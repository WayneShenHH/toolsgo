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
	case "InsertMessage":
		services.InsertMessage("worker:offer:message", "tmp")
	default:
		fmt.Println("Task doesn't exist.")
	}
}
