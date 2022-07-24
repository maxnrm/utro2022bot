package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	tt "github.com/maxnrm/utro2022bot/timetable"
	tele "gopkg.in/telebot.v3"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func timetableHandler(timetable *tt.Wrapper) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.BindJSON(&timetable) == nil {
			resp := ""
			for _, v := range timetable.Timetable[0].Events {
				resp += fmt.Sprintf("%s %s %s %s\n", v.Time, v.Name, v.Description, v.Place)
			}
			c.String(http.StatusOK, resp)
		}
	}
}

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

	var timetableWrapper tt.Wrapper

	var Token string = os.Getenv("TELEGRAM_BOT_KEY")
	// var DatabaseURL string = os.Getenv("DATABASE_URL")

	r := gin.Default()

	r.POST("/timetable", timetableHandler(&timetableWrapper))

	r.GET("/ping", ping)

	go r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

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
