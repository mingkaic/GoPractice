package main

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"github.com/mingkaic/GoPractice/Practice1/putil"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 3 {
		log.Fatalf("insufficient arguments: require 3")
	}
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}
	maxbase := putil.StrConv([]byte(args[0]))
	b10 := putil.StrConv([]byte(args[1]))
	nPerStr := int(putil.StrConv([]byte(args[2])))
	nOut := int(maxbase) - 1
	wg.Add(nOut)
	strBuffer := make([]string, nOut)
	strIdx := 0
	for i := 2; i <= int(maxbase); i++ {
		go func(base, b10 int64) {
			defer wg.Done()
			result := putil.BaseConv(base, b10)
			localstr := putil.EncodeToks(result)
			mutex.Lock()
			strBuffer[strIdx] = localstr
			strIdx++
			mutex.Unlock()
		}(int64(i), b10)
	}
	wg.Wait()
	for i, str := range strBuffer {
		if i%nPerStr == nPerStr-1 || i == nOut-1 {
			fmt.Println(str)
		} else {
			fmt.Print(str + ",")
		}
	}
}
