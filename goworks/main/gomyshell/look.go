package main

import(
	"fmt"
	"io/ioutil"
	"flag"
	"./my"
	"strings"
	"strconv"
	"os"
	"path/filepath"
)

var(
	nFlag bool
	rFlag bool
	dFlag bool
	fFlag bool
	namesub []string
	namesubnot []string
	linesub []string
	linesubnot []string
	diresub []string
	diresubnot []string
	nameF bool
	lineF bool
	direF bool
	count int
)
func main(){
	rF := flag.Bool("r", false, "!!!表示・処理過多注意!!!\n" +
		"\tRecursive 下層のファイル/ディレクトリをすべて検索")
	dF := flag.Bool("d", false, "Directory ディレクトリ検索モード")
	fF := flag.Bool("f", false, "FullPath 絶対パス表示モード")
	nF := flag.Bool("n", false, "!!!表示過多注意!!!\n" +
		"\tnon-stop ファイル内検索時にまとめて表示")
	direS := flag.String("dire", "", "ディレクトリ名検索。ヒットしたディレクトリの直下しか表示しません。\n" +
		"\tディレクトリ名に対して正規表現で検索をかけます。\n" +
		"\tカレントディレクトリは検索、表示対象から外されます。\n" +
		"\t自動的に-rが有効になります。\n" +
		"\tスペースは必ず\\sを指定してください。スペース区切りはAND")
	direnS := flag.String("diren", "", "-direの否定検索版")
	nameS := flag.String("name", "", "ファイル名検索。正規表現で検索します。\n" +
		"\tスペースは必ず\\sを指定してください。スペース区切りはAND")
	namenS := flag.String("namen", "", "-nameの否定検索版")
	lineS := flag.String("line", "", "!!!表示・処理過多注意!!!\n" +
		"\tファイル内検索。正規表現で検索します。\n" +
		"\tスペースは必ず\\sを指定してください。スペース区切りはAND\n" +
		"\t-fileを指定することで読み込むファイルを減らそう！")
	linenS := flag.String("linen", "", "-lineの否定検索版")
	fromS := flag.String("from", ".", "検索を開始するディレクトリ")
	
	flag.Parse()
	
	//フラグの受け取り
	rFlag = *rF
	dFlag = *dF
	fFlag = *fF
	nFlag = *nF
	diresub = strings.Split(*direS, " ")
	diresubnot = strings.Split(*direnS, " ")
	namesub = strings.Split(*nameS, " ")
	namesubnot = strings.Split(*namenS, " ")
	linesub = strings.Split(*lineS, " ")
	linesubnot = strings.Split(*linenS, " ")
	//フラグの有無判定
	nameF = (*nameS != "") || (*namenS != "")
	lineF = (*lineS != "") || (*linenS != "")
	direF = (*direS != "") || (*direnS != "")
	
	//フラグエラー
	if dFlag && (lineF || nFlag){
		fmt.Println("<-d(ディレクトリモード)を使用した場合は、-line/-linen(ファイル内検索)及び-n(ファイル内容連続表示)はすべて無効です>")
		lineF = false
		nFlag = false
	}
	if !lineF && nFlag{
		fmt.Println("<-line/-linenを指定していないので-nは無効です>\n")
	}
	if !rFlag && direF{
		rFlag = true
		fmt.Println("<-direによって-r(再起モード)が自動的に有効になります。>\n")
	}
	//ファイル内検索指定ありで拡張子許可
	if lineF{
		namesub = append(namesub, "(\\.txt$|\\.java$|\\.go$|\\.c$|\\.cpp$|\\.bat$|\\.html$|\\.css$|\\.js$)")
	}
	
	//ファイル探索開始(最後に\がなかったら付ける)
	cur := strings.Replace(*fromS, "\\", "/", 0)
	if cur[len(cur) - 1:] != "/"{
		cur = cur + "/"
	}
	
	recu(cur, "", !direF)
	
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
func recu(cur string, path string, dispflag bool){
	fds, _ := ioutil.ReadDir(cur + path)
	for _, v := range fds{
		if v.IsDir(){
			if dFlag && dispflag{//ディレクトリ検索
				if !nameF || my.MatchAll(v.Name(), namesub) && my.NotMatchAll(v.Name(), namesubnot){
					disp := cur + path + v.Name()
					if fFlag{//絶対パス化
						disp, _ = filepath.Abs(disp)
					}
					fmt.Println(disp)
					count++
				}
			}
			if rFlag{//-rで再帰
				p := path + v.Name()
				if p != ""{
					p = p + "/"
				}
				if !direF || my.MatchAll(v.Name(), diresub) && my.NotMatchAll(v.Name(), diresubnot){
					recu(cur, p, true)
				}else{
					recu(cur, p, false)
				}
			}
		}else{
			if !dFlag && dispflag{//ファイル検索
				if !nameF || my.MatchAll(v.Name(), namesub) && my.NotMatchAll(v.Name(), namesubnot){
					disp := cur + path + v.Name()
					if fFlag{//絶対パス化
						disp, _ = filepath.Abs(disp)
					}
					if lineF{
						fileCheck(disp)
					}else{
						fmt.Println(disp)
						count++
					}
				}
			}
		}
	}
}

//ファイル内文字列を探索して、発見したらfiledispを呼び出す
func fileCheck(name string){
	filestr, _ := ioutil.ReadFile(name)
	str, _ := my.AutoEnc(string(filestr))
	arr := strings.Split(str, "\n")
	str = ""
	for n, v := range arr{
		if !lineF || my.MatchAll(v, linesub) && my.NotMatchAll(v, linesubnot){
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
				fmt.Println("次を表示しない>skip(s)")
				fmt.Println("中断>exit(e)")
				fmt.Println("まとめて表示>all(a)")
				fmt.Println("次の内容を表示>Enter")
				first = false
			}
		}
		fmt.Println()
		fmt.Print("▼[" + name + "]")
		if !nFlag{
			var s string
			fmt.Scanln(&s)
			if s == "exit" || s == "e"{
				os.Exit(1)
			}
			if s == "all" || s == "a"{
				nFlag = true
			}
			if s == "skip" || s == "s"{
				return
			}
		}else{
			fmt.Println()
		}
		
		fmt.Println(str[:len(str) - 1])
		fmt.Println("▲")
		count++
	}
}