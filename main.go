package main

import (
	"fmt"

	"wayne.sdk/wayne/toolsgo/services"
	"wayne.sdk/wayne/toolsgo/tools"
)

func main() {
	taskSelector()
}
func taskSelector() {
	switch tools.TaskName() {
	case "MatchMessage":
		services.InsertMessage("worker:match:message", "match_message")
	default:
		fmt.Println("Task doesn't exist.")
	}
}
