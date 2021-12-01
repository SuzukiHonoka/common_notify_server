package main

import (
	"fmt"
	"nfly/cmd/server"
)

func main() {
	fmt.Printf("NFLY, a cross-plaform notify framework for devs and pros, which can simply secure your notification" +
		" also within datas.\nWritten in golang by starx.\n")
	server.Run()
}
