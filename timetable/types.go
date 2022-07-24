package timetable

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
	Timetable []Timetable `json:"timetable"`
}
