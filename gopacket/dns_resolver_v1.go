package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"net"
)

var (
	listener = "192.168.7.14:53"
	resolver = "8.8.8.8:53"
)

func CreateUdpListener(address string) (net.Conn, error) {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return nil, err
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func Listen(conn net.Conn) {
	defer conn.Close()

	for {
		data, addr, err := Read(conn)
		if err != nil {
			log.Println("Read error %v", err)
			return
		}
		log.Println("Received dns request from", addr)
		resp, err := Resolve(resolver, data)
		if err != nil {
			log.Println("Resolve error %v", err)
			return
		}
		_, err = Write(conn, addr, resp)
		if err != nil {
			log.Println("Write error %v", err)
			return
		}
		log.Println("sent the dns response to", addr)
	}
}

func Read(conn net.Conn) ([]byte, net.Addr, error) {
	buffer := make([]byte, 4096)
	n, addr, err := conn.(*net.UDPConn).ReadFrom(buffer)
	return buffer[:n], addr, err
}

func Write(conn net.Conn, addr net.Addr, data []byte) (int, error) {
	if addr == nil {
		return conn.Write(data)
	}
	return conn.(*net.UDPConn).WriteTo(data, addr.(*net.UDPAddr))
}

func Resolve(resolver string, data []byte) ([]byte, error) {
	return Get(resolver, data)
}

func Get(address string, data []byte) ([]byte, error) {
	conn, err := net.Dial("udp", address)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	transactionIdBytes := make([]byte, 2)
	_, err = rand.Read(transactionIdBytes)
	if err != nil {
		return nil, err
	}
	newData := append([]byte{}, data...)
	newData[0], newData[1] = transactionIdBytes[0], transactionIdBytes[1]
	if _, err := Write(conn, nil, newData); err != nil {
		return nil, err
	}
	log.Println("sent dns querry to", address)

	resp, addr, err := Read(conn)
	if err != nil {
		return nil, err
	}
	log.Println("received dns response from", addr)

	if resp[0] != transactionIdBytes[0] || resp[1] != transactionIdBytes[1] {
		return nil, fmt.Errorf("Incorrect transaction ID in UDP DNS response.")
	}
	resp[0], resp[1] = data[0], data[1]
	return resp, nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conn, err := CreateUdpListener(listener)
	if err != nil {
		log.Println("listen error %v", err)
	}
	log.Println("listener running on", listener)

	Listen(conn)
}
