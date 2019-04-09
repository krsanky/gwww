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

type Artist struct {
	Name string
}

func GetItems() ([]Item, error) {
	//	odb.Limit(15).Find(&items)
	//	data["items"] = items
	var items []Item

	return items, errors.New("GetItems() error...")
}


func GetArtists() ([]Artist, error) {
	//	odb := db.GetOpenDB()
	//	defer odb.Close()
	//	//	odb.Limit(15).Find(&items)
	//	//	data["items"] = items
	//	var artists []Artist
	//
	//	odb.Limit(15).Find(&items)
	//
	//	return artists, errors.New("GetArtists() error...")
	return nil, nil
}

func (a Artist) GetAlbums() {
}

//Limit
//Specify the max number of records to retrieve
//
//db.Limit(3).Find(&users)
////// SELECT * FROM users LIMIT 3;
//
//// Cancel limit condition with -1
//db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
////// SELECT * FROM users LIMIT 10; (users1)
////// SELECT * FROM users; (users2)

//Offset
//Specify the number of records to skip before starting to return the records
//
//db.Offset(3).Find(&users)
////// SELECT * FROM users OFFSET 3;
//
//// Cancel offset condition with -1
//db.Offset(10).Find(&users1).Offset(-1).Find(&users2)
////// SELECT * FROM users OFFSET 10; (users1)
////// SELECT * FROM users; (users2)

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
