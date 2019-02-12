package triggers

// 已註冊的 views
var initTriggers []string

// Register 註冊要初始的 view
func Register(trigger string) {
	if initTriggers == nil {
		initTriggers = []string{}
	}
	initTriggers = append(initTriggers, trigger)
}

// Init 初始
func Init() []string {
	Register(SampleTrigger)
	return initTriggers
}
