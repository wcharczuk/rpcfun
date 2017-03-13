package main

//go:generate protoc --go_out=import_path=proto:. proto/rpcfun.proto

import "flag"

func main() {
	server := flag.String("server", "localhost:5555", "The server to connect to")
	threads := flag.Int("threads", 4, "Number of threads to use")


}
