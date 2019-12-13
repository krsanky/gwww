package stocks

import (
	"bufio"
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

func LoadFromFile(filename string) error {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	//r := bufio.NewReader(f)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Printf("<<<<<<%s>>>>>>\n", scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}

	return nil
}
