package bot

import (
	"fmt"
	"log"

	"github.com/maxnrm/utro2022bot/db"
	tt "github.com/maxnrm/utro2022bot/timetable"
	tele "gopkg.in/telebot.v3"
)

type name struct {
	full  string
	short string
}

const (
	uralEco    = "Урал.Эко-сообщества"
	uralUrb    = "Урал.Урбанистические сообщества"
	uralCreate = "Урал.Креативные сообщества"
	uralInvolv = "Урал.Вовлеченные сообщества"
	uralEdu    = "Урал.Образовательные сообщества"
	uralHealth = "Урал.ЗОЖ-сообщества"
)

var (
	btnEco    = tele.InlineButton{Text: uralEco, Unique: "seturaleco"}
	btnUrb    = tele.InlineButton{Text: uralUrb, Unique: "seturalurb"}
	btnCreate = tele.InlineButton{Text: uralCreate, Unique: "seturalcreate"}
	btnInvolv = tele.InlineButton{Text: uralInvolv, Unique: "seturalinvolv"}
	btnEdu    = tele.InlineButton{Text: uralEdu, Unique: "seturaledu"}
	btnHealth = tele.InlineButton{Text: uralHealth, Unique: "seturalhealth"}

	setProgramInlineBtns = [][]tele.InlineButton{
		{btnEco},
		{btnUrb},
		{btnCreate},
		{btnInvolv},
		{btnEdu},
		{btnHealth},
	}
	setProgramInlineMarkup = &tele.ReplyMarkup{InlineKeyboard: setProgramInlineBtns}
)

// AddHandlers creates telegram bot
func AddHandlers(timetableWrapper *tt.Wrapper) func(*tele.Bot) *tele.Bot {

	dbHandler := db.New()

	return func(b *tele.Bot) *tele.Bot {
		b.Use(miniLogger())
		// b.Use(checkUserSubscribed(&dbHandler))

		b.Handle("/timetable", func(c tele.Context) error {
			resp := ""
			for _, v := range timetableWrapper.Timetable[0].Events {
				resp += fmt.Sprintf("%s %s %s %s\n", v.Time, v.Name, v.Description, v.Place)
			}

			return c.Send(resp)

		})

		b.Handle("/setprogram", func(c tele.Context) error {
			return c.Send("Hello!", setProgramInlineMarkup)
		})

		b.Handle(&btnEco, createSetProgramHandler(&btnEco, &dbHandler))
		b.Handle(&btnUrb, createSetProgramHandler(&btnUrb, &dbHandler))
		b.Handle(&btnCreate, createSetProgramHandler(&btnCreate, &dbHandler))
		b.Handle(&btnInvolv, createSetProgramHandler(&btnInvolv, &dbHandler))
		b.Handle(&btnEdu, createSetProgramHandler(&btnEdu, &dbHandler))
		b.Handle(&btnHealth, createSetProgramHandler(&btnHealth, &dbHandler))

		return b
	}
}

func createSetProgramHandler(btn *tele.InlineButton, dbHandler *db.Handler) tele.HandlerFunc {
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
			update := c.Update()
			messageID := update.ID
			l.Println(messageID, "ok")
			return next(c)
		}
	}
}
