package main
import(
    "fmt"
    "net"
)

func main(){
	lis, _ := net.Listen("tcp", ":8080")
	for{
		con, _ := lis.Accept()
		go conHandle(con)
	}
}

func conHandle(con net.Conn){
	fmt.Println("get connection")
	for{
		s := ""
		fmt.Scan(&s)
		con.Write([]byte(s))
	}
}