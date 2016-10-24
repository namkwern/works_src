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
	//log.Print(requ.RemoteAddr, requ.PostFormValue("text"))
	log.Printf("%#v", requ.RequestURI)
	fin, _ := os.Open("public" + requ.RequestURI)
	str, _ := ioutil.ReadAll(fin)
	fmt.Println(":", requ)
	if len(str) == 0 {
		resw.Write([]byte(htmltext))
		fmvalue := requ.FormValue("text")
		if len(fmvalue) != 0 {
			fmt.Println("\t\t", fmvalue)

			res, err := http.Get(fmvalue)
			if err != nil {
				return
			}
			str, _ := ioutil.ReadAll(res.Body)
			defer res.Body.Close()
			fmt.Println(string(str))
		}
	} else {
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
