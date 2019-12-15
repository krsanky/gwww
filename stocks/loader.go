package stocks

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
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
	fmt.Printf("file:%v\n", file)
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
		fixed = append(fixed, s+"-123")
	}
	return fixed
}

func GetColumnNames() {
	/* Read 1 line from file:
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	line := scanner.Text()
	fmt.Printf("%s\n", line)
	// Symbol|Security Name|Market Category|Test Issue|Financial Status|Round Lot Size|ETF|NextShares
	*/

	r := csv.NewReader(file)
	r.Comma = '|'

	cols, err := r.Read()
	if err != nil {
		panic(err)
	}

	fixed := FixColumnNames(cols)
	fmt.Printf("col-1:%s col1-1:%s\n", fixed[0], cols[0])
	for i := 0; i < len(cols); i++ {
		fmt.Printf("cols:%s fixed:%s\n", cols[i], fixed[i])
	}
}

// depr.
func LoadFromFile(filename string) error {
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
