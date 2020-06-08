package timer

import (
	"time"
)

// FirstDayOfISOWeek get week day
func FirstDayOfISOWeek(year int, week int, timezone *time.Location) time.Time {
	date := time.Date(year, 0, 0, 0, 0, 0, 0, timezone)
	isoYear, isoWeek := date.ISOWeek()

	// iterate back to Monday
	for date.Weekday() != time.Monday {
		date = date.AddDate(0, 0, -1)
		isoYear, isoWeek = date.ISOWeek()
	}

	// iterate forward to the first day of the first week
	for isoYear < year {
		date = date.AddDate(0, 0, 7)
		isoYear, isoWeek = date.ISOWeek()
	}

	// iterate forward to the first day of the given week
	for isoWeek < week {
		date = date.AddDate(0, 0, 7)
		_, isoWeek = date.ISOWeek()
	}

	return date
}

// TimeToSQLString 將時間轉換為SQL格式
func TimeToSQLString(t time.Time) string {
	loc := time.UTC
	return t.In(loc).Format("2006-01-02 15:04:05")
}

// TimeToString 時間轉字串的統一格式 UTC
func TimeToString(t time.Time) string {
	loc := time.UTC
	return t.In(loc).Format("2006-01-02 15:04:05 +00:00")
}

// TimeToStamp 時間轉 timestamp 統一做法
func TimeToStamp(t time.Time) int64 {
	return t.Unix() * 1000
}

// StampToTime timestamp 轉回時間的統一做法
func StampToTime(stamp int64) time.Time {
	return time.Unix(stamp/1000, 0)
}

// StringToTime 轉回時間的統一做法 UTC
func StringToTime(timeString string) time.Time {
	var ts time.Time
	var err error
	layouts := []string{
		"2006-01-02T15:04:05.000Z",
		"2006-01-02 15:04:05 +00:00",
		"2006-01-02 15:04:05 UTC",
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04:05+00:00",
		"2006-01-02 15:04:05",
	}
	for _, layout := range layouts {
		ts, err = time.Parse(layout, timeString)
		if err == nil {
			return ts
		}
	}
	return ts
}

// ToAccountDate 轉換為帳務日期，需要傳入UTC時間
// 帳務日期回傳 UTC+20 的當天 00:00
func ToAccountDate(utc time.Time) time.Time {
	location := time.UTC
	gameUTC := utc.UTC().Add(20 * time.Hour)
	year, month, day := gameUTC.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, location)
}

// TimeToDateString 時間轉MMDDHHmm，UTC
func TimeToDateString(t time.Time) string {
	loc := time.UTC
	return t.In(loc).Format("01021504")
}

// TimeToYMD 時間轉YYYY-MM-DD，UTC
func TimeToYMD(t time.Time) string {
	loc := time.UTC
	return t.In(loc).Format("2006-01-02")
}

// TimeToFormat 時間轉指定格式，UTC
func TimeToFormat(t time.Time, format string) string {
	return t.In(time.UTC).Format(format)
}

// TimeToCST 轉CST時間
func TimeToCST(t time.Time) time.Time {
	location := GetTaipeiTimeZone()
	return t.In(location)
}

// StringToTWTime 字串轉換成台灣time.Time
func StringToTWTime(dateStr string) (time.Time, error) {
	var date time.Time
	var err error
	format := "2006-01-02-15"
	location := GetTaipeiTimeZone()
	if date, err = time.ParseInLocation(format, dateStr, location); err != nil {
		return date, err
	}
	return date, nil
}

//GetFirstDayOfYearMonth 取得特定月份的第一天(UTC)
func GetFirstDayOfYearMonth(year int, month time.Month) time.Time {
	location := time.UTC
	firstDayOfYearMonth := time.Date(year, month, 1, 0, 0, 0, 0, location)
	return firstDayOfYearMonth
}

//GetLastDayOfYearMonth 取得特定月份的最後一天(UTC)
func GetLastDayOfYearMonth(year int, month time.Month) time.Time {
	lastDayOfYearMonth := GetFirstDayOfYearMonth(year, month).AddDate(0, 1, 0).Add(-time.Nanosecond)
	return lastDayOfYearMonth
}
