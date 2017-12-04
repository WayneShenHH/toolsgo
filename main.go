package main

import (
	"fmt"
	"strconv"

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
	case "OfferMessage":
		// services.InsertMessage("worker:offer:message", "offer_ou")
		services.InsertMessage("worker:offer:message", "offer_point")
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
	case "ExampleValid":
		services.ExampleValid()
	case "CreateJuMatch":
		mid := tools.CmdParameters()[0]
		i, _ := strconv.Atoi(mid)
		services.CreateJuMatch(uint(i))
	default:
		fmt.Println("Task doesn't exist.")
	}
}
