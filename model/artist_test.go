package model

import (
	"fmt"
	"os"
	"testing"

	"oldcode.org/repo/go/gow/db"
	"oldcode.org/repo/go/gow/settings"
)

func TestMain(m *testing.M) {
	settings.Init("/home/wise/data/GO/gow/settings.toml")
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
	db.Init()
	as, err := GetRawArtists()
	if err != nil {
		t.Errorf("err:")
	}
	for _, a := range as {
		fmt.Printf("a:%s\n", a)
	}

}
