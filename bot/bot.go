package bot

import (
	"fmt"
	"log"

	tt "github.com/maxnrm/utro2022bot/timetable"
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

// AddHandlers creates telegram bot
func AddHandlers(timetableWrapper *tt.Wrapper) func(*tele.Bot) *tele.Bot {
	return func(b *tele.Bot) *tele.Bot {
		b.Use(miniLogger())

		b.Handle("/timetable", func(c tele.Context) error {
			resp := ""
			for _, v := range timetableWrapper.Timetable[0].Events {
				resp += fmt.Sprintf("%s %s %s %s\n", v.Time, v.Name, v.Description, v.Place)
			}

			return c.Send(resp)

		})

		return b
	}
}
