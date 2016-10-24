package main 
import(
	"fmt"
	"time"
	"github.com/argandas/goduino"
)

var adn *goduino.Goduino
const LED = 0x0d
func main(){
	adn = goduino.New("Arduino1", "COM4")
	err := adn.Connect()
	defer adn.Disconnect()
	if err != nil {
		fmt.Println(err)
	}
	adn.PinMode(LED, goduino.Output)
	
	K002()
	
}

func K001(){
	adn.DigitalWrite(LED, 1)
	for{}
}

func K002(){
	count := 0
	for {
		
		adn.DigitalWrite(LED, count)
		time.Sleep(1 * time.Second)
		count++
		count %= 3
	}
}