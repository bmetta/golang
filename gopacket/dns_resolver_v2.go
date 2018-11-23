package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/routing"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

type scanner struct {
	iface *net.Interface
	// destination, gateway (if applicable), and source IP
	// addresses to use.
	dst, gw, src net.IP
	eth          layers.Ethernet
	ipv4         layers.IPv4
	handle       *pcap.Handle
	options      gopacket.SerializeOptions

	srcPort     layers.UDPPort
	srcPortLock sync.RWMutex
}

var (
	device       string = "eth0"
	snapshot_len int32  = 1024
	promiscuous  bool   = false
	err          error
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

func Listen(s *scanner, conn net.Conn, resolver string) {
	defer conn.Close()

	for {
		data, addr, err := Read(conn)
		if err != nil {
			log.Println("Read error %v", err)
			return
		}
		//log.Println("Received dns request from", addr)
		resp, err := Resolve(s, resolver, data)
		if err != nil {
			log.Println("Resolve error ", err)
			//return
		}
		_, err = Write(conn, addr, resp)
		if err != nil {
			log.Println("Write error %v", err)
			return
		}
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

func Resolve(s *scanner, resolver string, data []byte) ([]byte, error) {
	//return Get(resolver, data)
	return GoPacketGet(s, resolver, data)
}

func GoPacketGet(s *scanner, address string, data []byte) ([]byte, error) {
	s.srcPortLock.Lock()
	s.srcPort++
	if s.srcPort > 65530 {
		s.srcPort = 34000
	}
	udp := layers.UDP{
		SrcPort: s.srcPort,
		DstPort: 53,
	}
	s.srcPortLock.Unlock()

	udp.SetNetworkLayerForChecksum(&s.ipv4)

	transactionIdBytes := make([]byte, 2)
	_, err := rand.Read(transactionIdBytes)
	if err != nil {
		return nil, err
	}
	newData := append([]byte{}, data...)
	newData[0], newData[1] = transactionIdBytes[0], transactionIdBytes[1]

	buffer := gopacket.NewSerializeBuffer()
	if err := gopacket.SerializeLayers(buffer, s.options,
		&s.eth,
		&s.ipv4,
		&udp,
		gopacket.Payload(newData),
	); err != nil {
		return nil, err
	}

	outgoingPacket := buffer.Bytes()
	err = s.handle.WritePacketData(outgoingPacket)
	if err != nil {
		return nil, err
	}
	log.Println("Sent the dns request")

	start := time.Now()
	for {
		if time.Since(start) > time.Second*3 {
			return nil, errors.New("timeout getting DNS reply")
		}
		r_data, _, err := s.handle.ReadPacketData()
		if err != nil {
			return nil, err
		}
		packet := gopacket.NewPacket(r_data, layers.LayerTypeEthernet, gopacket.NoCopy)
		udpLayer := packet.Layer(layers.LayerTypeUDP)
		if udpLayer == nil {
			continue
		}
		udpLayerData := udpLayer.(*layers.UDP)
		if udpLayerData.SrcPort != 53 || udpLayerData.DstPort != udp.SrcPort {
			continue
		}

		resp := udpLayerData.Payload
		if resp[0] != transactionIdBytes[0] || resp[1] != transactionIdBytes[1] {
			continue
		}
		resp[0], resp[1] = data[0], data[1]
		log.Println("Received the dns response")
		return resp, nil
	}
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

func SetLogConfig() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func CreateScanner(resolver string) (*scanner, error) {
	router, err := routing.New()
	if err != nil {
		return nil, err
	}
	ip := net.ParseIP(resolver)
	if ip == nil {
		return nil, errors.New(resolver + " parse error")
	}
	s := &scanner{
		dst: ip,
		options: gopacket.SerializeOptions{
			FixLengths:       true,
			ComputeChecksums: true,
		},
		srcPort: 34000,
	}
	// Figure out the route to the IP.
	iface, gw, src, err := router.Route(ip)
	if err != nil {
		return nil, err
	}
	s.gw, s.src, s.iface = gw, src, iface

	handle, err := pcap.OpenLive(iface.Name, 65536, true, pcap.BlockForever)
	if err != nil {
		return nil, err
	}
	s.handle = handle
	hwaddr, err := s.getHwAddr()
	if err != nil {
		return nil, err
	}
	s.eth = layers.Ethernet{
		SrcMAC:       s.iface.HardwareAddr,
		DstMAC:       hwaddr,
		EthernetType: layers.EthernetTypeIPv4,
	}
	s.ipv4 = layers.IPv4{
		SrcIP:    s.src,
		DstIP:    s.dst,
		Version:  4,
		TTL:      64,
		Protocol: layers.IPProtocolUDP,
	}
	return s, nil
}

func (s *scanner) getHwAddr() (net.HardwareAddr, error) {
	start := time.Now()
	arpDst := s.dst
	if s.gw != nil {
		arpDst = s.gw
	}
	// Prepare the layers to send for an ARP request.
	eth := layers.Ethernet{
		SrcMAC:       s.iface.HardwareAddr,
		DstMAC:       net.HardwareAddr{0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		EthernetType: layers.EthernetTypeARP,
	}
	arp := layers.ARP{
		AddrType:          layers.LinkTypeEthernet,
		Protocol:          layers.EthernetTypeIPv4,
		HwAddressSize:     6,
		ProtAddressSize:   4,
		Operation:         layers.ARPRequest,
		SourceHwAddress:   []byte(s.iface.HardwareAddr),
		SourceProtAddress: []byte(s.src),
		DstHwAddress:      []byte{0, 0, 0, 0, 0, 0},
		DstProtAddress:    []byte(arpDst),
	}
	var buf gopacket.SerializeBuffer
	buf = gopacket.NewSerializeBuffer()
	if err := gopacket.SerializeLayers(buf, s.options, &eth, &arp); err != nil {
		return nil, err
	}

	// Send a single ARP request packet
	if err := s.handle.WritePacketData(buf.Bytes()); err != nil {
		log.Printf("GES_TEST WritePacketData Failed")
		return nil, err
	}

	// Wait 3 seconds for an ARP reply.
	for {
		if time.Since(start) > time.Second*3 {
			return nil, errors.New("timeout getting ARP reply")
		}
		data, _, err := s.handle.ReadPacketData()
		if err == pcap.NextErrorTimeoutExpired {
			continue
		} else if err != nil {
			return nil, err
		}
		packet := gopacket.NewPacket(data, layers.LayerTypeEthernet, gopacket.NoCopy)
		if arpLayer := packet.Layer(layers.LayerTypeARP); arpLayer != nil {
			arp := arpLayer.(*layers.ARP)
			if net.IP(arp.SourceProtAddress).Equal(net.IP(arpDst)) {
				return net.HardwareAddr(arp.SourceHwAddress), nil
			}
		}
	}
}

func main() {
	if len(os.Args) != 3 {
		log.Println("usage: dns_resolver <lister_ip> <dns_server_ip>")
		return
	}
	listener := os.Args[1] + ":53"
	resolver := os.Args[2]
	log.Println("listener, dns_server", listener, resolver)

	SetLogConfig()

	conn, err := CreateUdpListener(listener)
	if err != nil {
		log.Println("CreateUdpListener error %v", err)
	}

	scanner, err := CreateScanner(resolver)
	if err != nil {
		log.Println("CreateScanner error %v", err)
	}

	Listen(scanner, conn, resolver)
}
