package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/shuyi-tangerine/little_book_book/thrift"
	"os"
)

func Usage() {
	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
}

func main() {
	flag.Usage = Usage
	server := flag.Bool("server", false, "Run server")
	protocol := flag.String("P", "binary", "Specify the protocol (binary, compact, json, simplejson)")
	framed := flag.Bool("framed", false, "Use framed transport")
	buffered := flag.Bool("buffered", false, "Use buffered transport")
	addr := flag.String("addr", "localhost:9090", "Address to listen to")
	secure := flag.Bool("secure", false, "Use tls secure transport")

	flag.Parse()
	thriftServer := thrift.NewServer(*server, *protocol, *buffered, *framed, *addr, *secure)
	err := thriftServer.Start(context.Background())
	if err != nil {
		fmt.Println("start error", err)
	}
}
