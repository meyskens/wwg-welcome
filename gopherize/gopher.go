package gopherize

import (
	"sync"
)

// Gopher is the gopher image to be generated
type Gopher struct {
	images      []string
	imagesMutex sync.Mutex
}

// NewGopher gives a new gopher
func NewGopher() Gopher {
	return Gopher{
		imagesMutex: sync.Mutex{},
		images:      []string{},
	}
}

func (g *Gopher) SetImage(id string) {
	g.imagesMutex.Lock()
	g.images = append(g.images, id)
	g.imagesMutex.Unlock()
}

func (g *Gopher) GetImageURL() (string, error) {
	i := []string{}
	for _, id := range g.images {
		i = append(i, id)
	}
	gopherURL, err := ComposeImage(i)
	if err != nil {
		return "", err
	}
	return GetImageURL(gopherURL)
}
