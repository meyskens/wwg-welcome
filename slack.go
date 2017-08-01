package main

import (
	"encoding/json"
	"net/http"
	"os"

	resty "gopkg.in/resty.v0"
)

type slashCommand struct {
	Token       string
	TeamId      string
	TeamDomain  string
	ChannelId   string
	ChannelName string
	UserId      string
	UserName    string
	Command     string
	Text        string
	Hook        string
	ResponseURL string
}

type slackResponse struct {
	ResponseType string `json:"response_type,omitempty"`
	Text         string `json:"text"`
}

var token = os.Getenv("token")

func commandHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	v := r.Form

	sc := &slashCommand{
		Token:       v.Get("token"),
		TeamId:      v.Get("team_id"),
		TeamDomain:  v.Get("team_domain"),
		ChannelId:   v.Get("channel_id"),
		ChannelName: v.Get("channel_name"),
		UserId:      v.Get("user_id"),
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
	message := buildRandomGopher(sc.Text)

	resty.R().SetBody(slackResponse{
		Text: message,
	}).Post(sc.ResponseURL)
}
