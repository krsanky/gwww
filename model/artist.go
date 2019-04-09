package model

import (
	"oldcode.org/gow/db"
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
