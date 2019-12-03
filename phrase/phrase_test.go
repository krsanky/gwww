package phrase

import (
	"fmt"
	"os"
	"testing"

	"oldcode.org/home/wise/repo/go/gow/db"
	"oldcode.org/home/wise/repo/go/gow/settings"
)

func TestMain(m *testing.M) {
	settings.Init("/home/wise/data/GO/gow/settings.toml")
	os.Exit(m.Run())
}

func TestGetPhrases(t *testing.T) {
	db.InitDB()
	ps, err := GetPhrases()
	if err != nil {
		t.Errorf("err:")
	}
	for _, p := range ps {
		fmt.Printf("p:%v\n", p)
	}

}
