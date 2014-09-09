package appmanager

import (
	"../mongodb"
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Channel struct {
	Id       bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	Code     string        `json:"code"      bson:"code"`
	Name     string        `json:"name"      bson:"name"`
	Platform string        `json:"platform"  bson:"platform"`
	Created  time.Time     `json:"created"   bson:"created"`
	Updated  time.Time     `json:"updated"   bson:"updated,omitempty"`
}

var channelCollectionMap map[string]*mgo.Collection

func channelCollection(appid string) (*mgo.Collection, error) {
	if !AppExists(bson.M{"id": appid, "validate": true}) {
		return nil, errors.New("无效的App ID")
	}
	if channelCollectionMap == nil {
		channelCollectionMap = make(map[string]*mgo.Collection)
	}
	cName := KeyOfChannelCollection(appid)
	if channelCollectionMap[cName] == nil {
		channelCollectionMap[cName] = mongodb.Mdb.C(cName)
	}
	return channelCollectionMap[cName], nil
}

func KeyOfChannelCollection(appid string) string {
	return "channel_" + appid
}

func ListChannels(appid string, platform string) ([]Channel, error) {
	var result []Channel
	channelC, err := channelCollection(appid)
	if err == nil {
		var m bson.M
		if len(platform) != 0 {
			m = bson.M{"platform": platform}
		} else {
			m = bson.M{}
		}
		channelC.Find(m).All(&result)
	}
	return result, err
}

func CreateChannel(appid string, channel *Channel) error {
	channelC, err := channelCollection(appid)
	if err == nil {
		channel.Created = time.Now()
		channelC.Insert(channel)
	}
	return err
}

func ReadChannel(appid string, id int) (Channel, error) {
	var channel Channel
	channelC, err := channelCollection(appid)
	if err == nil {
		err = channelC.Find(bson.M{"_id": id}).One(&channel)
	}
	return channel, err
}

func UpdateChannel(appid string, id bson.ObjectId, m map[string]interface{}) (Channel, error) {
	var newChannel Channel
	channelC, err := channelCollection(appid)
	if err == nil {
		var change = mgo.Change{
			ReturnNew: true,
			Update: bson.M{
				"$set": m,
			},
		}
		changeInfo, err := channelC.FindId(id).Apply(change, &newChannel)
		fmt.Println(changeInfo)
		fmt.Println(err)
	}
	return newChannel, err
}

func DeleteChannel(appid string, id int) error {
	channelC, err := channelCollection(appid)
	if err == nil {
		err = channelC.RemoveId(id)
	}
	return err
}
