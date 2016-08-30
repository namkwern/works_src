package main
import(
	"fmt"
	"os"
	"image/png"
)

func main(){
	fp, err := os.Open("images/new.png")
	if err != nil{
		fmt.Println(err)
		return
	}
	img, err := png.Decode(fp)
	if err != nil{
		fmt.Println(err)
		return
	}
	fp.Close()
}