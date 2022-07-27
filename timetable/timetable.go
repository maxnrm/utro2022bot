package timetable

import (
	"fmt"
	"strings"
)

// FormatSelf is for
func (tt *Wrapper) FormatSelf() {
	formattedTimetables := []FormattedTimetable{}
	for _, v := range tt.Timetables {
		formattedTimetables = append(formattedTimetables, formatTimetableForSend(&v))
	}

	tt.FormattedTimeTables = formattedTimetables
}

func skipEmpty(str []string) []string {

	result := []string{}

	for _, value := range str {
		if strings.TrimSpace(value) != "" {
			result = append(result, value)
		}
	}

	return result
}

func formatTimetableForSend(tt *Timetable) FormattedTimetable {
	var ffTimetable FormattedTimetable
	formattedEvents := []string{}

	for _, event := range tt.Events {
		lines := []string{
			fmt.Sprintf("<b><u>%v</u></b>", event.Time),
			fmt.Sprintf("%v", event.Name),
			// fmt.Sprintf("<b>Спикеры:</b>%v", event.Speakers),
			// fmt.Sprintf("<b>Модератор:<b>%v", event.Moderator),
			fmt.Sprintf("%v", event.Description),
			fmt.Sprintf("%v", event.Place),
		}

		onlyFilledLines := skipEmpty(lines)
		stringEvent := strings.Join(onlyFilledLines, "\n")

		formattedEvents = append(formattedEvents, stringEvent)

	}

	ffTimetable.Name = tt.Name
	ffTimetable.FormattedEvents = skipEmpty(formattedEvents)

	return ffTimetable
}
