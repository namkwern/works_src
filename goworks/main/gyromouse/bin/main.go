package main

import (
	"fmt"
	"os"
	"os/exec"
	"io"
	"bufio"
	"net/textproto"
)

var PASS = []byte("25358")
var cmds [4]*exec.Cmd
var runFlag = true

func main(){
	//同じような処理
	go run("move", "54")
	run("event", "55")
}


func run(name string, uuid string){
	btsin, btsout := execute("bin\\btserver")
	btsin.Write([]byte(uuid + "\n"))
	mouin, _ := execute("bin\\mouse_" + name)
	btwraped := textproto.NewReader(bufio.NewReader(btsout))
	startIO(btwraped, mouin)
	
}


//マウスのイベント,IO実行
func startIO(stdout *textproto.Reader, stdin io.WriteCloser){
	for runFlag{
		line, err := stdout.ReadLine();check(err, 3)
		data := []byte(line)
		if isNum(data){
			exeIn(stdin, data)
		}else{
			fmt.Print(line + "\n")
		}
	}
}

//実行してioのパイプ取得
var cmdnum = 0
func execute(command string) (io.WriteCloser, io.ReadCloser){
	cmd := exec.Command(command)
	stdin, err := cmd.StdinPipe();check(err, 0)
	stdout, err := cmd.StdoutPipe();check(err, 1)
	cmd.Start()
	exeIn(stdin, PASS)
	cmds[cmdnum] = cmd
	cmdnum++
	return stdin, stdout
}

//stdinに書き込む
func exeIn(stdin io.WriteCloser, str ...[]byte){
	for n := 0; n < len(str); n++{
		stdin.Write(str[n])
		stdin.Write([]byte("\n"))
	}
}
/*
func readLine() string{
	var rv string
	fmt.Scan(&rv)
	return rv
}*/
func check(err error, num int){
	if err != nil{
		fmt.Printf("|ErrorNumber=%d|\n",num)
		finally()
	}
}

func isNum(b []byte) bool{
	flag := true
	for n := 0; n < len(b); n++{
		if (b[n] < '0' || '9' < b[n]) && (b[n] == '-' && n != 0){
			flag = false
		}
	}
	return flag
}
func finally(){
	for n := 0; n < len(cmds); n++{
		cmds[n].Process.Kill()
	}
	runFlag = false
	os.Exit(1)
}