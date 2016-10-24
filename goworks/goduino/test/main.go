package main 
import(
	"fmt"
//	"time"
	"github.com/argandas/goduino"
)

const LED = 0x0d
func main(){
	adn := goduino.New("Arduino1", "COM4")
	err := adn.Connect()
	defer adn.Disconnect()
	if err != nil {
		fmt.Println(err)
	}
	adn.PinMode(LED, goduino.Output)
	count := 0
	for {
		var s string
		fmt.Scanln(&s)
		switch s{
		case "write":
			adn.DigitalWrite(LED, count)
			count++
			count %= 2
		case "read" :
			onoff, _ := adn.DigitalRead(LED)
			adn.DigitalWrite(LED, onoff)
		case "exit" :
			adn.DigitalWrite(LED, 0)
			return
		}
	}
	
}
