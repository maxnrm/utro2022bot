package timetable

// Timetable is timetable
type Timetable struct {
	Name   string  `json:"name"`
	Events []Event `json:"rowObjects"`
}

type FormattedTimetable struct {
	Name            string `json:"name"`
	FormattedEvents []string
}

// Event is timetable event
type Event struct {
	Hidden       string `json:"Скрыть"`
	Order        string `json:"Иерархия"`
	Participants string `json:"Учавствуют"`
	Time         string `json:"Время"`
	Name         string `json:"Название"`
	Description  string `json:"Описание"`
	Place        string `json:"Место"`
}

// Wrapper is you know
type Wrapper struct {
	Timetables          []Timetable `json:"timetable"`
	FormattedTimeTables []FormattedTimetable
}
