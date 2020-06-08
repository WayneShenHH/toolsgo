// Package timer 系統使用取得時間介面
package timer

import (
	"fmt"
	"time"

	"github.com/WayneShenHH/toolsgo/pkg/module/logger"
)

const (
	// ServerTimezone server time zone
	ServerTimezone = 0.0
	// DefaultTimezone client default time zone
	DefaultTimezone = 8.0
)

// Clock local Clock info，將抓取現在時間的方法抽象化，以便測試程式可以指定當前時間為何
type Clock interface {
	UTC() float64
	Now() time.Time
	SetUTC(u float64)
	SetFixedNow(t time.Time)
	Counting(taskName string) time.Duration
	Reset()
	Today() time.Time
	GameToday() time.Time
	ReportWeekWithAccountDate() (time.Time, time.Time)
}

// Setter defines a context that enables setting values.
type Setter interface {
	Set(string, interface{})
}

type clock struct {
	utc float64
	cnt *time.Time
}

// Now custom time.Now
func (*clock) Now() time.Time {
	return time.Now()
}

// UTC get Clock UTC
func (c *clock) UTC() float64 {
	return c.utc
}

// SetUTC set Clock UTC
func (c *clock) SetUTC(u float64) {
	c.utc = u
}

// SetFixedNow
func (c *clock) SetFixedNow(t time.Time) {
	panic("shouldn't called by this implement")
}

// Counting timer counting & reset
func (c *clock) Counting(taskName string) time.Duration {
	duration := time.Since(*c.cnt)
	logger.Debug(fmt.Sprintf("[%v] used time : %v", taskName, duration.String()))
	c.Reset()
	return duration
}

// Reset timer reset
func (c *clock) Reset() {
	*c.cnt = time.Now()
}

// Today 取得 client 今日的 UTC 開始時間
func (c *clock) Today() time.Time {
	location := time.UTC
	year, month, day := c.Now().UTC().Date()
	return time.Date(year, month, day, 0, 0, 0, 0, location).Add(time.Hour * -time.Duration(int(c.UTC())))
}

// GameToday 目前的 Account Date 日期 (UTC+20)
// 如果是 UTC 時間開始時間應該是多少
func (c *clock) GameToday() time.Time {
	return ToAccountDate(c.Now()).Add(time.Hour * +time.Duration(4))
}

/*ReportWeekWithAccountDate 帳務週期:
以台灣時間星期一中午做為帳務週期的開始時間 ex:2018-04-23 00:00:00 +0000 UTC ~ 2018-04-30 00:00:00 +0000 UTC.
台灣位於UTC+8時區,中午12點相當於UTC+20時區的凌晨0點。再將帳務週期從日～六,轉為一～日。
*/
func (c *clock) ReportWeekWithAccountDate() (time.Time, time.Time) {
	startOfWeek := c.beginningOfWeek().Add(24 * time.Hour) //先取開始日,此套件的Week從禮拜日開始,這邊轉換成禮拜一開始計算
	startOfWeek = ToAccountDate(startOfWeek)               //轉換帳務日期基礎到台灣時間12點整
	if time.Now().After(startOfWeek) {
		return startOfWeek, startOfWeek.AddDate(0, 0, 7)
	}
	return startOfWeek.AddDate(0, 0, -7), startOfWeek
}

// beginningOfDay beginning of day
func (c *clock) beginningOfDay() time.Time {
	location := time.UTC
	y, m, d := c.Now().Date()
	return time.Date(y, m, d, 0, 0, 0, 0, location)
}

// beginningOfWeek beginning of week, sunday is beginning
func (c *clock) beginningOfWeek() time.Time {
	t := c.beginningOfDay()
	weekday := int(t.Weekday())
	return t.AddDate(0, 0, -weekday)
}

// New init clock
func New(tzUTC float64) Clock {
	now := time.Now()
	return &clock{
		utc: tzUTC,
		cnt: &now,
	}
}

// GetTaipeiTimeZone 取得台北時區
func GetTaipeiTimeZone() *time.Location {
	zone := time.FixedZone("Asia/Taipei", 8*60*60)
	return zone
}
