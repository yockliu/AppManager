package appmanager

import (
	"../mongodb"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type App struct {
	Id        bson.ObjectId "_id,omitempty"
	Name      string        "name"
	Platforms []string      "platforms"
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

func ReadApp(id int) (App, error) {
	var result App
	err := appCollection.Find(bson.M{"_id": id}).One(&result)
	return result, err
}

func UpdateApp(app *App) error {
	return appCollection.Update(bson.M{"_id": app.Id}, app)
}

func DeleteApp(id int) error {
	return appCollection.Remove(bson.M{"_id": id})
}
