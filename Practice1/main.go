package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/mingkaic/GoPractice/Practice1/putil"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		log.Fatalf("insufficient arguments: require 2")
	}
	base := putil.StrConv([]byte(args[0]))
	b10 := putil.StrConv([]byte(args[1]))
	result := putil.BaseConv(base, b10)
	fmt.Println(putil.EncodeToks(result))
}
