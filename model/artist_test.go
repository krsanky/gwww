package model

import (
	"fmt"
	"testing"
)

func TestGetArtists(t *testing.T) {

	//t.Errorf("asdasdasd")
	artists, err := GetArtists("A")
	if err != nil {
		t.Errorf("err:")
	}
	for _, a := range artists {
		fmt.Printf("a:%v\n", a)
	}
}
