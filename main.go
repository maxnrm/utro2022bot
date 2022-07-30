package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	tgb "github.com/maxnrm/utro2022bot/bot"
	"github.com/maxnrm/utro2022bot/db"
	tt "github.com/maxnrm/utro2022bot/timetable"
	ws "github.com/maxnrm/utro2022bot/webserver"
	tele "gopkg.in/telebot.v3"
)

var dbHandler db.Handler = db.DBHandler

func main() {

	var Token string = os.Getenv("TELEGRAM_BOT_KEY")

	var timetableWrapper tt.Wrapper
	ttStr, err := dbHandler.GetTimetable()
	if err == nil {
		json.Unmarshal([]byte(ttStr), &timetableWrapper)
		timetableWrapper.FormatSelf()
	} else {
		println("Error", err)
	}

	teleBotSettings := tele.Settings{
		Token:  Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	r := gin.Default()
	b, _ := tele.NewBot(teleBotSettings)

	r.SetTrustedProxies(nil)

	ws.CreateRoutes(&timetableWrapper)(r)
	tgb.AddHandlers(&timetableWrapper)(b)

	go r.Run()
	b.Start()
}
