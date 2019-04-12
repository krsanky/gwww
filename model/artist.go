package model

import (
	"net/url"
	"strings"

	"oldcode.org/gow/db"
	"oldcode.org/gow/lg"
)

type Artist struct {
	Name string
}

func GetArtists() ([]Artist, error) {
	artists := make([]Artist, 0)
	artists_, err := db.GetRawArtists()
	if err != nil {
		return nil, err
	}

	for _, a := range artists_ {
		artists = append(artists, Artist{Name: a})
	}

	return artists, nil
}

func (a *Artist) Url() string {
	name_part := url.QueryEscape(a.Name)
	url := strings.Join([]string{"/artist/", name_part, "/"}, "")
	return url
}

func (a *Artist) Albums() {
	lg.Log.Printf(".Albums() for %s", a.Name)
}
