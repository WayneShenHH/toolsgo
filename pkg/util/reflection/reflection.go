// Package reflection get struct info by reflect
package reflection

import "reflect"

// StructName get struct name
func StructName(data interface{}) (name string) {
	if reflect.TypeOf(data).Kind() == reflect.Ptr {
		name = reflect.TypeOf(data).Elem().Name()
	} else {
		name = reflect.ValueOf(data).Type().Name()
	}
	return
}
