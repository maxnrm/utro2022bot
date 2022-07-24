package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	tt "github.com/maxnrm/utro2022bot/timetable"
	ws "github.com/maxnrm/utro2022bot/webserver"
	tele "gopkg.in/telebot.v3"
)

func miniLogger() tele.MiddlewareFunc {
	l := log.Default()

	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			update := c.Update()
			messageID := update.Message.ID
			// data, _ := json.MarshalIndent(update, "", "  ")
			l.Println(messageID, "ok")
			return next(c)
		}
	}
}

func main() {

	var Token string = os.Getenv("TELEGRAM_BOT_KEY")
	// var DatabaseURL string = os.Getenv("DATABASE_URL")

	var timetableWrapper tt.Wrapper

	r := gin.Default()
	Routes := ws.CreateRoutes(&timetableWrapper)
	Routes(r)
	go r.Run()

	pref := tele.Settings{
		Token:  Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Use(miniLogger())

	b.Handle("/timetable", func(c tele.Context) error {
		resp := ""
		for _, v := range timetableWrapper.Timetable[0].Events {
			resp += fmt.Sprintf("%s %s %s %s\n", v.Time, v.Name, v.Description, v.Place)
		}

		return c.Send(resp)
		// return c.Send("timetable")
	})

	println("Bot started")
	b.Start()
}
