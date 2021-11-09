package main

import (
	"common_notify_server/cmd/server"
	"fmt"
)

func main() {
	fmt.Printf("NFLY, a cross-plaform notify framework for devs and pros, which can simply secure your notification" +
		" also within datas.\nWritten in golang by starx.\n")
	server.Run()
}
