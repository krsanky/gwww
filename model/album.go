package model

import (
	"fmt"

	"oldcode.org/home/wise/repo/go/gow/db"
	"oldcode.org/home/wise/repo/go/gow/lg"
)

type Album struct {
	Id          int
	Title       string `db:"album"`
	AlbumArtist string `db:"albumartist"`
	//  artpath BLOB, //sql.NullString
}

func (a *Album) String() string {
	return fmt.Sprintf("[Album %d %s %s]", a.Id, a.AlbumArtist, a.Title)
}

func (a *Album) Url() string {
	return fmt.Sprintf("/music/album?ai=%d", a.Id)
}

func AlbumById(id int) (*Album, error) {
	lg.Log.Printf(".AlbumById() for %d", id)
	var album Album
	err := db.DBX.QueryRowx(`
SELECT id, album, albumartist
FROM albums
WHERE id = $1
`, id).StructScan(&album)
	if err != nil {
		return nil, err
	}
	lg.Log.Printf(".AlbumById() album:%s", album.Title)
	return &album, nil
}

func (a *Album) Items() ([]Item, error) {
	items := make([]Item, 0)

	rows, err := db.DBX.Queryx(`
SELECT
id, album_id, path, title,
artist, albumartist, track, media       
FROM items where album_id = $1
ORDER by track
`, a.Id)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		i := Item{}
		err = rows.StructScan(&i)
		if err != nil {
			lg.Log.Printf("err:%s", err.Error())
		}
		lg.Log.Printf("album.Items(): %d %s", i.Id, i.Title)
		items = append(items, i)
	}

	return items, nil
}

func Albums(artist string) ([]Album, error) {
	lg.Log.Printf(".Albums() for %s", artist)
	albums := make([]Album, 0)

	rows, err := db.DBX.Queryx(`
SELECT id, album, albumartist
FROM albums
WHERE albumartist = $1
`, artist)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		a := Album{}
		err = rows.StructScan(&a)
		if err != nil {
			lg.Log.Printf("err:%s", err.Error())
		}
		lg.Log.Printf("Artist.Albums(): %d %s", a.Id, a.Title)
		albums = append(albums, a)
	}

	return albums, nil
}
