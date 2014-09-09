package appmanager

import (
	"../mongodb"
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Version struct {
	Id       bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	Code     string        `json:"code"      bson:"code"`
	Name     string        `json:"name"      bson:"name"`
	Platform string        `json:"platform"  bson:"platform"`
	GitTag   string        `json:"git_tag"   bson:"git_tag"`
	GitIndex string        `json:"git_index" bson:"git_index"`
	Created  time.Time     `json:"created"   bson:"created"`
	Updated  time.Time     `json:"updated"   bson:"updated,omitempty"`
}

var versionCollectionMap map[string]*mgo.Collection

func versionCollection(appid string) (*mgo.Collection, error) {
	if !AppExists(bson.M{"id": appid, "validate": true}) {
		return nil, errors.New("无效的App ID")
	}
	if versionCollectionMap == nil {
		versionCollectionMap = make(map[string]*mgo.Collection)
	}
	cName := KeyOfVersionCollection(appid)
	if versionCollectionMap[cName] == nil {
		versionCollectionMap[cName] = mongodb.Mdb.C(cName)
	}
	return versionCollectionMap[cName], nil
}

func KeyOfVersionCollection(appid string) string {
	return "version_" + appid
}

func ListVersion(appid string, platform string) ([]Version, error) {
	var result []Version
	versionC, err := versionCollection(appid)
	if err == nil {
		var m bson.M
		if len(platform) != 0 {
			m = bson.M{"platform": platform}
		} else {
			m = bson.M{}
		}
		versionC.Find(m).All(&result)
	}
	return result, err
}

func CreateVersion(appid string, version *Version) error {
	versionC, err := versionCollection(appid)
	if err == nil {
		version.Created = time.Now()
		err = versionC.Insert(version)
	}
	return err
}

func ReadVersion(appid string, id int) (Version, error) {
	var result Version
	versionC, err := versionCollection(appid)
	if err == nil {
		err = versionC.Find(bson.M{"_id": id}).One(&result)
	}
	return result, err
}

func UpdateVersion(appid string, id bson.ObjectId, m map[string]interface{}) (Version, error) {
	var newVersion Version
	versionC, err := versionCollection(appid)
	if err == nil {
		var change = mgo.Change{
			ReturnNew: true,
			Update: bson.M{
				"$set": m,
			},
		}
		changeInfo, err := versionC.FindId(id).Apply(change, &newVersion)
		fmt.Println(changeInfo)
		fmt.Println(err)
	}
	return newVersion, err
}

func DeleteVersion(appid string, id bson.ObjectId) error {
	versionC, err := versionCollection(appid)
	if err == nil {
		err = versionC.RemoveId(id)
	}
	return err
}
