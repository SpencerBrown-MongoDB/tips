package dates

import (
	"fmt"
	"strconv"
	"time"
)

const (
	parseDate      = "2006-01-02"
	parseTimeHours = "2006-01-02T15"
	parseTimeMins  = "2006-01-02T15:04"
	parseTimeSecs  = "2006-01-02T15:04:05"
)

// ShowDate formats a time to YYYY-MM-DD format. It will be UTC.
func ShowDate(t time.Time) string {
	return t.Format(parseDate)
}

// ShowTime formats a time to HH:MM:SS format. It will be UTC.
func ShowTime(t time.Time) string {
	return t.Format(parseTimeSecs)[11:]
}

// ConvertSimpleTimeString converts a time string like yyyy-mm-ddThh:mm:ss, assumed to be UTC,
// to its time.Time equivalent
func ConvertSimpleTimeString(ts string) (time.Time, error) {
	return time.Parse(parseTimeSecs, ts)
}

// ...strictly for test purposes...
var areTesting bool

const testYear = "2023"
const testTimeNow = 1701388800 * 1000

// ............................

const TimeStringHelp = "The string representation is yyyy-mm-ddThh:mm:ss and it is assumed to be UTC." +
	" The ss, mm:ss, or Thh:mm:ss are optional and default to zero." +
	" The yyyy- is also optional and defaults to this year."

// ConvertTimeString converts a string representation of a timestamp into a UTC time and a Unix epoch time in seconds.
// See "TimeStringHelp" above for the string representation format.
// if the arg is empty, the return will be the time now less the backoff value in seconds
func ConvertTimeString(arg string, backoff int64) (time.Time, int64, error) {
	theTime := time.Now().UTC()
	var theTimeMillis int64
	var err error
	// default to now minus backoff if arg is empty
	if arg == "" {
		theTimeMillis = theTime.UnixMilli()
		if areTesting { // just for test purposes
			theTimeMillis = testTimeNow
		}
		theTimeMillis -= backoff * 1000
		theTime = time.UnixMilli(theTimeMillis).UTC()
		return theTime, theTimeMillis / 1000, nil
	}
	// default the year if it's not there
	if len(arg) > 4 && arg[4] != '-' {
		if areTesting { // just for test purposes
			arg = testYear + "-" + arg
		} else {
			thisYear := strconv.Itoa(theTime.Year())
			arg = thisYear + "-" + arg
		}
	}
	theTime, err = time.Parse(parseTimeSecs, arg)
	if err != nil {
		theTime, err = time.Parse(parseTimeMins, arg)
		if err != nil {
			theTime, err = time.Parse(parseTimeHours, arg)
			if err != nil {
				theTime, err = time.Parse(parseDate, arg)
				if err != nil {
					return time.Time{}, 0, fmt.Errorf("invalid date/time format. should be yyyy-dd-mmThh:mm:ss, where the ss, mm, hh, or the entire T portion may be omitted. yyyy- may be omitted and defaults to this year. UTC is assumed.")
				}
			}
		}
	}
	// Note: time.Parse returns a UTC time when neither the format string nor the string to parse has a zone indicator.
	// It does not default to the local time zone, unlike time.Now().
	// round down to nearest second
	theEpochSecs := theTime.UnixMilli() / 1000
	return time.UnixMilli(theEpochSecs * 1000).UTC(), theEpochSecs, nil
}

// getTimesFromQuery gets start and end times as Unix epoch values
// and returns them as time.Time with appropriate defaults of end = now and start = now - backoff
func getTimesFromQuery(startTimeString string, endTimeString string, backoff int) (time.Time, time.Time, error) {
	startTime, err := getTimeFromEpochString(startTimeString)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid start time: %v", err)
	}
	endTime, err := getTimeFromEpochString(endTimeString)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid end time: %v", err)
	}
	if endTime.IsZero() {
		endTime = time.Now().UTC()
	}
	if startTime.IsZero() {
		startTime = endTime.Add(time.Duration(-backoff) * time.Second)
	}
	return startTime, endTime, nil
}

// getTimeFromEpochString converts a string with the Unix epoch in seconds
// into a time.Time. If the string is empty or "0", returns a zero time.
func getTimeFromEpochString(epochString string) (time.Time, error) {
	if epochString == "" {
		return time.Time{}, nil
	}
	epochVal, err := strconv.ParseInt(epochString, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	if epochVal == 0 {
		return time.Time{}, nil
	}
	if epochVal < 0 {
		return time.Time{}, fmt.Errorf("negative epoch %s", epochString)
	}
	return time.UnixMilli(epochVal * 1000).UTC(), nil
}
