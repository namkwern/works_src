package main

import (
	"fmt"
	"flag"
	"os"
	"os/exec"
    "./my"
)

func main(){
	//コマンドライン引数関係
	fF := flag.Bool("f", false, "FullPath 絶対パスの表示")
	dF := flag.Bool("d", false, "Directory ディレクトリ検索(未指定ではファイル)")
	flag.Parse()
	
	//カレントディレクトリ関係
	cd, _ := os.Getwd()
	cdlen := len(cd)
	
	//コマンド呼び出し関係
	dstr := "/a-d"
	if *dF{
		dstr = "/ad"
	}
	option := []string{"/k", "dir", "/s", "/b", dstr};
	cmd := exec.Command("cmd", option...)
	stdout, _ := cmd.StdoutPipe()
	wOut := my.SjisReader(stdout)
	cmd.Start()
	
	//メイン処理実行
	for {
		str, err := wOut.ReadLine()
		if err != nil{
			break
		}
		
		start := 0
		if !*fF{
			start = cdlen + 1
		}
		
		if cdlen + 1 < len(str){
			str = str[start:]
			if my.MatchAll(str, flag.Args()){
				fmt.Println(str)
			}
		}
	}
}