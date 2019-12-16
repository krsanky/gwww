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

Symbol  The one to four or five character identifier for each NASDAQ-listed security.

Security Name  Company issuing the security.

Market Category The category assigned to the issue by NASDAQ based
on Listing Requirements.  Values:
    Q = NASDAQ Global Select MarketSM
    G = NASDAQ Global MarketSM
    S = NASDAQ Capital Market

Test Issue  Indicates whether or not the security is a test security. Values: Y = yes, it is a test issue. N = no, it is not a test issue.

Financial Status Indicates when an issuer has failed to submit its
regulatory filings on a timely basis, has failed to meet NASDAQ's
continuing listing standards, and/or has filed for bankruptcy.
Values include:
    D = Deficient: Issuer Failed to Meet NASDAQ Continued Listing Requirements
    E = Delinquent: Issuer Missed Regulatory Filing Deadline
    Q = Bankrupt: Issuer Has Filed for Bankruptcy
    N = Normal (Default): Issuer Is NOT Deficient, Delinquent, or Bankrupt.
    G = Deficient and Bankrupt
    H = Deficient and Delinquent
    J = Delinquent and Bankrupt
    K = Deficient, Delinquent, and Bankrupt

Round Lot  Indicates the number of shares that make up a round lot for the given security.

File Creation Time: The last row of each Symbol Directory text file
contains a timestamp that reports the File Creation Time.  The file
creation time is based on when NASDAQ Trader generates the file and
can be used to determine the timeliness of the associated data.
The row contains the words File Creation Time followed by mmddyyyyhhmm
as the first field, followed by all delimiters to round out the
row.  An example: File Creation Time: 1217200717:03|||||

ZUMZ|Zumiez Inc. - Common Stock|Q|N|N|100|N|N
ZVO|Zovio Inc. - Common Stock|Q|N|N|100|N|N
ZVZZC|NASDAQ TEST STOCK Nextshares Test Security|G|Y|N|100||Y
ZVZZT|NASDAQ TEST STOCK|G|Y|N|100||N
ZWZZT|NASDAQ TEST STOCK|S|Y|N|100||N
ZXYZ.A|Nasdaq Symbology Test Common Stock|Q|Y|N|100||N
ZXZZT|NASDAQ TEST STOCK|G|Y|N|100||N
ZYNE|Zynerba Pharmaceuticals, Inc. - Common Stock|G|N|N|100|N|N
ZYXI|Zynex, Inc. - Common Stock|S|N|N|100|N|N
File Creation Time: 1213201911:01|||||||
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
			fmt.Printf("LAST LINE--")
			break
		}
		fmt.Printf("r0:%s\n", r[0])
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
