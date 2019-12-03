package model

import (
	"fmt"
	"net/url"
	"strings"

	"oldcode.org/repo/go/gow/db"
)

type Artist struct {
	Name string
}

func GetArtists(startswith string) ([]Artist, error) {
	if strings.Compare("", startswith) == 0 {
		startswith = "%"
	} else {
		startswith = startswith + "%"
	}
	rows, err := db.DB.Query(`
SELECT DISTINCT albumartist 
FROM albums 
WHERE albumartist like $1
ORDER by albumartist
`, startswith)
	if err != nil {
		return nil, err
	}

	var a string
	artists := make([]Artist, 0)
	for rows.Next() {
		rows.Scan(&a)
		artists = append(artists, Artist{a})
	}
	return artists, nil
}

func GetRawArtists() ([]string, error) {
	rows, err := db.DB.Query(`
SELECT DISTINCT albumartist 
FROM albums 
WHERE albumartist <> ''
ORDER by albumartist
`)
	if err != nil {
		return nil, err
	}

	var s string
	artists := make([]string, 0)
	for rows.Next() {
		rows.Scan(&s)
		artists = append(artists, s)
	}
	return artists, nil
}

func GetAllArtists() ([]Artist, error) {
	artists := make([]Artist, 0)
	artists_, err := GetRawArtists()
	if err != nil {
		return nil, err
	}

	for _, a := range artists_ {
		artists = append(artists, Artist{Name: a})
	}

	return artists, nil
}

func (a *Artist) String() string {
	return fmt.Sprintf("<artist name:%s>", a.Name)
}

func (a *Artist) Url() string {
	name_part := url.QueryEscape(a.Name)
	url := strings.Join([]string{"/music/artist?a=", name_part, ""}, "")
	return url
}
