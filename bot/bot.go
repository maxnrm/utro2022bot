package bot

import (
	tt "github.com/maxnrm/utro2022bot/timetable"
	tele "gopkg.in/telebot.v3"
)

// AddHandlers creates telegram bot
func AddHandlers(timetableWrapper *tt.Wrapper) func(*tele.Bot) *tele.Bot {

	return func(b *tele.Bot) *tele.Bot {
		b.Use(miniLogger())
		b.Handle("/start", startHandler)
		b.Handle("/help", helpHandler)
		b.Handle("/timetable", timetableHandler)
		b.Handle("/setprogram", setProgramHandler)
		return b
	}
}
