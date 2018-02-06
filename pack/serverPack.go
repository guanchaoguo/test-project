package main

import (
	"fmt"
	"net"
)

const BUF_SIZE = 20

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, BUF_SIZE)
	n, err := conn.Read(buf)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Printf("\n已接收：%d个字节，数据是：'%s'\n", n, string(buf))
}

func main() {
	ln, err := net.Listen("tcp", ":9872")

	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)
	}
}