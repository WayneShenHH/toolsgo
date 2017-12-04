package tools

import "os"
import "fmt"

func TaskName() string {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("There is no arg in command.")
		return ""
	}
	return args[0]
}

func CmdParameters() []string {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("There is no arg in command.")
		return nil
	}
	return args[1:]
}
