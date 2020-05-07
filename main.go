package main

import (
	"fmt"
	Mygreet "./greeting"
	"github.com/tenntenn/greeting"
)

func main() {
	greet := Mygreet.Do()
	fmt.Println(greet, greeting.Do())
}
