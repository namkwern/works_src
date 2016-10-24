package main
import(
	"fmt"
	"net"
	"os"
	"io/ioutil"
	"time"
)

func main(){
	addr, err := net.ResolveTCPAddr("tcp", ":80");catch(err,0)
	lis, err := net.ListenTCP("tcp", addr);catch(err,1)
	defer lis.Close()
	fmt.Println("Local Address:", lis.Addr())
	for{
		con, err := lis.AcceptTCP();catch(err,2)
		dur, _ := time.ParseDuration("2s")
		con.SetDeadline(time.Now().Add(dur))
		go res(con)
		go req(con)
	}
	
}

func res(con *net.TCPConn){
	fmt.Println("接続されました:", con.RemoteAddr())
	htmltext := "<html>" + head + body + "</html>"
	fmt.Fprintln(con, htmltext)
	defer con.CloseWrite()
}

func req(con *net.TCPConn){
	fmt.Println("接続されました:", con.RemoteAddr())	
	
	fmt.Println("リクエスト受信開始:", con.RemoteAddr())
	b, _ := ioutil.ReadAll(con);
	str := string(b)
	fmt.Println("▼リクエスト内容:", con.RemoteAddr())
	fmt.Println(str)
	fmt.Println("▲")
	defer con.CloseRead()
}

func catch(err error, num int){
	if err != nil{
		fmt.Println("err:", err)
		fmt.Println("errNum:", num)
		os.Exit(1)
	}
}

var(
head = `
<head>
<title>俳句</title>
<meta charset="utf-8">
</head>
`

body = `
<body>
<pre>
古池や
　　蛙飛び込む
　　　　　   <font color="#666666">水の音</font>
</pre>
<form>
	<input type="text">
	<input type="submit">
</form>
</body>
`
)
