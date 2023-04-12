package main

import (
	"context"
	"fmt"
	"github.com/shuyi-tangerine/little_book_book/gen-go/base"
	"github.com/shuyi-tangerine/little_book_book/gen-go/tangerine/little_book_book"

	"github.com/apache/thrift/lib/go/thrift"
)

var defaultCtx = context.Background()

func handleClient(client little_book_book.LittleBookBooker) (err error) {
	res, err := client.SaveContent(defaultCtx, &little_book_book.SaveContentRequest{
		Text: "test\ntest02\nbiubiubiu\n",
		Base: &base.RPCRequest{}})
	if err != nil {
		return err
	}
	fmt.Println("SaveContent res ==>", res.Base.Code, res.Base.Message)
	fmt.Println("SaveContent data ==>", res.Content)

	getContentRes, err := client.GetContent(defaultCtx, &little_book_book.GetContentRequest{
		Base: &base.RPCRequest{}})
	if err != nil {
		return err
	}
	fmt.Println("GetContent res ==>", getContentRes.Base.Code, getContentRes.Base.Message)
	fmt.Println("GetContent data ==>", getContentRes.Content)
	return
}

func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool, cfg *thrift.TConfiguration) error {
	var transport thrift.TTransport
	if secure {
		transport = thrift.NewTSSLSocketConf(addr, cfg)
	} else {
		transport = thrift.NewTSocketConf(addr, cfg)
	}
	transport, err := transportFactory.GetTransport(transport)
	if err != nil {
		return err
	}
	defer transport.Close()
	if err := transport.Open(); err != nil {
		return err
	}
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	return handleClient(little_book_book.NewLittleBookBookerClient(thrift.NewTStandardClient(iprot, oprot)))
}
