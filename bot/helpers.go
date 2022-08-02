package bot

import (
	"time"

	tt "github.com/maxnrm/utro2022bot/timetable"
	"golang.org/x/exp/slices"
)

func getToday() string {
	today := time.Now()
	format := "02.01.2006"
	shift := -time.Hour * 6

	return today.Add(shift).Format(format)
}

func getTimetableIndex(tts []tt.Timetable, name string) int {
	idx := slices.IndexFunc(tts, func(t tt.Timetable) bool { return t.Name == name })

	return idx
}
