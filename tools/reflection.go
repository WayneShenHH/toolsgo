package tools

import (
	"reflect"
	"strings"
)

// Information 紀錄結構體型態與屬性
type Information struct {
	class  reflect.Type
	fields map[string]reflect.Type
}

// Field 取得實例的指定屬性，該屬性須為Export，否則會panic
func Field(instance interface{}, field string) (value interface{}) {
	character := field[:1]
	if character == strings.ToLower(character) {
		return nil
	}
	switch temp := dereferncePtr(dereferncePtr(reflect.ValueOf(instance)).FieldByName(field)); temp.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value = temp.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		value = temp.Uint()
	case reflect.Float32, reflect.Float64:
		value = temp.Float()
	case reflect.Bool:
		value = temp.Bool()
	case reflect.String:
		value = temp.String()
	}
	return
}

func dereferncePtr(ptr reflect.Value) reflect.Value {
	if ptr.Kind() == reflect.Ptr {
		temp := reflect.Indirect(ptr)
		if temp.Kind() == reflect.Ptr {
			return dereferncePtr(temp)
		}
		return temp
	}
	return ptr
}

// Compare 比對兩個Struct，如果有不同的屬性就紀錄起來
func Compare(oldData, newData interface{}) (old, new map[string]interface{}) {
	old = make(map[string]interface{})
	new = make(map[string]interface{})

	oldDataV := reflect.ValueOf(oldData)
	newDataV := reflect.ValueOf(newData)
	if oldDataV.Type() == newDataV.Type() {
		length := oldDataV.NumField()

		for i := 0; i < length; i++ {
			character := oldDataV.Type().Field(i).Name[:1]
			if character == strings.ToLower(character) {
				continue
			}
			switch oldDataV.Field(i).Kind() {
			case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Struct:
				continue
			}
			oValue := oldDataV.Field(i).Interface()
			nValue := newDataV.Field(i).Interface()
			if oValue != nValue {
				fieldName := oldDataV.Type().Field(i).Name
				old[fieldName] = oValue
				new[fieldName] = nValue
			}
		}
	}
	return
}

// Info 建立結構體資訊，方便後續複製動作
func Info(t interface{}) *Information {
	clazz := dereferncePtr(reflect.ValueOf(t)).Type()
	info := Information{
		class:  clazz,
		fields: make(map[string]reflect.Type),
	}
	for i := 0; i < clazz.NumField(); i++ {
		field := clazz.Field(i)
		info.fields[field.Name] = field.Type
	}
	return &info
}

// Copy 將Orignal屬性複製到Destination相同名稱與型態的屬性上，dest必須為Pointer
func (info *Information) Copy(dest, orig interface{}) {
	real := dereferncePtr(reflect.ValueOf(dest))
	if real.Type() != (*info).class || !real.CanSet() {
		return
	}

	copy := dereferncePtr(reflect.ValueOf(orig))
	for n, t := range (*info).fields {
		if field := copy.FieldByName(n); field.IsValid() && field.Type() == t {
			real.FieldByName(n).Set(field)
		}
	}
}

// Copy 將Orignal屬性複製到Destination相同名稱與型態的屬性上，dest必須為Pointer
func Copy(dest, orig interface{}) {
	Info(dest).Copy(dest, orig)
}

// RemovePreload clear preloads
func RemovePreload(data interface{}) {
	value := reflect.ValueOf(data)
	if value.Kind() != reflect.Ptr {
		panic("must give a pointer")
	}
	dataV := dereferncePtr(value)
	length := dataV.NumField()
	for i := 0; i < length; i++ {
		character := dataV.Type().Field(i).Name[:1]
		if character == strings.ToLower(character) {
			continue
		}
		field := dataV.Field(i)
		switch field.Kind() {
		case reflect.Ptr:
			oValue := dereferncePtr(field)
			if oValue.Kind() == reflect.Struct {
				field.Set(reflect.Zero(field.Type()))
			}
		}
	}
	data = dataV.Addr()
}
