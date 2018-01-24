package main

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"github.com/mingkaic/Practice/Practice1/putil"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 3 {
		log.Fatalf("insufficient arguments: require 3")
	}
	var wg sync.WaitGroup
	maxbase := putil.StrConv([]byte(args[0]))
	b10 := putil.StrConv([]byte(args[1]))
	nStart := int(putil.StrConv([]byte(args[2])))
	if nStart < 2 {
		nStart = 2
	}
	wg.Add(int(maxbase) - nStart)
	var outstr string
	for i := nStart; i < int(maxbase); i++ {
		go func(wg *sync.WaitGroup, base, b10 int64) {
			result := putil.BaseConv(base, b10)
			outstr += putil.EncodeToks(result) + ","
			wg.Done()
		}(&wg, int64(i), b10)
	}
	wg.Wait()
	fmt.Println(outstr)
	fmt.Println(putil.EncodeToks(putil.BaseConv(maxbase, b10)))
}
