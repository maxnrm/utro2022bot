package bot

import "strings"

var (
	startTextArr []string = []string{"Привет! Я Ботутра, твой бот-помощник!",
		"Список команд:",
		"/timetable - Посмотреть расписание",
		"/setprogram - Выбрать программу",
		"/help - Список комманд"}

	helpTextArr []string = []string{"Список команд:",
		"/timetable - Посмотреть расписание",
		"/setprogram - Выбрать программу",
		"/help - Список комманд"}

	startText      string = strings.Join(startTextArr, "\n")
	helpText       string = strings.Join(helpTextArr, "\n")
	setProgramText string = strings.Join([]string{"Выбери программу, расписание которой хочешь видеть в боте:"}, "")
)
