package stocks

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
http://www.nasdaqtrader.com/trader.aspx?id=symboldirdefs
ftp://ftp.nasdaqtrader.com/symboldirectory/nasdaqlisted.txt
*/

var filename string
var columns []string
var file *os.File

func Init(filename string) {
	var err error
	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
}

func Cleanup() {
	err := file.Close()
	if err != nil {
		panic(err)
	}
}

func FixColumnNames(cols []string) []string {
	var fixed []string
	for _, s := range cols {
		tmp := strings.ToLower(s)
		tmp = strings.ReplaceAll(tmp, " ", "_")
		fixed = append(fixed, tmp)
	}
	return fixed
}

func GetColumnNames() {
	r := csv.NewReader(file)
	r.Comma = '|'

	cols, err := r.Read()
	if err != nil {
		panic(err)
	}

	columns := FixColumnNames(cols)
	for i := 0; i < len(cols); i++ {
		fmt.Printf("cols:%s columns[i]:%s\n", cols[i], columns[i])
	}
}

func ProcessNasdaqFile() error {
	r := csv.NewReader(file)
	r.Comma = '|'

	// throw away 1st line
	_, err := r.Read()
	if err != nil {
		panic(err)
	}

	for {
		r, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return(err)
		}
		if isLast(r[0]) {
			//fmt.Printf("LAST LINE--")
			break
		}
		//fmt.Printf("r0:%s\n", r[0])
		s := Stock{r[0], r[1]}
		err = s.Insert()
		if err != nil {
			panic(err)
		}
	}
	return nil	
}

func LoadFromFile2(filename string) error {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	//r := bufio.NewReader(f)

	scanner := bufio.NewScanner(f)
	run1 := true
	for scanner.Scan() {
		if run1 {
			fmt.Printf("<<<<<<%s>>>>>>\n", scanner.Text())
			run1 = false
			continue
		} else {
			break
		}
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}

	return nil
}

func isLast(line string) bool {
	return strings.HasPrefix(line, "File Creation Time:")
}
