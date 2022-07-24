package timetable

// Timetable is timetable
type Timetable struct {
	Name   string  `json:"name"`
	Events []Event `json:"rowObjects"`
}
