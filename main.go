package main

import (
	"fmt"
	"log"
	"net/http"

	gopherizeme "github.com/matryer/gopherize.me/server"
	"github.com/meyskens/wwg-welcome/gopherize"
)

var categories map[string]gopherizeme.Category // not that they are new images every day right?

func main() {
	var err error
	categories, err = gopherize.MapAllCategories()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/cmd", commandHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "This is the server that welcomes Gophers!")
	})
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
