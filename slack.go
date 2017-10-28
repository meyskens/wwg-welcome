package main

import (
	"encoding/json"
	"net/http"
	"os"

	resty "gopkg.in/resty.v0"
)

type slashCommand struct {
	Token       string
	TeamID      string
	TeamDomain  string
	ChannelID   string
	ChannelName string
	UserID      string
	UserName    string
	Command     string
	Text        string
	Hook        string
	ResponseURL string
}

type slackAttachment struct {
	Text     string `json:"text"`
	ImageURL string `json:"image_url"`
}
type slackResponse struct {
	ResponseType string            `json:"response_type,omitempty"`
	Text         string            `json:"text"`
	Attachments  []slackAttachment `json:"attachments"`
}

var token = os.Getenv("token")

func commandHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	v := r.Form

	sc := &slashCommand{
		Token:       v.Get("token"),
		TeamID:      v.Get("team_id"),
		TeamDomain:  v.Get("team_domain"),
		ChannelID:   v.Get("channel_id"),
		ChannelName: v.Get("channel_name"),
		UserID:      v.Get("user_id"),
		UserName:    v.Get("user_name"),
		Command:     v.Get("command"),
		Text:        v.Get("text"),
		Hook:        v.Get("hook"),
		ResponseURL: v.Get("response_url"),
	}

	if sc.Token != token {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Invalid token"))
		return
	}

	reply, _ := json.Marshal(slackResponse{
		ResponseType: "in_channel",
		Text:         "Generating Gopher...",
	})

	w.Header().Set("Content-Type", "application/json")
	w.Write(reply)
	go sendMessage(sc)
}

func sendMessage(sc *slashCommand) {
	url := buildRandomGopher(sc.Text)

	resty.R().SetBody(slackResponse{
		ResponseType: "in_channel",
		Text:         "Here is a Gopher I made using your input",
		Attachments: []slackAttachment{
			slackAttachment{
				ImageURL: url,
			},
			slackAttachment{
				Text: "Got the taste? Visit Gopherize.me",
			},
		},
	}).Post(sc.ResponseURL)
}
