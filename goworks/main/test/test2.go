package main

import (
	"fmt"
	"os"
	"../gomyshell/my"
)

func main() {
	nstdin := my.SjisReader(os.Stdin)
	s ,_ := nstdin.ReadLine()
	str, _ := my.AutoEnc(s)
	fmt.Println(s, []byte(s), str, []byte(str), []byte("え"))
}