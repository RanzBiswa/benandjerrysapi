package connectors

import (
	"github.com/benandjerrysapi/commonFramework/loggers"
	"github.com/benandjerrysapi/commonFramework/setup"
	"github.com/benandjerrysapi/configs"
	"gopkg.in/mgo.v2"
)

var connection *mgo.Session

var dbAddress []string

// ConnectMongo connects the mongo db database
func ConnectMongo() (*mgo.Session, string) {

	mgoDialInfo := &mgo.DialInfo{
		Addrs: configs.DbConfigs["mongo"].DBAddress,
		// This can be used when we have authorization to mongo db
		// Username:"",
		// Password:"",
	}
	connection, err := mgo.DialWithInfo(mgoDialInfo)
	//connection, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		loggers.LogError(setup.EvtAPIHandlerError,
			"Connect Mongo", "Conenction Error", nil)
		return nil,
			err.Error()
	}
	defer connection.Close()

	// Optional. Switch the session to a monotonic behavior.
	connection.SetMode(mgo.Monotonic, true)

	return connection.Clone(), ""

}
