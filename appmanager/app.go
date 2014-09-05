package appmanager

import (
	"../mongodb"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type App struct {
	Id        bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	Name      string        `json:"name"      bson:"name"`
	Platforms []string      `json:"platforms" bson:"platforms"`
	Created   time.Time     `json:"created"   bson:"created"`
	Updated   time.Time     `json:"updated"   bson:"updated,omitempty"`
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
	app.Created = time.Now()
	return appCollection.Insert(app)
}

func ReadApp(id bson.ObjectId) (App, error) {
	var result App
	err := appCollection.Find(bson.M{"_id": id}).One(&result)
	return result, err
}

func UpdateApp(id bson.ObjectId, app *App) error {
	app.Update = time.Now()
	return appCollection.Update(bson.M{"_id": id}, app)
}

func DeleteApp(id bson.ObjectId) error {
	return appCollection.Remove(bson.M{"_id": id})
}
