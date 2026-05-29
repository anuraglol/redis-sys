package main

import (
	"fmt"
	server "rediss/server"
)

func main() {
	fmt.Printf("hello world\n")
	server.RunSyncTCPServer()
}
