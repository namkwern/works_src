package main
import(
	"fmt"
	//"net"
	"net/http"
	"io/ioutil"
)

func main(){
	res, _ := http.Get("http://google.co.jp/")
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("\n\n%s\n\n", string(body))
}
