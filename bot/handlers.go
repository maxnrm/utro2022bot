package bot

import (
	"strings"

	"github.com/maxnrm/utro2022bot/db"
	tt "github.com/maxnrm/utro2022bot/timetable"
	tele "gopkg.in/telebot.v3"
)

var (
	dbHandler        db.Handler = db.DBHandler
	timetableWrapper tt.Wrapper

	setProgramBtns [][]tele.InlineButton = [][]tele.InlineButton{
		{tele.InlineButton{Text: "Урал.Эко-сообщества", Unique: "seturaleco", Data: "эко"}},
		{tele.InlineButton{Text: "Урал.Урбанистические сообщества", Unique: "seturalurb", Data: "урб"}},
		{tele.InlineButton{Text: "Урал.Креативные сообщества", Unique: "seturalcreate", Data: "креа"}},
		{tele.InlineButton{Text: "Урал.Вовлеченные сообщества", Unique: "seturalinvolv", Data: "вовл"}},
		{tele.InlineButton{Text: "Урал.Образовательные сообщества", Unique: "seturaledu", Data: "обр"}},
		{tele.InlineButton{Text: "Урал.ЗОЖ-сообщества", Unique: "seturalhealth", Data: "зож"}},
	}

	setProgramInlineMarkup = &tele.ReplyMarkup{InlineKeyboard: setProgramBtns}

	timetableHandler = timetableHandlerFactory(&timetableWrapper)
)

var (
	programCallbackHandlers []tele.HandlerFunc
)

func programCallbackHandlerFactory(btn *tele.InlineButton) tele.HandlerFunc {
	return func(c tele.Context) error {
		var user db.User

		user.ID = c.Chat().ID
		user.Group = btn.Data

		dbHandler.AddUser(&user, []string{"group"})

		c.Send("Ты выбрал программу " + btn.Text + ". Теперь ты будешь получать расписание для этой программы.")

		return c.Delete()
	}
}

func timetableHandlerFactory(ttw *tt.Wrapper) tele.HandlerFunc {
	return func(c tele.Context) error {
		userID := c.Chat().ID
		user := dbHandler.GetUser(userID)
		if user.ID == 0 {
			return setProgramHandler(c)
		}

		if !isUserSubscribed(c) {
			return c.Send("")
		}

		currentTimetable := ttw.Timetables[0].Events
		currentFormattedTimetable := ttw.Timetables[0].FormattedEvents

		for i := range currentTimetable {
			event := currentTimetable[i]
			formattedEvent := currentFormattedTimetable[i]

			isEmpty := strings.TrimSpace(formattedEvent) == ""
			isHidden := event.Hidden == "да"
			isMatchGroup := user.Group == event.Participants || strings.TrimSpace(event.Participants) == ""

			shouldSend := !isEmpty && !isHidden && isMatchGroup

			if shouldSend {
				c.Send(formattedEvent, tele.ModeHTML)
			}
		}

		return c.Send("")
	}
}

func setProgramHandler(c tele.Context) error {
	return c.Send(setProgramText, setProgramInlineMarkup)
}

func startHandler(c tele.Context) error {
	err := c.Send(startText)
	setProgramHandler(c)
	return err
}

func helpHandler(c tele.Context) error {
	return c.Send(helpText)
}
