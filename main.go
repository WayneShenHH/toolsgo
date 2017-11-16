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
		// services.InsertMessage("worker:match:message", "match_ml")
		// services.InsertMessage("worker:match:message", "match_ou")
		services.InsertMessage("worker:match:message", "match_point")
		// services.InsertMessage("worker:match:message", "tmp")
	case "OfferMessage":
		// services.InsertMessage("worker:offer:message", "offer_ou")
		services.InsertMessage("worker:offer:message", "tmp")
	case "SpiderOffer":
		services.InsertSpiderOffer("2888802_ht_point_0.0_0", "spider_message")
	case "GoroutineWithWaitGroup":
		services.GoroutineWithWaitGroup()
	case "GoroutineExample":
		services.GoroutineExample()
	case "CheckStatus":
		services.CheckStatus()
	case "ExampleAppend":
		services.ExampleAppend()
	default:
		fmt.Println("Task doesn't exist.")
	}
}
