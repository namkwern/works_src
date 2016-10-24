package main

import (
	"net"
	"fmt"
	"time"
	"strings"
)

func main() {
	fmt.Println(ScanServerUDP("37298"))
}


var (
	local, _ = net.ResolveUDPAddr("udp", ":0")
)

func ScanServerUDP(serverPort string) string {
	ch := make(chan string)
	go scanServerUDP(serverPort, ch)
	return <-ch
}
func scanServerUDP(serverPort string, arch chan string) {
	bcast, _ := net.ResolveUDPAddr("udp", "255.255.255.255:"+serverPort)
	ch := make(chan string)
	go listen(ch)
	addr := <-ch
	listenPort := addr[strings.LastIndex(addr, ":"):]
	sendUDP(listenPort, bcast)
	str := <-ch
	arch <- str
}
func sendUDP(str string, raddr *net.UDPAddr) {
again:
	con, err := net.DialUDP("udp", local, raddr)
	if err != nil {
		con.Close()
		goto again
	}
	defer con.Close()
	con.Write([]byte(str))
}
func listen(ch chan string) {
again:
	con, err := net.ListenUDP("udp", local)
	if err != nil {
		con.Close()
		goto again
	}
	ch <- con.LocalAddr().String()
	defer con.Close()
	now := time.Now()
	deadline := now.Add(1 << 34)
	con.SetReadDeadline(deadline)
	var buf = make([]byte, 100)
	_, raddr, err := con.ReadFromUDP(buf)
	if err != nil {
		con.Close()
		goto again
	}
	ch <- fmt.Sprintln(raddr)
}
