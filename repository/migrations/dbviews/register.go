package dbviews

// 已註冊的 views
var initViews []string

// Register 註冊要初始的 view
func Register(view string) {
	if initViews == nil {
		initViews = []string{}
	}
	initViews = append(initViews, view)
}

// Init 初始
func Init() []string {
	Register(SampleView)
	return initViews
}
