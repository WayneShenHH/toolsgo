// Package oddsupdate ws handler
package oddsupdate

import (
	"encoding/json"
	"fmt"

	"github.com/WayneShenHH/toolsgo/pkg/module/logger"
	"github.com/WayneShenHH/toolsgo/pkg/module/mq"
	"github.com/WayneShenHH/toolsgo/pkg/module/mq/topic"
	"github.com/WayneShenHH/toolsgo/pkg/transport/websocket"
)

// Hub oddsupdate hub, package again for wire can tell different ws handler
type Hub interface {
	websocket.Hub
}

// SubMsg 訂閱資訊
// swagger:model SampleSubMsg
type SubMsg struct {
	Indexes []SubIndex
}

// SubIndex 用於篩選的索引條件
type SubIndex struct {
	//篩選的編號
	ID uint
}

// NewHub oddsupdate hub
func NewHub(mq mq.MessageQueueService) Hub {
	return websocket.NewHub(mq, topic.UMSOddsUpdate, topic.DefaultChannel, getSubscribeRooms, getBroadcastRooms)
}

func getSubscribeRooms(msg []byte) []string {
	var sub SubMsg
	err := json.Unmarshal(msg, &sub)
	if err != nil {
		logger.Errorf(`sample/getSubscribeRooms error:%v`, err)
		return nil
	}
	rooms := []string{}
	for i := range sub.Indexes {
		name := getRoomName(sub.Indexes[i].ID)
		rooms = append(rooms, name)
	}
	return rooms
}

// BroadcastSampleMessage sample
// swagger:model BroadcastSampleMessage
type BroadcastSampleMessage struct {
	ID      uint
	Message string
}

func getBroadcastRooms(msg []byte) []string {
	m := &broadcastMessage{}
	err := json.Unmarshal(msg, m)
	if err != nil {
		logger.Error(err)
		return nil
	}

	rooms := []string{
		getRoomName(m.ID),
	}
	return rooms
}

func getRoomName(id uint) string {
	return fmt.Sprintf("room_%v", id)
}

// swagger:operation POST /trader/odds websocket broadcastodds
//
// 即時賠率更新訊息，Base URL : /ums/ws
//
//
// ---
// security:
// - bearer: []
// parameters:
// - name: Authorization
//   in: query
//   type: string
//   description: 身份驗證token
// - name: SubMsg
//   in: body
//   description: 訂閱要接收的賽事編號
//   schema:
//     type: object
//     $ref: "#/definitions/SampleSubMsg"
// responses:
//   200:
//     description: 回應成功
//     schema:
//       type: object
//       $ref: "#/definitions/BroadcastSampleMessage"
