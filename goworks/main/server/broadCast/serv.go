package main

import (
	"fmt"
	"net"
	"log"
	"strings"
)
var(
	addr, _ = net.ResolveUDPAddr("udp", ":37298")
	buf = make([]byte, 100)
)
func main() {
	for {
		str, raddr := getAddrUDP()
		fmt.Println(str, raddr)
		rstr := raddr.String()
		restr := rstr[:strings.LastIndex(rstr, ":")] + str
		remoteAddr, _ := net.ResolveUDPAddr("udp", restr)
		//send
		sendUDP("TCPPort:", remoteAddr)
	}
}

//受信文字列, アドレス
func getAddrUDP()(string, *net.UDPAddr){
	con, err := net.ListenUDP("udp", addr)
	if err != nil{
		log.Fatal(err)
	}
	defer con.Close()
	i, raddr, _ := con.ReadFromUDP(buf)
	return string(buf[:i]), raddr
}

func sendUDP(str string, raddr *net.UDPAddr){
	con, err := net.DialUDP("udp", addr, raddr)
	if err != nil{
		log.Fatal(err)
	}
	defer con.Close()
	con.Write([]byte(str))
}
