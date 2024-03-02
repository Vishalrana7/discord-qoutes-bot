package main

import(
	"fmt"
	"discord-qoute-bot/bot"
	"discord-qoute-bot/config"

)


func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}