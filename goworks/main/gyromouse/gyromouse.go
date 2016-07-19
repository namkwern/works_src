package main

import(
	"fmt"
	"os/exec"
	"io"
	"bufio"
	"net/textproto"
)

var maincmd *exec.Cmd
func main(){
	_, out := execute("bin\\bin")
	wrapedOut := textproto.NewReader(bufio.NewReader(out))
	startIO(wrapedOut)
}

//IO実行
func startIO(stdout *textproto.Reader){
	for {
		line, err := stdout.ReadLine();check(err, 3)
		fmt.Println(line);
	}
}

//実行してioのパイプ取得
func execute(command string) (io.WriteCloser, io.ReadCloser){
	cmd := exec.Command(command)
	stdin, err := cmd.StdinPipe();check(err, 0)
	stdout, err := cmd.StdoutPipe();check(err, 1)
	cmd.Start()
	maincmd = cmd
	return stdin, stdout
}
func check(err error, num int){
	if err != nil{
		fmt.Printf("|ErrorNumber=%d|\n",num)
		finally(num == 3)
	}
}
func finally(flag bool){
	maincmd.Process.Kill()
	if flag{
		main()
	}
}