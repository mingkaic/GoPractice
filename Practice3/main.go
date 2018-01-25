package main

import (
	"fmt"
)

func main() {
	cli := redisAPI.CreateLocal()
	cli.Connect()
	fmt.Print(cli.Ping())
}
