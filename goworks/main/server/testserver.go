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
	fmt.Println(resw)
	fmt.Println(requ.Header["Content-Type"])
	fmt.Println(requ.TransferEncoding)
	fmt.Println(requ.Host)

	by, _ := ioutil.ReadAll(requ.Body)
	fmt.Println(string(by))

	resw.Write([]byte(htmltext))
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

var htmltext = `
<html lang="ja">
<head>
<meta charset="utf-8" content="ja">
<title>テスト</title>
<head>
<body>
<form method="post">
<input type="text" name="text">
<input type="submit">
<img src="new.png">
</form>
</body>
</html>
`
