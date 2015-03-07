package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
		return
	}

	var inJson map[string]interface{} = make(map[string]interface{})
	err = json.Unmarshal(buf, &inJson)
	if err != nil {
		log.Fatal(err)
		return
	}
	outJson, err := json.MarshalIndent(inJson, "", "    ")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%s\n", outJson)
}
