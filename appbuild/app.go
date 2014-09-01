package appbuild

import (
	"../mongodb"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type Version struct {
	Code        string
	Name        string
	Tag         string
	Commit_hash string
}

type Channel struct {
	Code string
	Name string
}

func Init() {
}

func ListVersion(platform string) ([]Version, error) {
	var result []Version
	versionCollection := mongodb.Mdb.C(platform + "_version")
	versionCollection.Find(bson.M{}).All(&result)
	return result, nil
}

func SaveVersion(platform string, version *Version) error {
	versionCollection := mongodb.Mdb.C(platform + "_version")
	err := versionCollection.Insert(version)
	return err
}

func ListChannels() ([]Channel, error) {
	var result []Channel
	channelCollection := mongodb.Mdb.C("channel")
	channelCollection.Find(bson.M{}).All(&result)
	return result, nil
}

func SaveChannel(channel *Channel) error {
	fmt.Println(channel)
	channelCollection := mongodb.Mdb.C("channel")
	err := channelCollection.Insert(channel)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
