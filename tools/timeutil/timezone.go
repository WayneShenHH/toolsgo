package timeutil

import (
	"context"
	"fmt"
	"time"

	"github.com/jinzhu/now"
)

const (
	key = "localutc"
	// ServerTimeZone server time zone
	ServerTimeZone = 0.0
	// DefaultTimeZone client default time zone
	DefaultTimeZone = 8.0
)

// TimeZone local timezone info
type TimeZone struct {
	UTC float64
}

// Setter defines a context that enables setting values.
type Setter interface {
	Set(string, interface{})
}

// FromContext 從 context 中取得身份驗證資訊
func FromContext(c context.Context) (TimeZone, bool) {
	t, ok := c.Value(key).(TimeZone)
	return t, ok
}

// ToContext 將身份驗證資料傳入 context
func ToContext(c Setter, data TimeZone) {
	c.Set(key, data)
}

// FirstDayOfISOWeek get week day
func FirstDayOfISOWeek(year int, week int, zone *time.Location) time.Time {
	date := time.Date(year, 0, 0, 0, 0, 0, 0, zone)
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
		isoYear, isoWeek = date.ISOWeek()
	}

	return date
}

// TimeToString 時間轉字串的統一格式 UTC
func TimeToString(t time.Time) string {
	loc, _ := time.LoadLocation("UTC")
	return t.In(loc).Format("2006-01-02 15:04:05 +00:00")
}

// TimeToYMD 時間轉YYYY-MM-DD，UTC
func TimeToYMD(t time.Time) string {
	loc, _ := time.LoadLocation("UTC")
	return t.In(loc).Format("2006-01-02")
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
	layouts := []string{"2006-01-02T15:04:05.000Z", "2006-01-02 15:04:05 +00:00", "2006-01-02 15:04:05 UTC", "2006-01-02T15:04:05Z", "2006-01-02T15:04:05+00:00"}
	for _, layout := range layouts {
		ts, err = time.Parse(layout, timeString)
		if err == nil {
			return ts
		}
	}
	return ts
}

// TimeToDateString 時間轉MMDDHHmm，UTC
func TimeToDateString(t time.Time) string {
	loc, _ := time.LoadLocation("UTC")
	return t.In(loc).Format("01021504")
}

// JSONToUTCTime 將 json 格式的 time string convert to go time
func JSONToUTCTime(json string) time.Time {
	tpe, _ := time.LoadLocation("Asia/Taipei")
	layout := "2006-01-02T15:04:05.000Z"
	time, _ := time.Parse(layout, json)
	return time.In(tpe)
}

/*ReportDay 取得帳務日期
以中午為帳務日期開始的時間
*/
func ReportDay() (time.Time, time.Time) {
	midOfDay := now.BeginningOfDay().Add(12 * time.Hour)
	if time.Now().After(midOfDay) {
		return midOfDay, midOfDay.Add(24 * time.Hour)
	}
	return midOfDay.Add(-24 * time.Hour), midOfDay
}

/*ReportWeek 帳務週期
以星期一中午為帳務週期的開始時間
*/
func ReportWeek() (time.Time, time.Time) {
	startOfWeek := now.BeginningOfWeek().Add(36 * time.Hour)
	if time.Now().After(startOfWeek) {
		return startOfWeek, startOfWeek.AddDate(0, 0, 7)
	}
	return startOfWeek.AddDate(0, 0, -7), startOfWeek
}

// TimeToCST 轉CST時間
func TimeToCST(t time.Time) time.Time {
	loc, _ := time.LoadLocation("Asia/Taipei")
	return t.In(loc)
}

// Timer struct for timer
type Timer struct {
	cnt *time.Time
}

// GetTimer get a new timer
func GetTimer() *Timer {
	now := time.Now()
	return &Timer{
		cnt: &now,
	}
}

// Counting timer counting & reset
func (t *Timer) Counting(taskName string) time.Duration {
	duration := time.Since(*t.cnt)
	fmt.Printf("[%v] used time : %v\n", taskName, duration.String())
	t.Reset()
	return duration
}

// Reset timer reset
func (t *Timer) Reset() {
	*t.cnt = time.Now()
}
