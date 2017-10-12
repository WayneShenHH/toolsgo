package main

import (
	"fmt"

	"github.com/WayneShenHH/toolsgo/services"
	"github.com/WayneShenHH/toolsgo/tools"
)

func main() {
	taskSelector()
}
func taskSelector() {
	switch tools.TaskName() {
	case "MatchMessage":
		services.InsertMessage("worker:match:message", "match_point")
		services.InsertMessage("worker:match:message", "match_ou")
	case "OfferMessage":
		services.InsertMessage("worker:offer:message", "offer_ou")
	case "GoroutineWithWaitGroup":
		services.GoroutineWithWaitGroup()
	case "GoroutineExample":
		services.GoroutineExample()
	default:
		fmt.Println("Task doesn't exist.")
	}
}
