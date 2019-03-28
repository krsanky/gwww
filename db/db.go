package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var beets_db_file = "/home/wise/go/src/oldcode.org/gow/beets.db"

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func TestSql() {
	fmt.Printf("test (sql) beets: %s\n", beets_db_file)
	db, err := gorm.Open("sqlite3", beets_db_file)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	//var ODB *sql.DB
	ODB := db.DB()
	if err = ODB.Ping(); err != nil {
		panic("failed to ping database")
	}

	type Result struct {
		Id    int
		Title string
	}
	var result Result
	db.Raw("SELECT id, title from items where id = 10").Scan(&result)
	fmt.Printf("id:%d title:%s\n", result.Id, result.Title)
}

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

	item := Item{}
	db.First(&item, 10)
	fmt.Printf("item:%s %d a:%s aa:%s\n", item.Title, item.ID, item.Artist, item.AlbumArtist)
}

func Main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})

	// Read
	var product Product
	db.First(&product, 1)                   // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	db.Delete(&product)
}
