package appmanager

import (
	"../mongodb"
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type App struct {
	Id          bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	Name        string        `json:"name"      bson:"name"`
	Platforms   []string      `json:"platforms" bson:"platforms"`
	ProjectPath string        `json:"prj_path"  bson:"prj_path"`
	Created     time.Time     `json:"created"   bson:"created"`
	Updated     time.Time     `json:"updated"   bson:"updated,omitempty"`
	Forbidden   bool          `json:"forbidden" bson:"forbidden"`
	Validate    bool          `json:"validate"  bson:"validate"`
}

var appCollection *mgo.Collection

func InitApp() {
	appCollection = mongodb.Mdb.C("app")
}

func AppExists(where bson.M) bool {
	existCount, _ := appCollection.Find(where).Count()
	if existCount > 0 {
		return true
	} else {
		return false
	}
}

func ListApp() ([]App, error) {
	var result []App
	appCollection.Find(bson.M{"validate": true}).All(&result)
	return result, nil
}

func CreateApp(app *App) (App, error) {
	var newApp App
	exist := AppExists(bson.M{"name": app.Name, "validate": true})
	if exist {
		return newApp, errors.New("同名App已存在")
	}
	newId := bson.NewObjectId()
	app.Id = newId
	app.Created = time.Now()
	app.Forbidden = false
	app.Validate = true
	err := appCollection.Insert(app)
	if err == nil {
		appCollection.FindId(newId).One(&newApp)
	}
	return newApp, err
}

func ReadApp(id bson.ObjectId) (App, error) {
	var result App
	err := appCollection.Find(bson.M{"_id": id}).One(&result)

	if result.Validate == false && err == nil {
		err = errors.New("无效的App")
	}

	return result, err
}

func UpdateApp(id bson.ObjectId, m map[string]interface{}) (App, error) {
	var result App
	//	err := appCollection.Find(bson.M{"_id": id, "validate": true}).One(&result)
	//	if err != nil {
	//		err = errors.New("修改的App不存在或已删除")
	//		return result, err
	//	}
	m["updated"] = time.Now()
	var change = mgo.Change{
		ReturnNew: true,
		Update: bson.M{
			"$set": m,
		},
	}
	changeInfo, err := appCollection.FindId(id).Apply(change, &result)
	fmt.Println(changeInfo)
	fmt.Println(err)
	if err != nil {
		return result, err
	}
	return result, nil
}

func DeleteApp(id bson.ObjectId) error {
	var result App
	err := appCollection.Find(bson.M{"_id": id, "validate": true}).One(&result)
	if err != nil {
		err = errors.New("删除的App不存在或已删除")
		return err
	}
	var change = mgo.Change{
		ReturnNew: false,
		Update: bson.M{
			"$set": bson.M{
				"validate": false,
			},
		},
	}
	fmt.Println("change")
	changeInfo, err := appCollection.FindId(id).Apply(change, &result)
	fmt.Println(changeInfo)
	fmt.Println(err)
	if err != nil {
		return err
	}
	return nil
}
