package transformers

import (
	"fmt"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/wolftotem4/shaolin-ben-don/internal/client"
	typesjson "github.com/wolftotem4/shaolin-ben-don/internal/types/json"
)

type LinkItems struct {
	Items  []*typesjson.ProgressItem
	Client client.Client
}

func (link *LinkItems) String() string {
	var buffer []string
	for _, item := range link.Items {
		u, _ := link.Client.Endpoint(item.GetPath())

		if item.RemainSecondBeforeExpire < time.Hour {
			var expiration float32 = float32(item.RemainSecondBeforeExpire) / float32(time.Minute)
			buffer = append(buffer, fmt.Sprintf(
				"<b>%s</b>\n<i>即將在 %0.1f 分後過期</i>\n%s",
				tgbotapi.EscapeText(tgbotapi.ModeHTML, item.ShopName),
				expiration,
				tgbotapi.EscapeText(tgbotapi.ModeHTML, u.String()),
			))
		} else {
			buffer = append(buffer, fmt.Sprintf(
				"<b>#%s</b>\n%s",
				tgbotapi.EscapeText(tgbotapi.ModeHTML, item.ShopName),
				tgbotapi.EscapeText(tgbotapi.ModeHTML, u.String()),
			))
		}
	}
	return strings.Join(buffer, "\n\n")
}
