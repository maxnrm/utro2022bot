package bot

import (
	"strconv"
	"strings"

	"github.com/maxnrm/utro2022bot/db"
	tt "github.com/maxnrm/utro2022bot/timetable"
	tele "gopkg.in/telebot.v3"
)

type dayInfo struct {
	photo         *tele.Photo
	timetableName string
}

var (
	dbHandler        db.Handler = db.DBHandler
	timetableWrapper tt.Wrapper = tt.TimetableWrapper

	setProgramBtns [][]tele.InlineButton = [][]tele.InlineButton{
		{tele.InlineButton{Text: "Урал. ЗОЖ-сообщества", Unique: "seturalzozh", Data: "зож"}},
		{tele.InlineButton{Text: "Урал. Вовлеченные сообщества", Unique: "seturalinvolv", Data: "вовл"}},
		{tele.InlineButton{Text: "Урал. Образовательные сообщества", Unique: "seturaledu", Data: "обр"}},
		{tele.InlineButton{Text: "Менторская программа", Unique: "setmentor", Data: "ментор"}},
	}

	timetableBtn [][]tele.InlineButton = [][]tele.InlineButton{{tele.InlineButton{Text: "Посмотреть расписание", Unique: "seetimetable"}}}

	dayInfoMap map[string]dayInfo = map[string]dayInfo{
		"01.08.2022": {photo: &tele.Photo{File: tele.FromDisk("./img/1.jpg")}, timetableName: "смена1_день1"},
		"02.08.2022": {photo: &tele.Photo{File: tele.FromDisk("./img/2.jpg")}, timetableName: "смена1_день2"},
		"03.08.2022": {photo: &tele.Photo{File: tele.FromDisk("./img/3.jpg")}, timetableName: "смена1_день3"},
		"04.08.2022": {photo: &tele.Photo{File: tele.FromDisk("./img/4.jpg")}, timetableName: "смена1_день4"},
		"05.08.2022": {photo: &tele.Photo{File: tele.FromDisk("./img/5.jpg")}, timetableName: "смена1_день5"},
		"06.08.2022": {photo: &tele.Photo{File: tele.FromDisk("./img/1.jpg")}, timetableName: "смена2_день1"},
		"07.08.2022": {photo: &tele.Photo{File: tele.FromDisk("./img/2.jpg")}, timetableName: "смена2_день2"},
		"08.08.2022": {photo: &tele.Photo{File: tele.FromDisk("./img/3.jpg")}, timetableName: "смена2_день3"},
		"09.08.2022": {photo: &tele.Photo{File: tele.FromDisk("./img/4.jpg")}, timetableName: "смена2_день4"},
		"10.08.2022": {photo: &tele.Photo{File: tele.FromDisk("./img/5.jpg")}, timetableName: "смена2_день5"},
	}

	setProgramInlineMarkup = &tele.ReplyMarkup{InlineKeyboard: setProgramBtns}
	timetableInlineMarkup  = &tele.ReplyMarkup{InlineKeyboard: timetableBtn}
	timetableHandler       = timetableHandlerFactory(&timetableWrapper)
)

func programCallbackHandlerFactory(btn *tele.InlineButton) tele.HandlerFunc {
	return func(c tele.Context) error {
		var user db.User

		user.ID = c.Chat().ID
		user.Group = btn.Data

		dbHandler.AddUser(&user, []string{"group"})

		c.Send("Ты выбрал программу «"+btn.Text+"». Теперь ты будешь получать расписание для этой программы.", timetableInlineMarkup)

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

		currentDateString := getToday()
		currentTimetableIndex := getTimetableIndex(ttw.Timetables, dayInfoMap[currentDateString].timetableName)

		currentTimetable := ttw.Timetables[currentTimetableIndex].Events
		currentFormattedTimetable := ttw.Timetables[currentTimetableIndex].FormattedEvents
		sendableMessages := make(map[int]string)

		for i := range currentTimetable {
			event := currentTimetable[i]
			formattedEvent := currentFormattedTimetable[i]

			isEmpty := strings.TrimSpace(formattedEvent) == ""
			isHidden := event.Hidden == "да"
			isMatchGroup := user.Group == event.Participants || strings.TrimSpace(event.Participants) == ""

			shouldSend := !isEmpty && !isHidden && isMatchGroup

			order, _ := strconv.Atoi(strings.TrimSpace(event.Order))

			if shouldSend {
				sendableMessages[order] += "\n\n" + formattedEvent
			}

		}

		lastEventIdx := len(sendableMessages)
		lastEventOrder, _ := strconv.Atoi(currentTimetable[len(currentTimetable)-1].Order)

		c.Send(dayInfoMap[currentDateString].photo, tele.ModeHTML)

		for i := 1; i <= lastEventOrder; i++ {
			if i == lastEventIdx {
				return c.Send(sendableMessages[i], tele.ModeHTML)
			} else {
				c.Send(sendableMessages[i], tele.ModeHTML)
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
