package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type buff []byte
type jobj map[string]interface{}

func main() {
	defer func() {
		logError(recover())
	}()
	fmt.Printf("%s\n", readBuff().unmarshal().format())
}

func readBuff() buff {
	return buff(readInput())
}

func readInput() (buf []byte) {
	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	return
}

func (buf buff) unmarshal() jobj {
	return jobj(unmarshal([]byte(buf)))
}

func (j jobj) format() string {
	return format(j)
}

func unmarshal(buf []byte) (data map[string]interface{}) {
	data = make(map[string]interface{})
	if err := json.Unmarshal(buf, &data); err != nil {
		panic(err)
	}
	return
}

func format(data map[string]interface{}) (out string) {
	buf, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

func logError(err interface{}) {
	if err != nil {
		log.Fatal(err)
	}
}
