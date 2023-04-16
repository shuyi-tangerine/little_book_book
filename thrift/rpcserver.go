package thrift

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/shuyi-tangerine/little_book_book/gen-go/tangerine/little_book_book"
	"github.com/shuyi-tangerine/little_book_book/top"
	"os"

	"github.com/apache/thrift/lib/go/thrift"
)

type Server struct {
	isServer   bool
	protocol   string
	isBuffered bool
	isFramed   bool
	addr       string
	useSecure  bool
	c          chan error
}

func NewServer(isServer bool, protocol string, isBuffered bool, isFramed bool, addr string, useSecure bool) top.Server {
	return &Server{
		isServer:   isServer,
		protocol:   protocol,
		isBuffered: isBuffered,
		isFramed:   isFramed,
		addr:       addr,
		useSecure:  useSecure,
		c:          make(chan error),
	}
}

func (m *Server) IsBlock(ctx context.Context) (isBlock bool) {
	return m.isServer
}

func (m *Server) Start(ctx context.Context) (err error) {
	var protocolFactory thrift.TProtocolFactory
	switch m.protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactoryConf(nil)
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactoryConf(nil)
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryConf(nil)
	default:
		fmt.Fprint(os.Stderr, "Invalid protocol specified", m.protocol, "\n")
		flag.Usage()
		os.Exit(1)
	}

	var transportFactory thrift.TTransportFactory
	cfg := &thrift.TConfiguration{
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	if m.isBuffered {
		transportFactory = thrift.NewTBufferedTransportFactory(8192)
	} else {
		transportFactory = thrift.NewTTransportFactory()
	}

	if m.isFramed {
		transportFactory = thrift.NewTFramedTransportFactoryConf(transportFactory, cfg)
	}

	if m.isServer {
		handler := NewLittleBookBooker()
		processor := little_book_book.NewLittleBookBookerProcessor(handler)
		return runServer(transportFactory, protocolFactory, m.addr, m.useSecure, processor)
	} else {
		return runClient(transportFactory, protocolFactory, m.addr, m.useSecure, cfg)
	}
	return
}

func (m *Server) AsyncStart(ctx context.Context) {
	go func() {
		err := m.Start(ctx)
		if err != nil {
			fmt.Println("[AsyncStart] Start panic", err)
			m.c <- err
		}
	}()
}

func (m *Server) ErrorC() (c chan error) {
	return m.c
}
