// Package errors 錯誤處理
//nolint:unused // 預先定義錯誤類型接受沒有使用的定義，因此加入 nolint
package errors

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"
)

// Error 錯誤結構體
type Error struct {
	// Op 是執行的操作，通常為方法名稱
	Op Op
	// Kind 是錯誤的類型，例如權限、 I/O
	Kind Kind
	// Code 是錯誤代碼，如 1001 注單餘額不足
	Code Code
	// Err 是比現在的 Error 更底層的錯誤
	Err error
}

// Op 操作的描述，通常為套件與方法名
// ex: "errors/Error.String"
type Op string

// Op 預先定義的操作型態
const (
	AutoOp Op = "AUTO" // 自動抓取路徑，同 stack trace 一樣
)

// Kind 錯誤的種類描述，由 errors 定義可能的類型
// ex: errors.IO
type Kind uint8

// Code 四位數字的錯誤代碼，首位區分類型
// ex: 1000 （餘額不足）
type Code uint16

// Kind 預先定義的錯誤類型
const (
	Other      Kind = iota // 未知錯誤類型總括，不會被印出
	Database               // 資料庫操作發生的錯誤
	Invalid                // 各種因素導致操作無效
	Permission             // 使用者的權限不允許
	Data                   // 資料缺失或格式錯誤導致處理失敗
)

// Code 預先定義的錯誤代碼
const (
	Null Code = 0 // 空的錯誤代碼，用於標示、檢測
)

func (k Kind) String() string {
	switch k {
	case Other:
		return "未知錯誤類型"
	case Database:
		return "資料庫錯誤"
	case Permission:
		return "權限不足"
	case Invalid:
		return "無效的操作"
	case Data:
		return "資料的缺損"
	}
	return "無定義"
}

func (c Code) String() string {
	return desc[c]
}

// E 建立 Error
func E(args ...interface{}) error {
	if len(args) == 0 {
		panic("errors.E 未收到任何參數")
	}
	e := new(Error)
	for i := range args {
		switch arg := args[i].(type) {
		case Op:
			e.Op = arg
		case Kind:
			e.Kind = arg
		case Code:
			e.Code = arg
		case *Error:
			e.Err = arg
		case string:
			e.Err = Errorf(arg)
		case error:
			e.Err = arg
		default:
			_, file, line, _ := runtime.Caller(1)
			log.Printf("errors.E: bad call from %s:%d: %v", file, line, arg)
			return Errorf("無效型態 %T, value %v in error.E call", arg, arg)
		}
	}

	if e.Op == AutoOp {
		e.Op = caller()
	}

	prev, ok := e.Err.(*Error)
	if !ok {
		return e
	}

	if prev.Code == e.Code {
		prev.Code = 0
	}
	if prev.Kind == e.Kind {
		prev.Kind = Other
	}
	if e.Kind == Other {
		e.Kind = prev.Kind
		prev.Kind = Other
	}
	if e.Code == Null {
		e.Code = prev.Code
		prev.Code = Null
	}

	// logger.Error(e.Error())
	return e
}

// Error 格式化錯誤訊息
func (e *Error) Error() string {
	buffer := new(bytes.Buffer)
	if e.Op != "" {
		pad(buffer, ": ")
		buffer.WriteString(string(e.Op))
	}
	if e.Kind != 0 {
		pad(buffer, ": ")
		buffer.WriteString(e.Kind.String())
	}
	if e.Code != 0 {
		pad(buffer, ": ")
		buffer.WriteString(e.Desc())
		pad(buffer, " (")
		buffer.WriteString(strconv.FormatUint(uint64(e.Code), 10))
		pad(buffer, ") ")
	}
	if e.Err != nil {
		if prevErr, ok := e.Err.(*Error); ok {
			if !prevErr.IsZero() {
				pad(buffer, ":\n\t")
				buffer.WriteString(e.Err.Error())
			}
		} else {
			pad(buffer, ": ")
			buffer.WriteString(e.Err.Error())
		}
	}
	if buffer.Len() == 0 {
		return "無錯誤"
	}
	return buffer.String()
}

// Str 等效於 std/errors.Str
func Str(msg string) error {
	return errors.New(msg)
}

// Desc 取得 Code 相應的描述
func (e *Error) Desc() string {
	return e.Code.String()
}

// IsZero Error 是否為空值
func (e *Error) IsZero() bool {
	return e.Op == "" && e.Kind == 0 && e.Err == nil
}

// ExtractCode 從指定的 Error 取出錯誤代碼，若無則回傳 Null Code
func ExtractCode(err error) uint {
	if e, ok := err.(*Error); ok {
		return uint(e.Code)
	}
	return uint(Null)
}

// Is 是否為指定 Kind 的 Error
func Is(kind Kind, err error) bool {
	if e, ok := err.(*Error); ok && e.Kind == kind {
		return true
	}
	return false
}

// Match 是否與 Template 的錯誤相符
func Match(template, err error) bool {
	e1, ok := template.(*Error)
	if !ok {
		return false
	}
	e2, ok := err.(*Error)
	if !ok {
		return false
	}
	if e1.Op != "" && e2.Op != e1.Op {
		return false
	}
	if e1.Kind != 0 && e2.Kind != e1.Kind {
		return false
	}
	if e1.Code != 0 && e2.Code != e1.Code {
		return false
	}
	if e1.Err != nil {
		if _, ok := e1.Err.(*Error); ok {
			return Match(e1.Err, e2.Err)
		}
		if e2.Err == nil || e2.Err.Error() != e1.Err.Error() {
			return false
		}
	}
	return true
}

// Errorf 等效於 fmt.Errorf
func Errorf(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}

func pad(buffer *bytes.Buffer, s string) {
	if buffer.Len() == 0 {
		return
	}
	buffer.WriteString(s)
}

func caller() Op {
	pc := make([]uintptr, 1)
	if runtime.Callers(3, pc) != 0 {
		parts := strings.Split(runtime.FuncForPC(pc[0]).Name(), "/")
		return Op(strings.Replace(parts[len(parts)-1], ".", "/", 1))
	}
	return ""
}
