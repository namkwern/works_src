package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

type handler struct{}

func (h handler) ServeHTTP(resw http.ResponseWriter, requ *http.Request) {
	fmt.Println()
	log.Print(requ.RemoteAddr)
	fmt.Println(requ.URL.Path)
	fin, err := os.Open("public" + requ.URL.Path)
	if err != nil {
		return
	}
	str, _ := ioutil.ReadAll(fin)
	if len(str) != 0 {
		resw.Write(str)
	}
}

func main() {
	host, err := os.Hostname()
	catch(err)
	addrs, err := net.LookupHost(host)
	catch(err)
	log.Print("Service Start:")
	for _, v := range addrs {
		fmt.Println(v)
	}

	port := ":80"
	fmt.Println(port)

	http.ListenAndServe(port, handler{})
}

func catch(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
