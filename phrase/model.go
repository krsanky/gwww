package phrase

import (
	"fmt"
	"strings"

	"oldcode.org/repo/go/gow/db"
	"oldcode.org/repo/go/gow/lg"
)

type Phrase struct {
	Id     int
	Phrase string
	Path   string
	Source string
	Tags   string
	Order  int `db:"order_"`
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
	return fmt.Sprintf("<Phrase id:%d path:%s tags:%s \"%s\">",
		p.Id, p.Path, p.Tags, tmpstr)
}

func (p *Phrase) Url() string {
	return fmt.Sprintf("/phrase/edit?p=%d", p.Id)
}

func (p *Phrase) Insert() error {
	lg.Log.Printf("INSERT phrase: %s", p.String())

	sql := `INSERT INTO phrase
(phrase, tags, path, order_, source)
VALUES ($1, $2, $3, $4, $5)`
	_, err := db.DBX.Exec(sql,
		p.Phrase, p.Tags, p.Path, p.Order, p.Source)

	return err
}

func GetPhrase(id int) (Phrase, error) {
	var p Phrase
	err := db.DBX.QueryRowx("SELECT * FROM phrase where id = $1", id).StructScan(&p)
	if err != nil {
		return p, err
	}
	return p, nil
}

func GetPhrases() ([]Phrase, error) {
	var ps []Phrase
	//db := db.DBX.Unsafe()
	db := db.DBX

	rows, err := db.Queryx("SELECT * FROM phrase")
	if err != nil {
		lg.Log.Printf("GetPhrases ERR:%s", err.Error())
		return ps, err
	}

	for rows.Next() {
		var p Phrase
		err = rows.StructScan(&p)
		if err != nil {
			lg.Log.Printf("GetPhrases ERR:%s", err.Error())
			return nil, err
		}
		ps = append(ps, p)
	}
	return ps, nil
}
