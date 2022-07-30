package bot

import (
	"log"
	"os"
	"strconv"

	tele "gopkg.in/telebot.v3"
)

func isUserSubscribed(c tele.Context) bool {
	val, err := strconv.Atoi(os.Getenv("IMPORTANT_CHANNEL_ID"))
	if err != nil {
		println("Important channel ENV broken")
		return true
	}
	var importantChannelID int64 = int64(val)
	importantChannel, err := c.Bot().ChatByID(importantChannelID)
	if err != nil {
		println("Important channel ID seems to be wrong")
		return true
	}

	invite := importantChannel.InviteLink
	chatMember, err := c.Bot().ChatMemberOf(importantChannel, c.Chat())
	if err != nil {
		println("Important channel ID or user ID seems to be wrong")
		return true
	}

	isSubscribed := chatMember.Role != "left"

	if !isSubscribed {
		c.Send("Вижу ты не подписан на канал с очень важными обновлениями! Держиссылку " + invite)
	}

	return isSubscribed
}

func miniLogger() tele.MiddlewareFunc {
	l := log.Default()

	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			chatID := c.Chat().ID
			l.Println(chatID, "ok")
			return next(c)
		}
	}
}
