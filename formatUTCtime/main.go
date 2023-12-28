package main

import (
	"fmt"
	"time"
)

const formatString = "2006-01-02T15:04:05"
const testTimeString = "2019-03-07T12:21:24"
const testTimeStringZ = "2019-03-07T12:21:24Z"
const testTimeFormat = "2006-01-02T15:04:05"
const testTimeFormatZ = "2006-01-02T15:04:05Z"

func main() {
	timeNow := time.Now()
	timeYesterday := timeNow.Add(time.Hour * -24)
	timeNowUTC := timeNow.UTC()
	timeYesterdayUTC := timeNowUTC.Add(time.Hour * -24)
	fmt.Printf("timeNow: %s\n", timeNow)
	fmt.Printf("timeYesterday: %s\n", timeYesterday)
	fmt.Printf("timeNowUTC: %s and %s and %s\n", timeNowUTC, timeNowUTC.Format(time.UnixDate), timeNowUTC.Format(formatString))
	fmt.Printf("timeYesterdayUTC: %s\n", timeYesterdayUTC)
	fmt.Println("___________________________________")

	theEpochSecs := timeNow.UnixMilli() / 1000
	theEpochSecsUTC := timeNowUTC.UnixMilli() / 1000
	timeNowSecs := time.UnixMilli(theEpochSecs*1000)
	isEqual := timeNowSecs.Equal(time.UnixMilli(theEpochSecs*1000))
	fmt.Printf("Equal? %t\n", isEqual)
	theTimeFromEpochSecs := time.UnixMilli(theEpochSecs * 1000)
	theTimeFromEpochSecsUTC := time.UnixMilli(theEpochSecsUTC * 1000).UTC()
	fmt.Printf("NONUTC: Now: %s, EpochSecs: %s, %d\n", timeNow, theTimeFromEpochSecs, theEpochSecs)
	fmt.Printf("   UTC: Now: %s, EpochSecs: %s, %d\n", timeNowUTC, theTimeFromEpochSecsUTC, theEpochSecsUTC)
	fmt.Println("___________________________________")

	tt(testTimeFormat, testTimeString)
	tt(testTimeFormat, testTimeStringZ)
	tt(testTimeFormatZ, testTimeString)
	tt(testTimeFormatZ, testTimeStringZ)
}

func tt(tformat, tstring string) {
	testTime, err := time.Parse(tformat, tstring)
	if err != nil {
		fmt.Printf("Parse error: %v\n", err)
	} else {
		fmt.Printf("Format: %s, String: %s, Parsed time: %s\n", tformat, tstring, testTime)
	}
}
