package main

import (
	"net/url"
	"os"
	"strings"

	"github.com/ChimeraCoder/anaconda"
	"github.com/int128/slack"
)

func main() {
	webhook := os.Getenv("MATTERMOST_WEBHOOK")
	consumerKey := os.Getenv("TWITER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITER_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("TWITER_ACCESS_TOKEN_SECRET")

	api := anaconda.NewTwitterApiWithCredentials(accessToken, accessTokenSecret, consumerKey, consumerSecret)
	v := url.Values{}
	v.Set("track", "#ieLT")
	s := api.PublicStreamFilter(v)
	for t := range s.C {
		switch v := t.(type) {
		case anaconda.Tweet:
			if strings.Index(v.Text, "#IELT") != -1 {
				continue
			}
			imgURL := strings.Replace(v.User.ProfileImageUrlHttps, "_normal.", ".", 1)
			go slack.Send(webhook, &slack.Message{
				Username: v.User.Name,
				IconURL:  imgURL,
				Text:     v.Text,
			})
		}
	}
	return
}
