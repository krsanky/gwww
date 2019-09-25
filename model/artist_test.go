package model

import (
	"fmt"
	"os"
	"testing"

	"oldcode.org/home/wise/repo/go/oldcode.org/gow/db"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/settings"
)

func TestMain(m *testing.M) {
	settings.Init("/home/wise/GO/gow/settings.toml")
	os.Exit(m.Run())
}

func HHHTestGetArtists(t *testing.T) {
	artists, err := GetArtists("A")
	if err != nil {
		t.Errorf("err:")
	}
	for _, a := range artists {
		fmt.Printf("a:%s\n", a.Name)
	}
}

func TestGetRawArtists(t *testing.T) {
	db.InitDB()
	as, err := GetRawArtists()
	if err != nil {
		t.Errorf("err:")
	}
	for _, a := range as {
		fmt.Printf("a:%s\n", a)
	}

}
