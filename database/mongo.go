package database

import (
	"time"

	"gopkg.in/mgo.v2"
)

type Collection struct {
}

const (
	HOST               = "localhost:27017"
	DATABSE            = "muxgolang"
	USERNAME           = "admin"
	PASSWORD           = "1111"
	COLLECTION_PRODUCT = "products"
)

func (c *Collection) GetCollection(name string) *mgo.Collection {
	info := &mgo.DialInfo{
		Addrs:    []string{HOST},
		Timeout:  60 * time.Second,
		Database: DATABSE,
	}
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}
	return session.DB(DATABSE).C(name)
}
