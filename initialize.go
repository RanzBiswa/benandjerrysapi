package main

import (
	"log"

	_apiInit "github.com/zalora_icecream/commonFramework/init"
	"github.com/zalora_icecream/commonFramework/setup"
)

func initialize() {

	var e error
	s, e := _apiInit.Initialize("zalora.cfg",
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
