package bot

import (
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

var token string = os.Getenv("TELEGRAM_BOT_KEY")

// UtroBot is global bot instance
var UtroBot *tele.Bot = New()

// New create new telebot instance with added command handlers
func New() *tele.Bot {

	settings := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(settings)
	if err != nil {

	}

	addHandlers(b)

	return b
}

func addHandlers(b *tele.Bot) *tele.Bot {

	b.Use(miniLogger())

	for _, btn := range setProgramBtns {
		btnValue := &btn[0]
		b.Handle(btnValue, programCallbackHandlerFactory(btnValue))
	}

	b.Handle("/start", startHandler)
	b.Handle("/help", helpHandler)
	b.Handle("/timetable", timetableHandler)
	b.Handle("/setprogram", setProgramHandler)
	return b
}
