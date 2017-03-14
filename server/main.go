package main

import (
	"flag"
	"net"
)

var (
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "testdata/server1.pem", "The TLS cert file")
	keyFile  = flag.String("key_file", "testdata/server1.key", "The TLS key file")
	bind     = flag.String("bind", ":5555", "The bind address")
)

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", *bind)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

}
