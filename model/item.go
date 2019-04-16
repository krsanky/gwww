package model

import "fmt"

type Item struct {
	ID          int
	AlbumID     int `db:"album_id"`
	Path        string
	Title       string
	Artist      string
	AlbumArtist string
	Track       int
	Media       string
}

func (i *Item) String() string {
	return fmt.Sprintf("[Item:%d ai:%d trk:%d %s :: %s]", i.ID, i.AlbumID, i.Track, i.Artist, i.Title)
}

func (i *Item) Url() string {
	return fmt.Sprintf("/track?i=%d", i.ID)
}
