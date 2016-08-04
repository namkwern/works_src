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
	bottomCount int
)
func main(){
	hF := flag.Bool("h", false, "ヘルプを表示する")
	rF := flag.Bool("r", false, "!!!表示・処理過多注意!!!\n" +
		"\tRecursive 下層のファイル/ディレクトリをすべて検索")
	dF := flag.Bool("d", false, "Directory ディレクトリ検索モード")
	fF := flag.Bool("f", false, "FullPath 絶対パス表示モード")
	nF := flag.Bool("n", false, "!!!表示過多注意!!!\n" +
		"\tnon-stop 行検索時にまとめて表示")
	direS := flag.String("dire", "", "末端ディレクトリ名検索。ヒットしたディレクトリの直下しか表示しません。\n" +
		"\tディレクトリ名に対して正規表現で検索をかけます。\n" +
		"\tカレントディレクトリは検索、表示対象から外されます。\n" +
		"\t自動的に-rが有効になります。\n" +
		"\tスペースは必ず\\sを指定してください。スペース区切りはAND")
	direnS := flag.String("diren", "", "-direの否定検索版")
	nameS := flag.String("name", "", "ファイル名検索。正規表現で検索します。\n" +
		"\tスペースは必ず\\sを指定してください。スペース区切りはAND")
	namenS := flag.String("namen", "", "-nameの否定検索版")
	lineS := flag.String("line", "", "!!!表示・処理過多注意!!!\n" +
		"\t行検索。正規表現で検索します。\n" +
		"\tスペースは必ず\\sを指定してください。スペース区切りはAND\n" +
		"\t-nameを指定することで読み込むファイルを減らそう！")
	linenS := flag.String("linen", "", "-lineの否定検索版")
	fromS := flag.String("from", ".", "検索を開始するディレクトリ")
	
	//-hでフラグ詳細表示
	flag.Parse()
	if *hF || len(flag.Args()) != 0{
		fmt.Println("<ファイル検索ツール>")
		fmt.Println("例:下層のjavaファイルからpublic要素を検索する")
		fmt.Println("\tlook -r -name=\"java$\" -line=\"public\"")
		fmt.Println()
		flag.PrintDefaults()
		return
	}
	
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
		fmt.Println("<-d(ディレクトリモード)を使用した場合は、-line/-linen(行検索)及び-n(連続表示)はすべて無効です>")
		lineF = false
		nFlag = false
	}
	if !lineF && nFlag{
		fmt.Println("<-line/-linen(行検索)を指定していないので-n(連続表示)は無効です>\n")
	}
	if !rFlag && direF{
		fmt.Println("<-r(再起モード)を指定していないので-dire(末端ディレクトリ名検索)は使用できません。>\n")
		return
	}
	//ファイル内検索指定ありで拡張子許可
	if lineF{
		namesub = append(namesub, "(\\.txt$|\\.java$|\\.go$|\\.c$|\\.cpp$|\\.bat$|\\.html$|\\.css$|\\.js$)")
	}
	
	//ファイル探索開始(最後に\がなかったら付ける)
	cur := strings.Replace(*fromS, "\\", "/", -1)
	if cur[len(cur) - 1] != '/'{
		cur = cur + "/"
	}
	err := os.Chdir(cur)
	if err != nil{
		fmt.Println("-from指定エラー")
		return
	}
	disp, _ := filepath.Abs(cur)
	
	if *fromS != "."{
		fmt.Println("< from = " + disp + " >\n")
	}
	
	recu("./", !direF)
	
	//個数表示
	var fileType string
	if dFlag{
		fileType = "Directorys"
	}else{
		fileType= "Files"
	}
	fmt.Print("\ndone.\n")
	fmt.Println(count, fileType)
	if dFlag && count != bottomCount{
		fmt.Println(bottomCount, "Bottom Directorys")
	}
}

//カレントディレクトリ、ファイルパス
func recu(path string, dispflag bool) bool{
	bottom := true//末端ディレクトリを判定
	fds, err := ioutil.ReadDir(path)
	if err != nil{
		fmt.Println("err:notfound dir")
	}
	for _, v := range fds{
		if v.IsDir(){
			bottom = false
			bt := false
			if rFlag{//-rで再帰
				b := !direF || my.MatchAll(v.Name(), diresub) && my.NotMatchAll(v.Name(), diresubnot)//ヒットしたディレクトリは表示許可
				bt = recu(path + v.Name() + "/", b)
			}
			if !rFlag || bt{//末端のディレクトリか、-r未指定
				if dFlag && dispflag{//ディレクトリ検索
					if !nameF || my.MatchAll(v.Name(), namesub) && my.NotMatchAll(v.Name(), namesubnot){
						disp := path + v.Name()
						if fFlag{//-fで絶対パス化
							disp, _ = filepath.Abs(disp)
						}
						fmt.Println(disp)
						bottomCount++
					}
				}
			}
			if dFlag{
				count++
			}
		}else{
			if !dFlag && dispflag{//-dなしでファイル検索
				if !nameF || my.MatchAll(v.Name(), namesub) && my.NotMatchAll(v.Name(), namesubnot){
					disp := path + v.Name()
					if fFlag{//-fで絶対パス化
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
	return bottom
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

//ファイル内容を表示
var first = true
func filedisp(name string, str string){
	if str != ""{
		if !nFlag{
			if first{//最初のメッセージ
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
			if s == "skip" || s == "s"{
				return
			}
			if s == "exit" || s == "e"{
				os.Exit(1)
			}
			if s == "all" || s == "a"{
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
