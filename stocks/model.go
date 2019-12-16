package stocks

import (
	"fmt"

	"oldcode.org/repo/go/gow/db"
	"oldcode.org/repo/go/gow/lg"
)

type Stock struct {
	Symbol string
	Name   string
}

func (s *Stock) String() string {
	return fmt.Sprintf("<stock %s>", s.Symbol)
}

func (s *Stock) Insert() error {
	sql := `INSERT INTO stock (symbol, name) VALUES ($1, $2)`
	lg.Log.Printf("stocks.Insert().......")
	_, err := db.DBX.Exec(sql, s.Symbol, s.Name)
	return err
}

func GetAll() ([]Stock, error) {
	var ss []Stock
	db := db.DBX

	rows, err := db.Queryx("SELECT * FROM stock")
	if err != nil {
		lg.Log.Printf("stocks.GetAll ERR:%s", err.Error())
		return ss, err
	}

	for rows.Next() {
		var s Stock
		err = rows.StructScan(&s)
		if err != nil {
			lg.Log.Printf("stocks.GetAll ERR:%s", err.Error())
			return nil, err
		}
		ss = append(ss, s)
	}
	return ss, nil
}
