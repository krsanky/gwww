package model

import (
	"errors"
)

type Item struct {
	//gorm.Model
	ID          int `gorm:"primary_key"`
	Path        string
	Title       string
	Artist      string
	AlbumArtist string `gorm:"column:albumartist"`
	Track       int
	Media       string
}

func GetItems() ([]Item, error) {
	//	odb.Limit(15).Find(&items)
	//	data["items"] = items
	var items []Item

	return items, errors.New("GetItems() error...")
}

//CREATE TABLE items(
//  id INTEGER PRIMARY KEY,
//  path BLOB,
//  album_id INTEGER,
//  title TEXT,
//  artist TEXT,
//  artist_sort TEXT,
//  artist_credit TEXT,
//  album TEXT,
//  albumartist TEXT,
//  albumartist_sort TEXT,
//  albumartist_credit TEXT,
//  genre TEXT,
//  lyricist TEXT,
//  composer TEXT,
//  composer_sort TEXT,
//  arranger TEXT,
//  grouping TEXT,
//  year INTEGER,
//  month INTEGER,
//  day INTEGER,
//  track INTEGER,
//  tracktotal INTEGER,
//  disc INTEGER,
//  disctotal INTEGER,
//  lyrics TEXT,
//  comments TEXT,
//  bpm INTEGER,
//  comp INTEGER,
//  mb_trackid TEXT,
//  mb_albumid TEXT,
//  mb_artistid TEXT,
//  mb_albumartistid TEXT,
//  mb_releasetrackid TEXT,
//  albumtype TEXT,
//  label TEXT,
//  acoustid_fingerprint TEXT,
//  acoustid_id TEXT,
//  mb_releasegroupid TEXT,
//  asin TEXT,
//  catalognum TEXT,
//  script TEXT,
//  language TEXT,
//  country TEXT,
//  albumstatus TEXT,
//  media TEXT,
//  albumdisambig TEXT,
//  disctitle TEXT,
//  encoder TEXT,
//  rg_track_gain REAL,
//  rg_track_peak REAL,
//  rg_album_gain REAL,
//  rg_album_peak REAL,
//  r128_track_gain INTEGER,
//  r128_album_gain INTEGER,
//  original_year INTEGER,
//  original_month INTEGER,
//  original_day INTEGER,
//  initial_key TEXT,
//  length REAL,
//  bitrate INTEGER,
//  format TEXT,
//  samplerate INTEGER,
//  bitdepth INTEGER,
//  channels INTEGER,
//  mtime REAL,
//  added REAL
