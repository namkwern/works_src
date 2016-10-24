package main
import(
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/event/touch"
	"golang.org/x/mobile/exp/sprite"
)

func main(){
	app.Main(func(a app.App){
		for e := range a.Events(){
			switch e.(type){
			case lifecycle.Event:
			case paint.Event:
			case size.Event:
			case touch.Event:
			}
		}
	})
}