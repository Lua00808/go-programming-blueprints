package main

import (
	"log"

	"gopkg.in/mgo.v2"
)

var db *mgo.Session

func dialDb() error {
	var err error
	log.Println("MongoDB にダイアル中: localhost")
	db, err = mgo.Dial("localhost")
	return err
}
func closeDb() {
	db.Close()
	log.Println("データベース接続が閉じられました")
}
