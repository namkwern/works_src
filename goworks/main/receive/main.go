package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
)

var (
	reg = regexp.MustCompile("//.*/")
)

type handler struct{}

func (h handler) ServeHTTP(resw http.ResponseWriter, requ *http.Request) {
	log.Print(requ.RequestURI[1:])
	fmt.Println(requ.Referer())

	b, _ := ioutil.ReadAll(requ.Body)
	requ.Body.Close()

	url := "https://google.com/"

	req, err := http.NewRequest("GET", url, bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	req.URL = req.URL
	req.Header = requ.Header
	req.Header["Referer"] = []string{url}
	fmt.Printf("::%#v", requ)
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	b, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()
	resw.Write(b)
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
	fmt.Println()

	http.ListenAndServe(":80", handler{})
}

func catch(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
