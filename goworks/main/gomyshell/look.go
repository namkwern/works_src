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
	namesubnot []string
	linesub []string
	linesubnot []string
	diresub []string
	diresubnot []string
	lineF bool
	nFlag bool
	rFlag bool
	count int
)
func main(){
	
	rF := flag.Bool("r", false, "!!!表示・処理過多注意!!!\n\tRecursive 下層のファイル/ディレクトリをすべて検索")
	dF := flag.Bool("d", false, "Directory ディレクトリ検索に切り替え")
	nF := flag.Bool("n", false, "!!!表示過多注意!!!\n\tnon-stop ファイル内検索時にまとめて表示")
	direS := flag.String("dire", "", "ディレクトリ名検索。パスに対して検索をかけます。-dか-rで検索する必要があります。")
	direnS := flag.String("diren", "", "-direの否定検索版")
	fileS := flag.String("file", "", "ファイル名検索。正規表現で検索します。\n\tスペースは必ず\\sを指定してください。スペース区切りはAND")
	filenS := flag.String("filen", "", "fileの否定検索版")
	lineS := flag.String("line", "", "!!!表示・処理過多注意!!!\n\tファイル内検索。正規表現で検索します。\n\tスペースは必ず\\sを指定してください。スペース区切りはAND\n\t-fileを指定することで読み込むファイルを減らそう！")
	linenS := flag.String("linen", "", "-inの否定検索版")
	fromS := flag.String("from", ".", "検索を開始するディレクトリ")
	
	flag.Parse()
	
	//フラグの受け取り
	rFlag = *rF
	dFlag = *dF
	nFlag = *nF
	diresub = strings.Split(*direS, " ")
	diresubnot = strings.Split(*direnS, " ")
	namesub = strings.Split(*fileS, " ")
	namesubnot = strings.Split(*filenS, " ")
	linesub = strings.Split(*lineS, " ")
	linesubnot = strings.Split(*linenS, " ")
	
	lineF = (*lineS != "") || (*linenS != "")
	
	
	//フラグエラー
	if dFlag && (lineF || nFlag){
		fmt.Println("<-d(ディレクトリモード)を使用した場合は、-file/-filen(ファイル名検索)-line/-linen(ファイル内検索)及び-n(ファイル内容連続表示)はすべて無効です>")
		lineF = false
		nFlag = false
	}
	if !lineF && nFlag{
		fmt.Println("<-line/-linenを指定していないので-nは無効です>\n")
	}
	if (!rFlag && !dFlag) && *direS != ""{
		fmt.Println("<-r(再起モード)もしくは-d(ディレクトリモード)を指定していないので-direは無効です。>\n")
	}
	//ファイル内検索指定ありで拡張子許可
	if lineF{
		namesub = append(namesub, "(\\.txt$|\\.java$|\\.go$|\\.c$|\\.cpp$|\\.bat$|\\.html$|\\.css$|\\.js$)")
	}
	
	//ファイル探索開始(最後に\がなかったら付ける)
	cur := *fromS
	if cur[len(cur) - 1:] != "\\"{
		cur = cur + "\\"
	}
	recu(cur, "")
	
	//個数表示
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
	fds, _ := ioutil.ReadDir(cur + path)
	for _, v := range fds{
		if v.IsDir(){
			if dFlag{
				if my.MatchAll(v.Name(), diresub) && my.NotMatchAll(v.Name(), diresubnot){
					fmt.Println(cur + path + v.Name())
					count++
				}
			}
			if rFlag{
				p := path + v.Name()
				if my.MatchAll(p, diresub) && my.NotMatchAll(p, diresubnot){
					if p != ""{
						p = p + "\\"
					}
					recu(cur, p)
				}
			}
		}else{
			if !dFlag{
				if my.MatchAll(v.Name(), namesub) && my.NotMatchAll(v.Name(), namesubnot){
					if lineF{
						fileCheck(cur + path + v.Name())
					}else{
						fmt.Println(cur + path + v.Name())
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
		if my.MatchAll(v, linesub) && my.NotMatchAll(v, linesubnot){
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