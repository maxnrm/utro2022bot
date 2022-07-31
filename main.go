package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maxnrm/utro2022bot/bot"
	db "github.com/maxnrm/utro2022bot/db"
	"github.com/maxnrm/utro2022bot/webserver"
	tele "gopkg.in/telebot.v3"
)

var dbHandler db.Handler = db.DBHandler
var ws *gin.Engine = webserver.WebServer
var b *tele.Bot = bot.UtroBot

func main() {
	// run bot server and http server
	go ws.Run()
	b.Start()
}
