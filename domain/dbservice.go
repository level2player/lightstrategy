package domain

import (
	"log"

	"gopkg.in/mgo.v2"
)

const url = "127.0.0.1:27017"

var Session *mgo.Session

func init() {
	session, err := mgo.Dial(url)
	if err != nil {
		log.Println("mgo connect err")
	} else {
		Session = session
		log.Println("mgo connect sucess")
	}
}
