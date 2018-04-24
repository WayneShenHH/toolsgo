package tools

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

// LoadJSON load a json file
func LoadJSON(file string) []byte {
	f, err := ioutil.ReadFile("./assets/jsons/" + file + ".json")
	if err != nil {
		panic(err)
	}
	dst := new(bytes.Buffer)
	json.Compact(dst, f)
	return dst.Bytes()
}
