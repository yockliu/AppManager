package appbuild

import (
	"../mongodb"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type App struct {
	ObjectId  bson.ObjectId "_id"
	Name      string
	Platforms []string
}

var appCollection = mongodb.Mdb.C("app")

func ListApp() ([]App, error) {
	var result []App
	appCollection.Find(bson.M{}).All(&result)
	return result, nil
}

func CreateApp(app *App) error {
	return appCollection.Insert(app)
}

func ReadApp(id int) (App, error) {
	var result App
	err := appCollection.Find(bson.M{"_id", id}).One(&result)
	return result, err
}

func UpdateApp(app *App) error {
	return appCollection.Update(bson.M{"_id": app.ObjectId}, app)
}

func DeleteApp(id int) error {
	return appCpllection.Remove(bson.M{"_id": id})
}
