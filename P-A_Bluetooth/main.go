package main

import (
	"fmt"
	"os"
	"os/exec"
	"io"
	"strconv"
)

var PASS = []byte("25358")

var cmds [4]*exec.Cmd

func main(){
	//同じような処理
	go run("move", "54")
	run("event", "55")
}

var cmdnum = 0
func run(event string, uuid string){
	bts, btsin, btsout := execute("bin\\btserver")
	btsin.Write([]byte(uuid + "\n"))
	mou, mouin, _ := execute("bin\\mouse_" + event)
	
	cmds[cmdnum] = bts
	cmdnum++
	cmds[cmdnum] = mou
	cmdnum++
	
	mouse(btsout, mouin)
	
}


//マウスのクリック(イベント)スレッド
func mouse(stdout io.ReadCloser, stdin io.WriteCloser){
	buf := make([]byte, 128)
	for{
		n, err := stdout.Read(buf);check(err, 3)
		if n > 0{
			fmt.Print(string(buf[0:n]))
			_, err := strconv.Atoi(string(buf[0:n]))
			if err == nil{
				exeIn(stdin, buf[0:n])
			}
		}
	}
}

//実行してioのパイプ取得
func execute(command string) (*exec.Cmd, io.WriteCloser, io.ReadCloser){
	cmd := exec.Command(command)
	stdin, err := cmd.StdinPipe();check(err, 0)
	stdout, err := cmd.StdoutPipe();check(err, 1)
	cmd.Start()
	exeIn(stdin, PASS)
	return cmd, stdin, stdout
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
		fmt.Printf("---Num=%d|Error=%s---\n",num, err.Error)
		for n := 0; n < len(cmds); n++{
			cmds[n].Process.Kill()
		}
		os.Exit(1)
	}
}
