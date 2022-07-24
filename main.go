package main

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	tgb "github.com/maxnrm/utro2022bot/bot"
	tt "github.com/maxnrm/utro2022bot/timetable"
	ws "github.com/maxnrm/utro2022bot/webserver"
	tele "gopkg.in/telebot.v3"
)

func main() {

	var Token string = os.Getenv("TELEGRAM_BOT_KEY")

	var timetableWrapper tt.Wrapper

	teleBotSettings := tele.Settings{
		Token:  Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	r := gin.Default()
	b, _ := tele.NewBot(teleBotSettings)

	ws.CreateRoutes(&timetableWrapper)(r)
	tgb.AddHandlers(&timetableWrapper)(b)

	go r.Run()
	b.Start()
}
