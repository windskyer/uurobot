package main

import (
	"fmt"
	"github.com/wechaty/go-wechaty/wechaty"
	"github.com/wechaty/go-wechaty/wechaty-puppet/schemas"
	"github.com/wechaty/go-wechaty/wechaty/user"
)

func main1() {
	_ = wechaty.NewWechaty().
		OnScan(func(ctx *wechaty.Context, qrCode string, status schemas.ScanStatus, data string) {
			fmt.Printf("Scan QR Code to login: %s\nhttps://wechaty.github.io/qrcode/%s\n", status, qrCode)
		}).
		OnLogin(func(ctx *wechaty.Context, user *user.ContactSelf) { fmt.Printf("User %s logined\n", user) }).
		OnMessage(func(ctx *wechaty.Context, message *user.Message) { fmt.Printf("Message: %s\n", message) }).
		Start()
}

func main() {
	bot := wechaty.NewWechaty()
	// TODO: init your bot at here...
	bot.Start()
}