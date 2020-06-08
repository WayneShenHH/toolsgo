// Package mocktimer mock timer
package mocktimer

import (
	"fmt"
	"time"

	"github.com/WayneShenHH/toolsgo/module/logger"
	"github.com/WayneShenHH/toolsgo/module/timer"
)

type clock struct {
	utc      float64
	fixednow time.Time
	cnt      *time.Time
}

// New 初始
func New() timer.Clock {
	return &clock{}
}

// UTC 時區
func (c *clock) UTC() float64 {
	return c.utc
}

func (c *clock) Now() time.Time {
	return c.fixednow
}

// SetUTC set Clock UTC
func (c *clock) SetUTC(u float64) {
	c.utc = u
}

// SetFixedNow 設定預期的當前系統時間
func (c *clock) SetFixedNow(t time.Time) {
	c.fixednow = t
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
	location, _ := time.LoadLocation("UTC")
	year, month, day := c.Now().UTC().Date()
	return time.Date(year, month, day, 0, 0, 0, 0, location).Add(time.Hour * -time.Duration(int(c.UTC())))
}

// GameToday 目前的 Account Date 日期 (UTC+20)
// 如果是 UTC 時間開始時間應該是多少
func (c *clock) GameToday() time.Time {
	return timer.ToAccountDate(c.Now()).Add(time.Hour * +time.Duration(4))
}

/*ReportWeekWithAccountDate 帳務週期:
以台灣時間星期一中午做為帳務週期的開始時間 ex:2018-04-23 00:00:00 +0000 UTC ~ 2018-04-30 00:00:00 +0000 UTC.
台灣位於UTC+8時區,中午12點相當於UTC+20時區的凌晨0點。再將帳務週期從日～六,轉為一～日。
*/
func (c *clock) ReportWeekWithAccountDate() (time.Time, time.Time) {
	startOfWeek := c.beginningOfWeek().Add(24 * time.Hour) //先取開始日,此套件的Week從禮拜日開始,這邊轉換成禮拜一開始計算
	startOfWeek = timer.ToAccountDate(startOfWeek)         //轉換帳務日期基礎到台灣時間12點整
	if time.Now().After(startOfWeek) {
		return startOfWeek, startOfWeek.AddDate(0, 0, 7)
	}
	return startOfWeek.AddDate(0, 0, -7), startOfWeek
}

// beginningOfDay beginning of day
func (c *clock) beginningOfDay() time.Time {
	location, _ := time.LoadLocation("UTC")
	y, m, d := c.Now().Date()
	return time.Date(y, m, d, 0, 0, 0, 0, location)
}

// beginningOfWeek beginning of week, sunday is beginning
func (c *clock) beginningOfWeek() time.Time {
	t := c.beginningOfDay()
	weekday := int(t.Weekday())
	return t.AddDate(0, 0, -weekday)
}
