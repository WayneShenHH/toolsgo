package sngheader

import (
	"testing"

	sng "github.com/gmallard/stompngo"
)

func TestMap(t *testing.T) {
	type args struct {
		m map[string]string
	}
	tests := []struct {
		name  string
		args  args
		wantH sng.Headers
	}{
		{
			args: args{map[string]string{
				"key1": "value1",
				"key2": "value2",
			}},
			wantH: sng.Headers{
				"key1", "value1",
				"key2", "value2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < len(tt.wantH); i += 2 {
				if value, exist := tt.args.m[tt.wantH[i]]; !exist {
					t.Errorf("%s not exist", value)
				}
			}
		})
	}
}
