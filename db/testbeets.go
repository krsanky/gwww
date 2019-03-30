package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"oldcode.org/gow/model"
)

func TestBeets() {
	fmt.Printf("test beets: %s\n", beets_db_file)
	db, err := gorm.Open("sqlite3", beets_db_file)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	if err = db.DB().Ping(); err != nil {
		panic("failed to ping database")
	}

	item := model.Item{}
	db.First(&item, 10)
	fmt.Printf("item:%s %d a:%s aa:%s\n", item.Title, item.ID, item.Artist, item.AlbumArtist)
}
