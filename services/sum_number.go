package services

import (
	"fmt"

	"github.com/WayneShenHH/toolsgo/tools"
)

func Sum(numbers ...int) {
	len := len(numbers)
	tools.MultiStart(2, func() {
		sum := 0
		for _, v := range numbers[:len/2] {
			sum += v
		}
		fmt.Println(sum)
	})
}
