package main

import (
	"fmt"
	"./greeting"
)

func main() {
	greet := greeting.Do()
	fmt.Println(greet)
}
