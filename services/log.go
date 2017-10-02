package services

import (
	"encoding/json"
	"fmt"
)

func Log(o interface{}) {
	j, _ := json.Marshal(o)
	fmt.Println(string(j))
}
