package main

import (
	"log"

	"github.com/nsqio/go-nsq"
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

type poll struct {
	Options []string
}

func loadOptions() ([]string, error) {
	var options []string
	iter := db.DB("ballots").C("polls").Find(nil).Iter()
	var p poll
	for iter.Next(&p) {
		options = append(options, p.Options...)
	}
	iter.Close()
	return options, iter.Err()
}

func publishVotes(votes <-chan string) <-chan struct{} {
	stopchan := make(chan struct{}, 1)
	pub, _ := nsq.NewProducer("localhost:4150", nsq.NewConfig())
	go func() {
		defer func() {
			stopchan <- struct{}{}
		}()
		for vote := range votes {
			pub.Publish("votes", []byte(vote)) // 投票内容をパブリッシュします
		}
		log.Println("Publisher: 停止中です")
		pub.Stop()
		log.Println("Publisher: 停止しました")
	}()
	return stopchan
}
