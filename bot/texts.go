package bot

import "strings"

var (
	startTextArr []string = []string{"Привет! Я Ботутра, твой бот-помощник!",
		"",
		"Список команд:",
		"/timetable - Посмотреть расписание",
		"/setprogram - Выбрать программу",
		"/help - Список комманд",
		"",
		"Команды так же доступны по кнопке Menu, слева от строки ввода текста"}

	helpTextArr []string = []string{"Список команд:",
		"",
		"/timetable - Посмотреть расписание",
		"/setprogram - Выбрать программу",
		"/help - Список комманд",
		"",
		"Команды так же доступны по кнопке Menu, слева от строки ввода текста"}

	startText      string = strings.Join(startTextArr, "\n")
	helpText       string = strings.Join(helpTextArr, "\n")
	setProgramText string = strings.Join([]string{"Выбери программу, расписание которой хочешь видеть в боте:"}, "")
)
