package main
import (
	"fmt"
	"net"
	"os"
	"../gomyshell/my"
)

func main(){
	conn, err := net.Dial("tcp", "192.168.112.1:3845");catch(err)
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	out := my.SjisReader(conn)
	for{
		status, err := out.ReadLine();catch(err)
		fmt.Print(status)
	}
}
func catch(err error){
	if err != nil{
		fmt.Println("err =", err)
		os.Exit(1)
	}
}