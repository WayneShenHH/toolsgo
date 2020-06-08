// Package version application version info for ldflags
package version

import (
	"fmt"
	"strings"
)

const name = "txgo"

// // MajorVersion 主版號
// const MajorVersion = 0

// // MinorVersion 次版號
// const MinorVersion = 1

// // PatchVersion 修訂號
// const PatchVersion = 0

// Version 版本
var version string

// Revision git 版控 hash 值
var revision string

// Branch 版本名稱
var branch string

// Format 輸出版本編號
func Format() string {
	versionFormat := "0.0.0-edge"

	// 如果有值傳入
	if version != "" || branch != "" || revision != "" {
		versionsArr := []string{}
		if version != "" {
			versionsArr = append(versionsArr, version)
		}
		if branch != "" {
			versionsArr = append(versionsArr, branch)
		}
		if revision != "" {
			versionsArr = append(versionsArr, revision)
		}
		versionFormat = strings.Join(versionsArr, "-")
	}
	return versionFormat
}

// Full 完整編號包含名稱
func Full() string {
	return fmt.Sprintf(`%v@%v`, Name(), Format())
}

// Name 專案名稱
func Name() string {
	return name
}
