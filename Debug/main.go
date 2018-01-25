package main

import (
	"fmt"
)

func main() {
	done := make(chan bool)

	for i := 0; i < 2; i++ {
		go func(i int) {
			squared := i * i
			fmt.Println(squared)
			done <- true
		}(i)
	}

	<-done
	<-done
}
