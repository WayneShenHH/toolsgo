package services

import (
	"fmt"

	"github.com/WayneShenHH/toolsgo/tools"
)

func Sum(numbers ...int) {
	len := len(numbers)
	tools.MultiStartWithGroup(1000, func(i int) {
		sum := 0
		for _, v := range numbers[:len/2] {
			sum += v
		}
		fmt.Println(i, ":", sum)
	})
}
