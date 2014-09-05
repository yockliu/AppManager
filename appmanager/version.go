package appmanager

import (
	"../mongodb"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Version struct {
	Id       bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	Code     string        `json:"code"      bson:"code"`
	Name     string        `json:"name"      bson:"name"`
	GitTag   string        `json:"git_tag"   bson:"git_tag"`
	GitIndex string        `json:"git_index" bson:"git_index"`
	Created  time.Time     `json:"created"   bson:"created"`
	Updated  time.Time     `json:"updated"   bson:"updated,omitempty"`
}

var cMap map[string]*mgo.Collection

func versionCollection(appid string, platform string) *mgo.Collection {
	if cMap == nil {
		cMap = make(map[string]*mgo.Collection)
	}
	cName = "version_" + appid + "_" + platform
	if cMap[cName] == nil {
		cMap[cName] = mongodb.Mdb.C(cName)
	}
	return cMap[cName]
}

func ListVersion(appid string, platform string) ([]Version, error) {
	var result []Version
	versionCollection(appid, platform).Find(bson.M{}).All(&result)
	return result, nil
}

func CreateVersion(appid string, platform string, version *Version) error {
	err := versionCollection(appid, platform).Insert(version)
	return err
}

func ReadVersion(appid string, platform string, id int) (Version, error) {
	var result Version
	err := versionCollection(appid, platform).Find(bson.M{"_id": id}).One(&result)
	return result, err
}

func UpdateVersion(appid string, platform string, version Version) error {
	return versionCollection(appid, platform).Update(bson.M{"_id": version.Id}, version)
}

func DeleteVersion(appid string, platform string, id int) error {
	return versionCollection(appid, platform).RemoveId(id)
}
