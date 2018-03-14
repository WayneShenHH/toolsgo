package services

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	input := 8
	for i := 0; i < input; i++ {
		for j := 0; j < 2*input; j++ {
			if j < (input-i) || j > (input+i) {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Print("\n")
	}
}
