package main

import (
	"fmt"
	"log"
	"os"

	"github.com/iporsut/jsonformat"
)

func main() {
	formated, err := jsonformat.From(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(formated)
}
