package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello, HTTP server")
}

func main() {
	type Person struct {
		Name string `json: "name"`
		Age  int    `json: "age"`
	}
	p := &Person{Name: "wataru", Age: 21}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(p); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
}
