package jsonformat

import (
	"bytes"
	"fmt"
	"testing"
)

func ExampleJsonFormat() {
	r := bytes.NewReader([]byte(`{"name":"Weerasak"}`))
	formated, _ := From(r)
	fmt.Println(formated)
	// Output:
	// {
	//     "name": "Weerasak"
	// }
}

func TestFormatError(t *testing.T) {
	r := bytes.NewReader([]byte(`"name":"Weerasak"}`))
	_, err := From(r)
	if err == nil {
		t.Error("Expect json syntax error but nil")
	}
}
