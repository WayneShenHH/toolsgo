package tools

import (
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

const isLogMode = true

func Log(o ...interface{}) {
	if !isLogMode {
		return
	}
	br := getNewLine()
	for _, v := range o {
		j, _ := json.Marshal(v)
		fmt.Println(br)
		fmt.Println(string(j))
	}
	fmt.Println(br)
}
func getNewLine() string {
	fd := int(os.Stdout.Fd())
	termWidth, _, _ := terminal.GetSize(fd)
	br := ""
	for i := 0; i < termWidth; i++ {
		br += "-"
	}
	return br
}
