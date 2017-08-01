package gopherize

import (
	"strings"
	"time"

	gopherizeme "github.com/matryer/gopherize.me/server"
	resty "gopkg.in/resty.v0"
)

type artworkResponse struct {
	Categories        []gopherizeme.Category `json:"categories"`
	TotalCombinations int64                  `json:"total_combinations"`
}

// <3 https://mholt.github.io/json-to-go/
type gopherResponse struct {
	ID           string    `json:"id"`
	Images       []string  `json:"images"`
	OriginalURL  string    `json:"original_url"`
	URL          string    `json:"url"`
	ThumbnailURL string    `json:"thumbnail_url"`
	Ctime        time.Time `json:"ctime"`
}

func init() {
	resty.SetRedirectPolicy(resty.FlexibleRedirectPolicy(5))
}

// GetAllCategories gives back all categories with their images
func GetAllCategories() (categories []gopherizeme.Category, err error) {
	response := artworkResponse{}
	_, err = resty.R().
		SetResult(&response).Get("https://gopherize.me/api/artwork/")
	categories = response.Categories

	return
}

func MapAllCategories() (categories map[string]gopherizeme.Category, err error) {
	categories = map[string]gopherizeme.Category{}

	var c []gopherizeme.Category
	c, err = GetAllCategories()

	for _, category := range c {
		categories[category.Name] = category
	}

	return
}

// ComposeImage gets the image URL of a set of images
func ComposeImage(input []string) (url string, err error) {
	get := strings.Join(input, "|")

	var response *resty.Response
	response, err = resty.R().Get("https://gopherize.me/save?images=" + get)
	url = response.RawResponse.Request.URL.String()

	return
}

// GetImageURL gets the image URL of your gopher!
func GetImageURL(gopherURL string) (url string, err error) {
	response := gopherResponse{}
	_, err = resty.R().
		SetResult(&response).Get(gopherURL + "/json")
	url = response.URL

	return
}
