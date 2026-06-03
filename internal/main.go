package main

import (
	"fmt"
	// proxy "rediss/proxy"
	server "rediss/server"
)

func main() {
	fmt.Printf("hello world\n")
	server.RunAsyncTCPServer()
	// proxy.InitProxyServer()
}
