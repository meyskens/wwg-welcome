package gopherize

import (
	"sync"
)

// Gopher is the gopher image to be generated
type Gopher struct {
	images      map[string]string
	imagesMutex sync.Mutex
}

// NewGopher gives a new gopher
func NewGopher() Gopher {
	return Gopher{
		imagesMutex: sync.Mutex{},
		images:      map[string]string{},
	}
}

func (g *Gopher) SetImage(name, id string) {
	g.imagesMutex.Lock()
	g.images[name] = id
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
