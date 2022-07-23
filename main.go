package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	tele "gopkg.in/telebot.v3"
)

// Event is timetable event
type Event struct {
	Hidden       string `json:"Скрыть"`
	Order        string `json:"Иерархия"`
	Participants string `json:"Учавствуют"`
	Time         string `json:"Время"`
	Name         string `json:"Название"`
	Description  string `json:"Описание"`
	Place        string `json:"Место"`
}

// Timetable is timetable
type Timetable struct {
	Name   string  `json:"name"`
	Events []Event `json:"rowObjects"`
}

// TimetableWrapper is you know
type TimetableWrapper struct {
	Timetable []Timetable `json:"timetable"`
}

func main() {
	var timetable TimetableWrapper
	var tt string

	go func() {
		r := gin.Default()

		r.POST("/timetable", func(c *gin.Context) {
			if c.BindJSON(&timetable) == nil {
				tt = ""
				for _, v := range timetable.Timetable[0].Events {
					tt += fmt.Sprintf("%s %s %s %s\n", v.Time, v.Name, v.Description, v.Place)
				}
				c.String(http.StatusOK, tt)
			}
		})

		r.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	}()

	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	// b.Use(middleware.Logger())

	b.Handle("/timetable", func(c tele.Context) error {
		return c.Send(tt)
		// return c.Send("timetable")
	})
	println("Bot started")
	b.Start()
}
