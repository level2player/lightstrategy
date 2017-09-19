package domain

import (
	"log"

	"gopkg.in/mgo.v2"
)

const url = "127.0.0.1:27017"

var Session *mgo.Session

func init() {
	var conncetErr error
	Session, conncetErr = mgo.Dial(url)
	if conncetErr != nil {
		log.Println("mgo connect err")
	}
}
