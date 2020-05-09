package main

import (
		"fmt"
		"os"
		"flag"
		"strings"

    "rsc.io/quote/v2"
)

var msg = flag.String("msg", "defalt", "説明")
var n int

func init() {
	flag.IntVar(&n, "n", 1, "回数")
}

func main() {
		fmt.Println(quote.HelloV2())
		flag.Parse()
		fmt.Println(strings.Repeat(*msg, n))
		fmt.Println(os.Args)
}
