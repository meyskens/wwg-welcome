package main

import "net/http"
import "os"

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
	}

	if sc.Token != token {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Invalid token"))
		return
	}

	w.Write([]byte(buildRandomGopher(sc.Text)))
}
