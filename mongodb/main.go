package mongodb

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

var Mdb *mgo.Database

func Init() {
	fmt.Println("mongodb init")

	session, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Println("mgo err")
		fmt.Println(err)
		panic(err)
	}
	//defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	Mdb = session.DB("test")
}
