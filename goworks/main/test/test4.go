package main
import (
	"fmt"
	"net"
)

func main() {
	con, _ := net.Dial("tcp", "localhost:8080")
	for{
		buf := make([]byte, 100)
		size, _ := con.Read(buf)
		if size != 0{
			fmt.Println(string(buf[:size]))
		}
	}
}