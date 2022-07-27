package timetable

// Timetable is timetable
type Timetable struct {
	Name            string   `json:"name"`
	Day             string   `json:"day"`
	Week            string   `json:"week"`
	Events          []Event  `json:"rowObjects"`
	FormattedEvents []string `json:"-"`
}

// Event is timetable event
type Event struct {
	Hidden       string `json:"Скрыть"`
	Order        string `json:"Иерархия"`
	Participants string `json:"Участники"`
	Speakers     string `json:"Спикеры"`
	Moderator    string `json:"Модератор"`
	Time         string `json:"Время"`
	Name         string `json:"Название"`
	Description  string `json:"Описание"`
	Place        string `json:"Место"`
}

// Wrapper is you know
type Wrapper struct {
	Timetables []Timetable `json:"timetables"`
}
