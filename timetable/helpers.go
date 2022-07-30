package timetable

import "strings"

func skipEmpty(str []string) []string {

	result := []string{}

	for _, value := range str {
		if strings.TrimSpace(value) != "" {
			result = append(result, value)
		}
	}

	return result
}

func skipEmptySendEvents(str []sendEvent) []sendEvent {

	result := []sendEvent{}

	for _, value := range str {
		if strings.TrimSpace(value.src) != "" {
			result = append(result, value)
		}
	}

	return result
}
