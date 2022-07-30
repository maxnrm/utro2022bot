package timetable

import (
	"fmt"
	"strings"
)

type sendEvent struct {
	pre string
	src string
}

//FormatSelf is for creating text events to send to tg
func (t *Wrapper) FormatSelf() {
	for i, tt := range t.Timetables {
		t.Timetables[i].FormattedEvents = make([]string, len(tt.Events))
		for j, event := range tt.Events {
			lines := []sendEvent{
				{pre: "", src: fmt.Sprintf("<b><u>%v</u></b>", event.Time)},
				{pre: "", src: strings.ToUpper(event.Name)},
				{pre: "<b>Спикеры:</b>", src: fmt.Sprintf("%v", event.Speakers)},
				{pre: "<b>Модератор:</b>", src: fmt.Sprintf("%v", event.Moderator)},
				{pre: "", src: fmt.Sprintf("%v", event.Description)},
				{pre: "", src: strings.ToUpper(event.Place)},
			}

			onlyFilledLines := []string{}
			for _, v := range skipEmptySendEvents(lines) {
				formattedLine := strings.TrimSpace(fmt.Sprintf("%v %v", v.pre, v.src))
				onlyFilledLines = append(onlyFilledLines, formattedLine)
			}
			stringEvent := strings.Join(onlyFilledLines, "\n")

			t.Timetables[i].FormattedEvents[j] = stringEvent
		}
	}
}
