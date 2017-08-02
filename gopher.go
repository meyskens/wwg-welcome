package main

import (
	"math/rand"

	"github.com/meyskens/wwg-welcome/gopherize"

	"fmt"
)

var buildorder = []string{ // or else it will end weird...
	"Body",
	"Eyes",
	"Shirts",
	"Hair",
	"Glasses",
	"Hats_and_Hair_Accessories",
	"Extras",
}

func buildRandomGopher(name string) string {
	rand.Seed(nameToSeed(name))

	gopher := gopherize.NewGopher()

	for _, b := range buildorder {
		category := categories[b]

		fmt.Println(category.Name)
		image := category.Images[rand.Intn(len(category.Images))]
		gopher.SetImage(image.ID)
	}

	url, err := gopher.GetImageURL()
	if err != nil {
		return "Sorry something bad happened"
	}
	return url
}

// hehehe
func nameToSeed(name string) int64 {
	bytes := []byte(name)

	var i int64
	for _, b := range bytes {
		i += int64(b)
	}

	return i
}
