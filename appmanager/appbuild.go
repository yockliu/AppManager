package appmanager

import (
	"archive/zip"
	"errors"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"sync"
)

type AppBuilder struct {
	appid           string
	platform        string
	scheduleCond    *sync.Cond
	taskMutex       *sync.Mutex
	scheduleLooping bool
}

var appbuilderMap map[string]*AppBuilder = make(map[string]*AppBuilder)
var appbuilderMutex sync.Mutex = sync.Mutex{}

func GetAppBuilder(appid string, platform string) (*AppBuilder, error) {
	fmt.Println("GetAppBuilder")
	appbuilderMutex.Lock()
	if appbuilderMap[appid] == nil {
		app, err := ReadApp(bson.ObjectIdHex(appid))
		if err != nil {
			return nil, err
		}
		if len(app.ProjectPath) == 0 {
			err = errors.New("App的工程未设置")
			return nil, err
		}
		ab := new(AppBuilder)
		ab.appid = appid
		ab.platform = platform
		locker := new(sync.Mutex)
		ab.scheduleCond = sync.NewCond(locker)
		ab.taskMutex = new(sync.Mutex)
		appbuilderMap[appid] = ab
	}
	appbuilderMutex.Unlock()
	return appbuilderMap[appid], nil
}

func (ab *AppBuilder) AddBuild(versionid string, channels []string) (*AppBuildTask, error) {
	task := new(AppBuildTask)
	task.Appid = ab.appid
	task.Platform = ab.platform
	task.Versionid = versionid
	task.Channels = channels

	ab.taskMutex.Lock()
	newTask, err := CreateAppBuildTask(task)
	ab.taskMutex.Unlock()

	go ab.scheduleLoop()

	return &newTask, err
}

func (ab *AppBuilder) CheckSchedule() {
	go ab.scheduleLoop()
}

func (ab *AppBuilder) scheduleLoop() {
	ab.scheduleCond.L.Lock()

	if ab.scheduleLooping == true {
		return
	}
	ab.scheduleLooping = true

	for {
		var task AppBuildTask
		ab.taskMutex.Lock()

		m := map[string]interface{}{}
		m["status"] = T_ABTask_ST_RUNNING
		task, err := FindAppBuildTask(m)
		fmt.Println("appbuild schedule loop find running task")
		fmt.Println(task)
		if err != nil || &task == nil {
			m["status"] = T_ABTask_ST_INIT
			task, err = FindAppBuildTask(m)
			fmt.Println("appbuild schedule loop find init task")
			fmt.Println(task)
			ab.taskMutex.Unlock()
			if err != nil || &task == nil {
				break
			}
		} else {
			ab.taskMutex.Unlock()
		}

		ab.runBuild(&task)
	}

	ab.scheduleLooping = false
	ab.scheduleCond.L.Unlock()
}

func (ab *AppBuilder) runBuild(task *AppBuildTask) error {
	ab.taskMutex.Lock()
	m := make(map[string]interface{})
	m["status"] = T_ABTask_ST_RUNNING
	UpdateAppBuildTask(task.Id, m)
	ab.taskMutex.Unlock()

	app, err := ReadApp(bson.ObjectIdHex(task.Appid))
	if err != nil {
		return err
	}
	if len(app.ProjectPath) == 0 {
		err = errors.New("App的工程未设置")
		return err
	}

	version, err := ReadVersion(task.Appid, bson.ObjectIdHex(task.Versionid))
	if err != nil {
		return err
	}
	if len(version.GitTag) == 0 && len(version.GitIndex) == 0 {
		err = errors.New("版本的git路径未设置")
		return err
	}

	changeDir := "cd " + app.ProjectPath + "\n"

	//gitStash := "git stash\n"

	checkoutTag := "git checkout " + version.GitTag + "\n"

	//gitStashPop := "git stash pop\n"

	gradleClean := "./gradlew clean\n"

	shcmd := changeDir + checkoutTag + gradleClean

	fmt.Println(shcmd)

	var chString string = ""
	for i, v := range task.Channels {
		chString += v
		if i < len(task.Channels) {
			chString += ","
		}
	}
	cmd := exec.Command("/bin/sh", "-c", "./build.sh "+app.Name+app.Id.Hex()+" "+app.ProjectPath+" "+version.GitTag+" "+chString)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("appbuilder cmd start")
	cmd.Start()
	cmd.Wait()
	fmt.Println("appbuilder cmd run end")

	workPath, _ := os.Getwd()
	sourceDir := workPath + "/static/apk/" + app.Name + app.Id.Hex() + "/" + version.GitTag + "/"
	destFile := workPath + "/static/apk/zip/" + task.Id.Hex() + ".zip"
	var countOfChannels = len(task.Channels)
	var sourceFileNames = make([]string, countOfChannels)
	for i, v := range task.Channels { //zhoumo-a1-release.apk
		sourceFileNames[i] = "zhoumo-" + strings.Replace(v, "and-", "", 1) + "-release.apk"
	}
	zipResult := zipFiles(sourceDir, sourceFileNames, destFile)

	ab.taskMutex.Lock()
	m = make(map[string]interface{})
	if !zipResult {
		m["status"] = T_ABTask_ST_ERR
	} else {
		m["status"] = T_ABTask_ST_FINISH
		m["file"] = "/apk/zip/" + task.Id.Hex() + ".zip"
	}
	UpdateAppBuildTask(task.Id, m)
	ab.taskMutex.Unlock()

	return nil
}

func zipFiles(dir string, sourceNames []string, to string) bool {
	zipfile, err := os.Create(to)
	if err != nil {
		return false
	}
	defer zipfile.Close()

	zipw := zip.NewWriter(zipfile)
	for _, filename := range sourceNames {
		fmt.Println(dir + filename)
		fmt.Println("...")
		file, err := os.Open(dir + filename)
		if err != nil {
			return false
		}
		defer file.Close()
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			return false
		}
		zipF, err := zipw.Create(filename)
		if err != nil {
			return false
		}
		_, err = zipF.Write([]byte(bytes))
		if err != nil {
			return false
		}
	}
	defer zipw.Close()
	return true
}
