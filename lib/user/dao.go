package user

import (
	"log"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserDao struct {
	Server   string
	Database string
}

var db *mgo.Database

func (u *UserDao) Connect() {
	session, err := mgo.Dial(u.Server)
	if err != nil {

		log.Fatal(err)
	}
	db = session.DB(u.Database)
}