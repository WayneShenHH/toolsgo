package tools

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

func LoadJson(file string) []byte {
	f, err := ioutil.ReadFile("./jsons/" + file + ".json")
	if err != nil {
		panic(err)
	}
	dst := new(bytes.Buffer)
	json.Compact(dst, f)
	return dst.Bytes()
}
