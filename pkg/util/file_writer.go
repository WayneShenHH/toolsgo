package tools

import (
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func WriteFile(bytes []byte) {
	err := ioutil.WriteFile("/tmp/dat1", bytes, 0644)
	check(err)
}
