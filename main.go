package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("listener error")
	}
	defer ln.Close()
	log.Printf("server started at %v", ln.Addr())

	go listenInput(func() { log.Fatal("server closed"); ln.Close(); return })

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("connection error")
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	log.Printf("%v", conn)
	conn.Write([]byte{'h', 'e', 'l', 'l', 'l', 'o'})
	conn.Close()
}

func listenInput(fn func()) {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		if text == "q\n" {
			fn()
		}
	}
}
