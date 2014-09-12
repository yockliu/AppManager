package appmanager

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"os"
	"os/exec"
)

type AppBuilder struct {
	appid   string
	running bool
}

var appbuilderMap map[string]*AppBuilder = make(map[string]*AppBuilder)

//func (ab *AppBuilder) Run1() {
//	stdout, err := ab.cmd.StdoutPipe()
//	if err != nil {
//		fmt.Println("Error: %s\n", err)
//		return
//	}
//	if err := ab.cmd.Start(); err != nil {
//		fmt.Println("Error: %s\n", err)
//		return
//	}
//	d, _ := ioutil.ReadAll(stdout)
//	if err := ab.cmd.Wait(); err != nil {
//		fmt.Println("Error: %s\n", err)
//		return
//	}
//	fmt.Println(string(d))
//}

func (ab *AppBuilder) RunBuild(appid string, versionid string, channels []string) error {
	if ab.running {
		return errors.New("正在打包")
	}

	app, err := ReadApp(bson.ObjectIdHex(ab.appid))
	if err != nil {
		return err
	}
	if len(app.ProjectPath) == 0 {
		err = errors.New("App的工程未设置")
		return err
	}

	version, err := ReadVersion(appid, bson.ObjectIdHex(versionid))
	if err != nil {
		return err
	}
	if len(version.GitTag) == 0 && len(version.GitIndex) == 0 {
		err = errors.New("版本的git路径未设置")
		return err
	}

	shcmd := "cd " + app.ProjectPath + "\n" + "git checkout " + version.GitTag + "\n" // + "./gradlew clean"
	cmd := exec.Command("/bin/sh", "-c", shcmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("appbuilder cmd start")
	ab.running = true
	cmd.Start()
	cmd.Wait()
	ab.running = false
	fmt.Println("cmd run end")
	return nil
}

func (ab *AppBuilder) isRunning() bool {
	return ab.running
}

func GetAppBuilder(appid string) (*AppBuilder, error) {
	fmt.Println("GetAppBuilder")
	if appbuilderMap[appid] == nil {
		app, err := ReadApp(bson.ObjectIdHex(appid))
		if err != nil {
			return nil, err
		}
		if len(app.ProjectPath) == 0 {
			err = errors.New("App的工程未设置")
			return nil, err
		}
		appbuilderMap[appid] = &AppBuilder{appid, false}
	}
	return appbuilderMap[appid], nil
}
