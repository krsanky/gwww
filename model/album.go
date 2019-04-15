package model

type Album struct {
	ID          int    `db:"id"`
	Album       string //sql.NullString
	AlbumArtist string `db:"albumartist"`
	//  artpath BLOB,
}

func (a Album) String() string {
	//return fmt.Sprint("%d %s %s", a.ID, a.Album, a.AlbumArtist)
	return a.Album
}
