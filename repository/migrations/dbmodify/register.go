package dbmodify

import "github.com/go-gormigrate/gormigrate"

// 已註冊的 views
var initViews []*gormigrate.Migration

// Register 註冊要初始的 view
func Register(view *gormigrate.Migration) {
	if initViews == nil {
		initViews = []*gormigrate.Migration{}
	}
	initViews = append(initViews, view)
}

// Init 初始
func Init() []*gormigrate.Migration {
	Register(initTable)
	return initViews
}
