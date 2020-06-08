package timer

import (
	"context"
)

type key string

const tzkey key = "localutc"

// FromContext 從 context 中取得timer
func FromContext(c context.Context) Clock {
	tz := c.Value(string(tzkey))
	if tz == nil {
		tz = c.Value(tzkey)
	}
	if tz == nil {
		tz = New(DefaultTimezone)
	}
	return tz.(Clock)
}

// ToContext 將timer 傳入 context
func ToContext(c Setter, data Clock) {
	c.Set(string(tzkey), data)
}

// SetContext add store to context
func SetContext(c context.Context, data Clock) context.Context {
	return context.WithValue(c, tzkey, data)
}
