package main

import (
	"time"
)

// message は一つのメッセージを表す
type message struct {
	Name      string
	Message   string
	When      time.Time // TODO: メッセージを送信した時刻を表示させる
	AvatarURL string
}
