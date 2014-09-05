package appmanager

import (
	"../mongodb"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type App struct {
	Id        bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	Name      string        `json:"name"      bson:"name"`
	Platforms []string      `json:"platforms" bson:"platforms"`
}

var appCollection *mgo.Collection

func InitApp() {
	appCollection = mongodb.Mdb.C("app")
}

func ListApp() ([]App, error) {
	var result []App
	appCollection.Find(bson.M{}).All(&result)
	return result, nil
}

func CreateApp(app *App) error {
	return appCollection.Insert(app)
}

func ReadApp(id bson.ObjectId) (App, error) {
	var result App
	err := appCollection.Find(bson.M{"_id": id}).One(&result)
	return result, err
}

func UpdateApp(id bson.ObjectId, app *App) error {
	return appCollection.Update(bson.M{"_id": id}, app)
}

func DeleteApp(id bson.ObjectId) error {
	return appCollection.Remove(bson.M{"_id": id})
}
