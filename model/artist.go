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
	url := strings.Join([]string{"/artist?a=", name_part, ""}, "")
	return url
}

func (a *Artist) Albums() ([]Album, error) {
	lg.Log.Printf(".Albums() for %s", a.Name)
	albums := make([]Album, 0)

	db.Open()
	rows, err := db.DB.Queryx(`
SELECT ID, Album, AlbumArtist
FROM albums
WHERE albumartist = ?
`, a.Name)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var a Album
		err = rows.StructScan(&a)
		if err != nil {
			lg.Log.Printf("err:%s", err.Error())
		}
		lg.Log.Printf("Artist.Albums(): %d %s", a.ID, a.Album)
		albums = append(albums, a)
	}

	return albums, nil
}



