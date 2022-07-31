package timetable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/maxnrm/utro2022bot/db"
)

var dbHandler db.Handler = db.DBHandler

// TimetableWrapper is gloabal timetable object
var TimetableWrapper = New()

//New create new wrapper
func New() Wrapper {
	var timetableWrapper *Wrapper

	ttStr, err := dbHandler.GetTimetable()
	if err == nil {
		json.Unmarshal([]byte(ttStr), &timetableWrapper)
		timetableWrapper.FormatSelf()
	} else {
		println("Error", err)
	}

	return *timetableWrapper
}

//FormatSelf is for creating text events to send to tg
func (t *Wrapper) FormatSelf() {
	for i, tt := range t.Timetables {
		t.Timetables[i].FormattedEvents = make([]string, len(tt.Events))
		for j, event := range tt.Events {
			lines := []sendEvent{
				{pre: "&#128337;<b><u>", post: "</u></b>", src: event.Time},
				{pre: "", post: "", src: strings.ToUpper(event.Name)},
				{pre: "<b>Спикеры:</b>", post: "", src: event.Speakers},
				{pre: "<b>Модератор:</b>", post: "", src: event.Moderator},
				{pre: "", post: "", src: event.Description},
				{pre: "&#128205;", post: "", src: strings.ToUpper(event.Place)},
			}

			onlyFilledLines := []string{}
			for _, v := range skipEmptySendEvents(lines) {
				formattedLine := strings.TrimSpace(fmt.Sprintf("%v %v %v", v.pre, v.src, v.post))
				onlyFilledLines = append(onlyFilledLines, formattedLine)
			}
			stringEvent := strings.Join(onlyFilledLines, "\n")

			t.Timetables[i].FormattedEvents[j] = stringEvent
		}
	}
}
