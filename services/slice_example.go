package services

import (
	"github.com/WayneShenHH/toolsgo/tools"
)

func ExampleAppend() {
	ids := []interface{}{}
	ids = tools.UniqueAppend(ids, 1)
	ids = tools.UniqueAppend(ids, 1)
	ids = tools.UniqueAppend(ids, 1)
	ids = tools.UniqueAppend(ids, 2)
	ids = tools.UniqueAppend(ids, 3)
	ids = tools.UniqueAppend(ids, 3)
	ids = tools.UniqueAppend(ids, 4)
	ids = tools.UniqueAppend(ids, 4)
	tools.Log(ids)
}
