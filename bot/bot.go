package bot

import (
	"log"

	"github.com/maxnrm/utro2022bot/db"
	tt "github.com/maxnrm/utro2022bot/timetable"
	tele "gopkg.in/telebot.v3"
)

const (
	uralEco    = "Урал.Эко-сообщества"
	uralUrb    = "Урал.Урбанистические сообщества"
	uralCreate = "Урал.Креативные сообщества"
	uralInvolv = "Урал.Вовлеченные сообщества"
	uralEdu    = "Урал.Образовательные сообщества"
	uralHealth = "Урал.ЗОЖ-сообщества"
)

var (
	dbHandler db.Handler = db.New()

	btnEco    = tele.InlineButton{Text: uralEco, Unique: "seturaleco"}
	btnUrb    = tele.InlineButton{Text: uralUrb, Unique: "seturalurb"}
	btnCreate = tele.InlineButton{Text: uralCreate, Unique: "seturalcreate"}
	btnInvolv = tele.InlineButton{Text: uralInvolv, Unique: "seturalinvolv"}
	btnEdu    = tele.InlineButton{Text: uralEdu, Unique: "seturaledu"}
	btnHealth = tele.InlineButton{Text: uralHealth, Unique: "seturalhealth"}

	setProgramInlineMarkup = &tele.ReplyMarkup{InlineKeyboard: setProgramInlineBtns}
	setProgramInlineBtns   = [][]tele.InlineButton{
		{btnEco},
		{btnUrb},
		{btnCreate},
		{btnInvolv},
		{btnEdu},
		{btnHealth},
	}
)

// AddHandlers creates telegram bot
func AddHandlers(timetableWrapper *tt.Wrapper) func(*tele.Bot) *tele.Bot {

	return func(b *tele.Bot) *tele.Bot {
		b.Use(miniLogger())
		// b.Use(checkUserSubscribed(&dbHandler))

		b.Handle("/start", startHandler)
		b.Handle("/help", helpHandler)

		b.Handle("/timetable", createTimetableHandler(timetableWrapper))

		// setprogram handlers
		b.Handle("/setprogram", setProgramHandler)
		b.Handle(&btnEco, createSetProgramVariantHandler(&btnEco, &dbHandler))
		b.Handle(&btnUrb, createSetProgramVariantHandler(&btnUrb, &dbHandler))
		b.Handle(&btnCreate, createSetProgramVariantHandler(&btnCreate, &dbHandler))
		b.Handle(&btnInvolv, createSetProgramVariantHandler(&btnInvolv, &dbHandler))
		b.Handle(&btnEdu, createSetProgramVariantHandler(&btnEdu, &dbHandler))
		b.Handle(&btnHealth, createSetProgramVariantHandler(&btnHealth, &dbHandler))

		return b
	}
}

func createTimetableHandler(ttw *tt.Wrapper) tele.HandlerFunc {
	return func(c tele.Context) error {

		currentTimetable := ttw.Timetables[0].Events
		currentFormattedTimetable := ttw.FormattedTimeTables[0].FormattedEvents

		var sendOptions tele.SendOptions
		sendOptions.ParseMode = tele.ModeHTML

		for i := range currentTimetable {
			c.Send(currentFormattedTimetable[i], tele.ModeHTML)
		}

		userID := c.Chat().ID

		user := dbHandler.GetUser(userID)
		if user.ID == 0 {
			return setProgramHandler(c)
		}

		return c.Send("")
	}
}

func startHandler(c tele.Context) error {
	err := c.Send(`Привет! Я Ботутра, твой бот-помощник!
Список команд:
/timetable - Посмотреть расписание
/setprogram - Выбрать программу
/help - Список комманд
`)
	setProgramHandler(c)
	return err
}

func helpHandler(c tele.Context) error {
	return c.Send(`Список команд:
/timetable - Посмотреть расписание
/setprogram - Выбрать программу
/help - Список комманд
`)
}

func setProgramHandler(c tele.Context) error {
	return c.Send("Выбери программу, расписание которой хочешь видеть в боте:", setProgramInlineMarkup)
}

func createSetProgramVariantHandler(btn *tele.InlineButton, dbHandler *db.Handler) tele.HandlerFunc {
	return func(c tele.Context) error {
		var user db.User

		user.ID = c.Chat().ID
		user.Group = btn.Unique

		dbHandler.AddUser(&user, []string{"group"})

		c.Send("Ты выбрал программу " + btn.Text + ". Теперь ты будешь получать расписание для этой программы.")

		return c.Delete()
	}
}

// func checkUserSubscribed(dbHandler *db.Handler) tele.MiddlewareFunc {

// 	return func(next tele.HandlerFunc) tele.HandlerFunc {
// 		return func(c tele.Context) error {

// 			// user = dbHandler.GetUser(userID)

// 			var importantChannelID int64 = os.GetEnv("IMPORTANT_CHANNEL_ID")
// 			importantChannel, err := c.Bot().ChatByID(importantChannelID)
// 			user, err := c.Bot().ChatMemberOf(importantChannel, c.Chat())
// 			if err != nil {
// 				fmt.Println("Chat ID is not ok!")
// 			}
// 			fmt.Println(user.Anonymous)
// 			channelLink := importantChannel.InviteLink

// 			//return if user is joined chat
// 			_, err = c.Bot().ChatMemberOf(importantChannel, c.Chat())
// 			if err == nil {
// 				return next(c)
// 			}

// 			c.Send("Кажется ты еще не подписан на наш канал с очень важными апдейтами! Держи ссылку" + channelLink)

// 			return next(c)
// 		}
// 	}
// }

func miniLogger() tele.MiddlewareFunc {
	l := log.Default()

	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			chatID := c.Chat().ID
			l.Println(chatID, "ok")
			return next(c)
		}
	}
}
