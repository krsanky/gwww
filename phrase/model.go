package phrase

import (
	"fmt"
	"strings"

	"oldcode.org/home/wise/repo/go/gow/db"
	"oldcode.org/home/wise/repo/go/gow/lg"
)

type Phrase struct {
	Id     int
	Phrase string
	Path   string
	Tags   string
	Order  int `schema:"order_"`
}

func (p *Phrase) Str() string {
	return p.String()
}
func (p *Phrase) String() string {
	var tmpstr = p.Phrase
	if len(p.Phrase) > 9 {
		tmpstr = p.Phrase[:9] + "..."
	}
	idx := strings.Index(tmpstr, "\n")
	if (idx != -1) && (idx > 0) {
		tmpstr = tmpstr[:idx-1]
	}
	return fmt.Sprintf("<Phrase id:%d \"%s\">", p.Id, tmpstr)
}

func (p *Phrase) Insert() error {
	lg.Log.Printf("INSERT phrase ...")

	sql := `INSERT INTO phrase
(phrase, tags, path, order_)
VALUES ($1, $2, $3, $4)`
	_, err := db.DBX.Exec(sql,
		p.Phrase, p.Tags, p.Path, p.Order)

	return err
}

func GetPhrases() ([]Phrase, error) {
	var ps []Phrase
	db := db.DBX.Unsafe()

	rows, err := db.Queryx("SELECT * FROM phrase")
	if err != nil {
		return ps, err
	}

	for rows.Next() {
		var p Phrase
		err = rows.StructScan(&p)
		if err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}
	return ps, nil
}
