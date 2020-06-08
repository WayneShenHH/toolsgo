package mockqueue_test

import (
	"fmt"
	"testing"

	json "github.com/json-iterator/go"

	"github.com/WayneShenHH/toolsgo/pkg/module/mq/mockqueue"
)

func Test_mockNSQD_Produce(t *testing.T) {

	queue := mockqueue.New()

	tests := []struct {
		name    string
		topic   string
		wantMsg string
		ch      string
	}{
		// TODO: Add test cases.
		{
			name:    "t1",
			topic:   "test_topic",
			wantMsg: "hello",
			ch:      "ch1",
		},
		{
			name:    "t2",
			topic:   "test_topic2",
			wantMsg: "hello",
			ch:      "ch1",
		},
	}
	for _, tt := range tests {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			go queue.Produce(tc.topic, tc.wantMsg)
			done := queue.WaitTopicFinish(tc.topic)
			go queue.Consume(tc.topic, tc.ch, func(msg []byte) error {
				var obj string
				err := json.Unmarshal(msg, &obj)

				if err != nil {
					t.Errorf("json Unmarshal error: %s", err.Error())
					return fmt.Errorf("json Unmarshal error: %s", err)
				}

				if tc.wantMsg != obj {
					t.Errorf("mockNSQD.Consume() msg = %v, wantMsg %v", obj, tc.wantMsg)
					return fmt.Errorf("mockNSQD.Consume() msg = %v, wantMsg %v", obj, tc.wantMsg)
				}
				return nil
			})
			<-done
		})
	}
}
