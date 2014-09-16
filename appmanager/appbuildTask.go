package appmanager

import (
	"../mongodb"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type T_ABTask_Status int

const (
	T_ABTask_ST_ERR     = -1
	T_ABTask_ST_INIT    = iota
	T_ABTask_ST_RUNNING = iota
	T_ABTask_ST_FINISH  = iota
)

type AppBuildTask struct {
	Id        bson.ObjectId   `json:"id"          bson:"_id,omitempty"`
	Appid     string          `json:"appid"       bson:"appid"`
	Platform  string          `json:"platform"    bson:"platform"`
	Versionid string          `json:"versionid"   bson:"versionid"`
	Channels  []string        `json:"channels"    bson:"channels"`
	Status    T_ABTask_Status `json:"status"      bson:"status"`
	Created   time.Time       `json:"created"     bson:"created"`
	Updated   time.Time       `json:"updated"     bson:"updated"`
}

var taskCollection *mgo.Collection

func InitAppBuildTask() {
	taskCollection = mongodb.Mdb.C("app_build_task")
}

func ReadAppBuildTaskList(m map[string]interface{}) ([]AppBuildTask, error) {
	fmt.Println(m)
	fmt.Println(bson.M(m))
	var list []AppBuildTask
	err := taskCollection.Find(m).All(&list)
	return list, err
}

func CreateAppBuildTask(task *AppBuildTask) (AppBuildTask, error) {
	var newTask AppBuildTask
	task.Id = bson.NewObjectId()
	task.Created = time.Now()
	task.Status = T_ABTask_ST_INIT
	err := taskCollection.Insert(task)

	if err == nil {
		taskCollection.FindId(task.Id).One(&newTask)
	}

	return newTask, err
}

func ReadAppBuildTask(id bson.ObjectId) (AppBuildTask, error) {
	var newTask AppBuildTask
	err := taskCollection.FindId(id).One(&newTask)
	return newTask, err
}

func FindAppBuildTask(m bson.M) (AppBuildTask, error) {
	var newTask AppBuildTask
	err := taskCollection.Find(m).One(&newTask)
	return newTask, err
}

func UpdateAppBuildTask(id bson.ObjectId, m map[string]interface{}) (AppBuildTask, error) {
	var newTask AppBuildTask
	m["updated"] = time.Now()
	var change = mgo.Change{
		ReturnNew: true,
		Update: bson.M{
			"$set": m,
		},
	}
	changeInfo, err := taskCollection.FindId(id).Apply(change, &newTask)
	fmt.Println(changeInfo)
	fmt.Println(err)
	if err != nil {
		return newTask, err
	}
	return newTask, nil
}
