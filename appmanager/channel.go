package appmanager

import (
	"../mongodb"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Channel struct {
	Id      bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	Code    string        `json:"code"      bson:"code"`
	Name    string        `json:"name"      bson:"name"`
	Created time.Time     `json:"created"   bson:"created"`
	Updated time.Time     `json:"updated"   bson:"updated,omitempty"`
}

var channelCollectionMap map[string]*mgo.Collection

func channelCollection(app string, platform string) *mgo.Collection {
	if channelCollectionMap == nil {
		channelCollectionMap = make(map[string]*mgo.Collection)
	}
	cName := "channel_" + app + "_" + platform
	if channelCollectionMap[cName] == nil {
		channelCollectionMap[cName] = mongodb.Mdb.C(cName)
	}
	return channelCollectionMap[cName]
}

func ListChannels(app string, platform string) ([]Channel, error) {
	var result []Channel
	channelCollection(app, platform).Find(bson.M{}).All(&result)
	return result, nil
}

func CreateChannel(app string, platform string, channel *Channel) error {
	err := channelCollection(app, platform).Insert(channel)
	if err != nil {
		panic(err)
	}
	return err
}

func ReadChannel(app string, platform string, id int) (Channel, error) {
	var channel Channel
	err := channelCollection(app, platform).Find(bson.M{"_id": id}).One(&channel)
	return channel, err
}

func UpdateChannel(app string, platform string, channel Channel) error {
	err := channelCollection(app, platform).Update(bson.M{"_id": channel.Id}, channel)
	return err
}

func DeleteChannel(app string, platform string, id int) error {
	return channelCollection(app, platform).RemoveId(id)
}
