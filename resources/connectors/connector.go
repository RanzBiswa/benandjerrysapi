package connectors

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
)

var connection   *mgo.Session

// ConnectMongo connects the mongo db database
func ConnectMongo() *mgo.Session {
	connection, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
	}
	defer connection.Close()

	// Optional. Switch the session to a monotonic behavior.
	connection.SetMode(mgo.Monotonic, true)

	fmt.Print(connection.DB("IceCreams").CollectionNames())


	return connection.Clone()

}
