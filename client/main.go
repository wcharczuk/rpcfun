package main

import (
	"flag"
	"fmt"

	logger "github.com/blendlabs/go-logger"
)

func main() {
	server := flag.String("server", "localhost:5555", "The server to connect to")
	threads := flag.Int("threads", 4, "Number of threads to use")

	for i := 0; i < *threads; i++ {
		go pumpMessages(*server)
	}
}
func pumpMessages(server string) {
	defer func() {
		if r := recover(); r != nil {
			logger.Default().Fatal(fmt.Errorf("%v", r))
		}
	}()
}
