package appbuild

import (
	"../mongodb"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type Channel struct {
	ObjectId bson.ObjectId "_id"
	Code     string
	Name     string
}

func channelCollection(app string, platform string) *Collection {
	return mongodb.Mdb.C("channel_" + app + "_" + platform)
}

func ListChannels(app string, platform string) ([]Channel, error) {
	var result []Channel
	channelCollection(app, platform).Find(bson.M{}).All(&result)
	return result, nil
}

func CreateChannel(app string, platform string, channel *Channel) error {
	err := channelCollection(app, platform).Insert(channel)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func ReadChannel(app string, platform string, id int) (Channel, error) {
	var channel Channel
	err := channelCollection(app, platform).Find(bson.M{"_id": id}).One(&channel)
	return channel, err
}

func UpdateChannel(app string, platform string, channel Channel) error {
	err := channelCollection(app, platform).Update(bson.M{"_id": channel.ObjectId}, channel)
	return err
}

func DeleteChannel(app string, platform string, id int) error {
	return channelCollection(app, platform).RemoveId(id)
}
