package model

import (
	"fmt"

	"oldcode.org/home/wise/repo/go/oldcode.org/gow/db"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/lg"
)

type Album struct {
	ID          int
	Title       string `db:"album"`
	AlbumArtist string `db:"albumartist"`
	//  artpath BLOB, //sql.NullString
}

func (a *Album) String() string {
	return fmt.Sprintf("[Album %d %s %s]", a.ID, a.AlbumArtist, a.Title)
}

func (a *Album) Url() string {
	return fmt.Sprintf("/album?ai=%d", a.ID)
}

func AlbumByID(id int) (*Album, error) {
	lg.Log.Printf(".AlbumByID() for %d", id)
	var album Album
	err := db.BeetsDB.QueryRowx(`
SELECT id, album, albumartist
FROM albums
WHERE id = ?
`, id).StructScan(&album)
	if err != nil {
		return nil, err
	}
	lg.Log.Printf(".AlbumByID() album:%s", album.Title)
	return &album, nil
}

func (a *Album) Items() ([]Item, error) {
	items := make([]Item, 0)

	rows, err := db.BeetsDB.Queryx(`
SELECT
id, album_id, path, title,
artist, albumartist, track, media       
FROM items where album_id = ?
`, a.ID)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		i := Item{}
		err = rows.StructScan(&i)
		if err != nil {
			lg.Log.Printf("err:%s", err.Error())
		}
		lg.Log.Printf("album.Items(): %d %s", i.ID, i.Title)
		items = append(items, i)
	}

	return items, nil
}
