package model

import (
	"fmt"
	"net/url"
	"strings"

	"oldcode.org/home/wise/repo/go/oldcode.org/gow/db"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/lg"
)

type Artist struct {
	Name string
}

func T1() {
	fmt.Printf("db.TestDB() from model\n")
	db.TestDB()
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
WHERE albumartist like ?
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
//	rows, err := db.DB.Query(`
//SELECT DISTINCT albumartist 
//FROM albums 
//WHERE albumartist <> ''
//ORDER by albumartist
//`)
//	if err != nil {
//		return nil, err
//	}
//
//	var s string
//	artists := make([]string, 0)
//	for rows.Next() {
//		rows.Scan(&s)
//		artists = append(artists, s)
//	}
//	return artists, nil
//	db.TestDB()
	return nil, nil
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
	url := strings.Join([]string{"/artist?a=", name_part, ""}, "")
	return url
}

func (a *Artist) Albums() ([]Album, error) {
	lg.Log.Printf(".Albums() for %s", a.Name)
	albums := make([]Album, 0)

	rows, err := db.DBX.Queryx(`
SELECT id, album, albumartist
FROM albums
WHERE albumartist = ?
`, a.Name)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		a := Album{}
		err = rows.StructScan(&a)
		if err != nil {
			lg.Log.Printf("err:%s", err.Error())
		}
		lg.Log.Printf("Artist.Albums(): %d %s", a.ID, a.Title)
		albums = append(albums, a)
	}

	return albums, nil
}
