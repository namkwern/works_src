package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	host, _ := os.Hostname()
	addrs, _ := net.LookupHost(host)
	log.Print("Service Start:")
	for _, v := range addrs {
		fmt.Println(v)
	}
	const port = ":37298"
	fmt.Println(port)
	lis, err := net.Listen("tcp", port)
	fatal(err)
	for {
		con, err := lis.Accept()
		errPrint(err)
		log.Println("Connect:", con.RemoteAddr())
		go handle(con)
	}
}

func handle(con net.Conn) {

	buf := make([]byte, 100)
	for {
		i, _ := con.Read(buf)
		if i == 0{
			continue
		}
		fmt.Println(con.RemoteAddr(), buf[:i])
		con.Write([]byte("通知"))
	}
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func errPrint(err error) {
	if err != nil {
		log.Print(err)
	}
}
