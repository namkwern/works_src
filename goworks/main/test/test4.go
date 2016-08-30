package main
import (
	"fmt"
	"net"
	"bufio"
	"os"
	"../gomyshell/my"
)

func main(){
	lis, err := net.Listen("tcp", ":3845");catch(err)
	conn, err := lis.Accept();catch(err)
	status, err := bufio.NewReader(conn).ReadString('\n');catch(err)
	fmt.Println("name:", status)
	for{
		var s string
		fmt.Scanf("%s", &s)
		str, _ := my.FromSjis(s)
		fmt.Fprintln(conn, str)
	}
}
func catch(err error){
	if err != nil{
		fmt.Println("err =", err)
		os.Exit(1)
	}
}