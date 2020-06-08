package timer

import (
	"reflect"
	"testing"
	"time"
)

type testUnit struct {
	name string
	args args
	want time.Time
}
type args struct {
	utc time.Time
}

func TestToAccountTime(t *testing.T) {

	tests := []testUnit{
		// {name: "", args: args{utc:}},
	}
	timeFormat := "2006-01-02 15:04:05 +0000"
	gameTime, _ := time.Parse(timeFormat, "2018-04-18 13:15:22 +0800")
	accountDate, _ := time.Parse(timeFormat, "2018-04-19 00:00:00")
	tests = append(tests, testUnit{name: "04/18 下午開賽帳務日期為 04/19", args: args{utc: gameTime}, want: accountDate})

	gameTime, _ = time.Parse(timeFormat, "2018-04-18 11:15:22 +0800")
	accountDate, _ = time.Parse(timeFormat, "2018-04-18 00:00:00")
	tests = append(tests, testUnit{name: "04/18 上午開賽帳務日期為 04/18", args: args{utc: gameTime}, want: accountDate})

	gameTime, _ = time.Parse(timeFormat, "2018-04-18 12:00:00 +0800")
	accountDate, _ = time.Parse(timeFormat, "2018-04-19 00:00:00")
	tests = append(tests, testUnit{name: "04/18 中午整點開賽帳務日期為 04/19", args: args{utc: gameTime}, want: accountDate})

	for idx := range tests {
		tt := tests[idx]
		t.Run(tt.name, func(t *testing.T) {
			if got := ToAccountDate(tt.args.utc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToAccountTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
