package main

import(
	"fmt"
	"io/ioutil"
	"flag"
	"./my"
	"strings"
	"strconv"
)

var(
	dFlag bool
	namesub []string
	textsub []string
	namesubnot []string
	textsubnot []string
	inF bool
	innF bool
	nFlag bool
	rFlag bool
)
func main(){
	rF := flag.Bool("r", false, "!!!表示・処理過多注意!!!\n\tRecursive 下層のファイル/ディレクトリをすべて検索")
	dF := flag.Bool("d", false, "Directory ディレクトリ検索に切り替え")
	nF := flag.Bool("n", false, "!!!表示過多注意!!!non-stop ファイル内検索時にまとめて表示")
	forS := flag.String("for", "", "ファイル名検索。正規表現で検索します。\n\tスペースは必ず\\sを指定してください。スペース区切りはAND")
	fornS := flag.String("forn", "", "forの否定検索版")
	inS := flag.String("in", "", "!!!表示・処理過多注意!!!\n\tファイル内検索。正規表現で検索します。\n\tスペースは必ず\\sを指定してください。スペース区切りはAND\n\t-forを指定することで読み込むファイルを減らそう！")
	innS := flag.String("inn", "", "-inの否定検索版")
	fromS := flag.String("from", ".", "検索を開始するディレクトリ")
	
	flag.Parse()
	
	//フラグの受け取り
	rFlag = *rF
	dFlag = *dF
	nFlag = *nF
	namesub = strings.Split(*forS, " ")
	textsub = strings.Split(*inS, " ")
	namesubnot = strings.Split(*fornS, " ")
	textsubnot = strings.Split(*innS, " ")
	
	inF = (*inS != "")
	innF = (*innS != "")
	
	
	//フラグエラー
	if dFlag && (inF || innF || nFlag){
		fmt.Println("<-d(ディレクトリ指定)を使用した場合は、-in/-inn(ファイル内検索)及び-n(ファイル内詳細)はすべて無効です>")
		inF = false
		innF = false
		nFlag = false
	}
	if !inF && !innF && nFlag{
		fmt.Println("<-inもしくは-innを指定していないので-nは無効です>\n")
	}
	//ファイル内検索指定ありで拡張子許可
	if inF || innF{
		namesub = append(namesub, "(\\.txt$|\\.java$|\\.go$|\\.c$|\\.cpp$|\\.bat$|\\.html$|\\.css$|\\.js$)")
	}
	
	//ファイル探索開始
	ch1 := make(chan bool)
	ch2 := make(chan string)
	ch3 := make(chan int)
	go recu(*fromS, "", ch1, ch2, ch3, true)
	if inF || innF{
		if !nFlag{
			fmt.Print("中断>Ctrl+C\nまとめて表示>all\n次の内容を表示>Enter:")
		}
		for {
			count := <- ch3
			first := <- ch1
			
			if {
				
			}
		}
		for ;count != 0; count--{
			v := <- ch2
			if !nFlag && v != ""{
				var s string
				fmt.Scanln(&s)
				if s == "all"{
					nFlag = true
				}
			}
			fmt.Print(v)
		}
	}
	fmt.Println("done.")
}

//カレントディレクトリ、ファイルパス
func recu(cur string, path string, ch1 chan bool, ch2 chan string, ch3 chan int, first bool){
	count := 0
	fds, _ := ioutil.ReadDir(cur + "\\" + path)
	for _, v := range fds{
		if v.IsDir(){
			if dFlag{
				if my.MatchAll(v.Name(), namesub) && my.NotMatchAll(v.Name(), namesubnot){
					fmt.Println(path + v.Name())
				}
			}
			if rFlag{
				p := path + v.Name()
				if p != ""{
					p = p + "\\"
				}
				recu(cur, p, ch1, ch2, false)
			}
		}else{
			if !dFlag{
				if my.MatchAll(v.Name(), namesub) && my.NotMatchAll(v.Name(), namesubnot){
					if inF || innF{
						count++
						fileCheck(cur + "\\" +  path + v.Name(), ch2)
						
					}else{
						fmt.Println(path + v.Name())
					}
				}
			}
		}
	}
	ch3 <- count
	ch1 <- first
}

func fileCheck(name string, ch2 chan string){
	str, _ := my.ReadAssets(name)
	arr := strings.Split(str, "\n")
	str = ""
	for n, v := range arr{
		if my.MatchAll(v, textsub) && my.NotMatchAll(v, textsubnot){
			line := strings.Replace(v, "\t", " ", -1)
			num := strconv.Itoa(n + 1)
			str =  str + num + "\t" + line + "\n"
		}
	}
	if str != ""{
		disp := "▼[" + name + "]\n" +
		str[:len(str) - 1] + 
		"\n▲\n\n"
		ch2 <- disp
	}else{
		ch2 <- ""
	}
}
