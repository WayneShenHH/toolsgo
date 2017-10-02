package services_test

import (
	"testing"

	"github.com/WayneShenHH/toolsgo/services"
)

func Test_LogModel(t *testing.T) {
	m := services.Message{}
	services.Log(m)
}
