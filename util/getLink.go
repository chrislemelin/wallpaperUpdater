package util

import (
	"net/http"
	"os"
	"time"

	_ "github.com/turnage/graw"
	"github.com/turnage/graw/reddit"
)

type data struct {
	data string
	kind string
}

var (
	username = os.Getenv("REDDIT_USERNAME")
	password = os.Getenv("REDDIT_PASSWORD")
)
var myClient = &http.Client{Timeout: 10 * time.Second}

// GetLink gets link to be downlaoded
func GetLink() (bool, string) {

	bot, err := reddit.NewBotFromAgentFile("urlFetcher.agent", 0)
	if err != nil {
		print(err)
	}
	harvest, err := bot.Listing("/r/wallpapers", "")
	if err != nil {
		print(err)
	}
	if len(harvest.Posts) >= 1 {
		return true, harvest.Posts[0].URL
	}
	return false, ""

}
