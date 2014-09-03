package appbuild

import (
	"../mongodb"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type Version struct {
	ObjectId    bson.ObjectId "_id"
	Code        string
	Name        string
	Tag         string
	CommitIndex string
}

func versionCollection(app string, platform string) *Collection {
	return mongodb.Mdb.C("version_" + app + "_" + platform)
}

func ListVersion(app string, platform string) ([]Version, error) {
	var result []Version
	versionCollection(app, platform).Find(bson.M{}).All(&result)
	return result, nil
}

func CreateVersion(app string, platform string, version *Version) error {
	err := versionCollection(app, platform).Insert(version)
	return err
}

func ReadVersion(app string, platform string, id int) (Version, error) {
	var result Version
	err := versionCollection(app, platform).Find(bson.M{"_id": id}).One(&result)
	return result, err
}

func UpdateVersion(app string, platform string, version Version) error {
	return versionCollection(app, platform).Update(bson.M{"_id", version.ObjectId}, version)
}

func DeleteVersion(app string, platfrom string, id int) error {
	return versionCollection(app, platform).RemoveId(id)
}
