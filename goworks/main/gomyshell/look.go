package main

import(
	"fmt"
	"io/ioutil"
	"flag"
	"./my"
	"strings"
	"strconv"
	"os"
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
	count int
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
	recu(*fromS, "")
	
	var fileType string
	if dFlag{
		fileType = "Directorys"
	}else{
		fileType= "Files"
	}
	fmt.Print("\ndone.\n")
	fmt.Println(count, fileType)
}

//カレントディレクトリ、ファイルパス
func recu(cur string, path string){
	fds, _ := ioutil.ReadDir(cur + "\\" + path)
	for _, v := range fds{
		if v.IsDir(){
			if dFlag{
				if my.MatchAll(v.Name(), namesub) && my.NotMatchAll(v.Name(), namesubnot){
					fmt.Println(path + v.Name())
					count++
				}
			}
			if rFlag{
				p := path + v.Name()
				if p != ""{
					p = p + "\\"
				}
				recu(cur, p)
			}
		}else{
			if !dFlag{
				if my.MatchAll(v.Name(), namesub) && my.NotMatchAll(v.Name(), namesubnot){
					if inF || innF{
						fileCheck(cur + "\\" +  path + v.Name())
					}else{
						fmt.Println(path + v.Name())
						count++
					}
				}
			}
		}
	}
}

func fileCheck(name string){
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
	filedisp(name, str)
}

var first = true
func filedisp(name string, str string){
	if str != ""{
		if !nFlag{
			if first{
				fmt.Println("中断>exit\nまとめて表示>all\n次の内容を表示>Enter")
				first = false
			}
		}
		fmt.Println()
		fmt.Print("▼[" + name + "]")
		if !nFlag{
			var s string
			fmt.Scanln(&s)
			if s == "exit"{
				os.Exit(1)
			}
			if s == "all"{
				nFlag = true
			}
		}else{
			fmt.Println()
		}
		
		fmt.Println(str[:len(str) - 1])
		fmt.Println("▲")
		count++
	}
}