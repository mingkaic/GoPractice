package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/mingkaic/Practice/Practice1/putil"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		log.Fatalf("insufficient arguments: require 2")
	}
	var base int64 = putil.StrConv([]byte(args[0]))
	var b10 int64 = putil.StrConv([]byte(args[1]))
	result := putil.BaseConv(base, b10)
	fmt.Println(putil.EncodeToks(result))
}
