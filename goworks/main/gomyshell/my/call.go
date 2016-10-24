package my

import(
	"os/exec"
	"os"
	"fmt"
)


func Execute(command ...string) string{
	cmd := exec.Command(command[0], command[1:]...)
	str, _ := cmd.CombinedOutput()
	return string(str)
}
func check(err error, num int){
	if err != nil{
		fmt.Printf("|ErrorNumber=%d|\n",num)
		os.Exit(1)
	}
}