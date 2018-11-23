package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: 1234})
	if err != nil {
		log.Fatal("Listen:", err)
	}
	sendHello(conn, &net.UDPAddr{IP: net.IP{8, 8, 8, 8}, Port: 53})
	sendHello(conn, &net.UDPAddr{IP: net.IP{8, 8, 4, 4}, Port: 53})
}

func sendHello(conn *net.UDPConn, addr *net.UDPAddr) {
	n, err := conn.WriteTo([]byte("hello"), addr)
	if err != nil {
		log.Fatal("Write:", err)
	}
	fmt.Println("Sent", n, "bytes", conn.LocalAddr(), "->", addr)
}
