package jsonformat

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
)

func From(r io.Reader) (formated string, err error) {
	defer func() {
		if recovered := recover(); recovered != nil {
			err = recovered.(error)
		}
	}()
	formated = format(read(r))
	return
}

func read(r io.Reader) (b []byte) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return
}

func format(b []byte) string {
	dst := new(bytes.Buffer)
	if err := json.Indent(dst, b, "", "    "); err != nil {
		panic(err)
	}
	return string(dst.Bytes())
}
