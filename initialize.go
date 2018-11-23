package main

import (
	"log"

	_apiInit "github.com/benandjerrysapi/commonFramework/init"
	"github.com/benandjerrysapi/commonFramework/setup"
)

func initialize() {

	var e error
	s, e := _apiInit.Initialize("benandjerrysapi.cfg",
		setup.IceCreamAPI)
	if e != nil {
		log.Fatal(e)
	}

	if s == nil {
		log.Fatal("Failed to initialize")
	}

	c = s.C
	port = s.Port
	secure = s.Secure
	pool = s.Pool
}
