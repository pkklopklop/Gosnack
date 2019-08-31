package main

import (
	"fmt"
	"time"

	"kbtg.tech/pbf-lib/ktime"
)

const (
	//Date time const
	DateTimeFormat                   = "2006-01-02T15:04:05.000+07:00"
	DateTimeFormat_YYYY_MM_DDThhmmss = "2006-01-02T15:04:05"
	DateTimeFormat_YYYYMMDDhhmmss    = "20060102150405" // format to YYYYMMDDhhmmss
	DateTimeFormat_YYYYMMDDhh        = "2006010215"     // format to YYYYMMDDhh
	DateTimeFormat_YYYYMMDD          = "20060102"       // format to YYYYMMDD
	DateTimeFormat_MMDD              = "0102"           // format to MMDD
	DateTimeFormat_YYYY_MM_DD        = "2006-01-02"     // format to YYYY-MM-DD
	DataTimeFormat_HH_MM_SS          = "15:04:05"       // format to HH:MM:SS
	DateTimeFormat_YYMMDD            = "060102"         // format to YYMMDD
	DataTimeFormat_HHMMSS            = "150405"         // format to HHMMSS
	DataTimeFormat_HHMM              = "1504"           // format to HHMM
)

func main() {

	//2019
	fmt.Println("2019 year next month-end")
	var firstDayOfTheMonth, currEndOfMonth time.Time
	for i := 1; i <= 12; i++ {
		firstDayOfTheMonth, _ = ktime.Date(2019, i, 1)
		currEndOfMonth = firstDayOfTheMonth.AddDate(0, 1, -1)
		fmt.Printf("Month %01d: %s\n", i, currEndOfMonth.Format(DateTimeFormat_YYYY_MM_DD))
	}

	//2019
	fmt.Println("2020 year next month-end")
	for i := 1; i <= 12; i++ {
		firstDayOfTheMonth, _ = ktime.Date(2020, i, 1)
		currEndOfMonth = firstDayOfTheMonth.AddDate(0, 1, -1)
		fmt.Printf("Month %01d: %s\n", i, currEndOfMonth.Format(DateTimeFormat_YYYY_MM_DD))
	}

}

//Input: date in year, month, day integer type.
//Output: date in time.Time type.
func date(year, month, day int) (time.Time, error) {
	terr := validateDate(year, month, day)
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC), terr
}

//Validate date if the date is valid (Ex. 2019-02-31 is not valid.)
func validateDate(yyyy int, mm int, dd int) error {
	_, terr := time.Parse(DateTimeFormat_YYYY_MM_DD, fmt.Sprintf("%04d-%02d-%02d", yyyy, mm, dd))
	return terr
}
