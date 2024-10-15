package main

import (
	"fmt"

	"github.com/slack-go/slack"
)

func main() {
	// アクセストークンを使用してクライアントを生成する
	tkn := "sample-token"
	c := slack.New(tkn)

	// MsgOptionText() の第二引数に true を設定すると特殊文字をエスケープする
	if _, _, err := c.PostMessage("#1_1_times_平野元気", slack.MsgOptionText("test slack text", true)); err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("success")
	}
}
