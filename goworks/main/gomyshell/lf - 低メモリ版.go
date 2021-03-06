package main

import(
	"fmt"
	"io/ioutil"
	"flag"
	"strings"
	"strconv"
	"os"
	"path/filepath"
	"regexp"
	"./my"
)

type reglist struct{
	contain []*regexp.Regexp	//検索リスト
	not []*regexp.Regexp		//否定検索リスト
}
//それぞれ正規表現で引数から与えられた文字列をチェックするメソッド　上から両方、含む、含まず
func (r reglist)check(str string) (bool){
	fo := len(r.contain) == 0 || my.MatchAll(str, r.contain, true)
	fn := len(r.not) == 0 || my.MatchAll(str, r.not, false)
	return fo && fn
}
func (r reglist)checkContain(str string) (bool){
	fo := len(r.contain) == 0 || my.MatchAll(str, r.contain, true)
	return fo
}
func (r reglist)checkNot(str string) (bool){
	fn := len(r.not) == 0 || my.MatchAll(str, r.not, false)
	return fn
}

var(
	nFlag bool		//-n
	rFlag bool		//-r
	dFlag bool		//-d
	fFlag bool		//-f
	tFlag bool		//-t
	usFlag bool		//-unsafe
	namesub reglist	//-name(n)=""　文字列
	linesub reglist	//-line(n)=""
	diresub reglist	//-dire(n)=""
	nameF bool		//-name(n)　有無
	lineF bool		//-line(n)
	direF bool		//-dire(n)
	count int
	bottomCount int
)

func main(){
	hF := flag.Bool("h", false, 
		"ヘルプを表示する",
	)
	flag.BoolVar(&rFlag, "r", false, 
		"!!!表示・処理過多注意!!!\n" +
		"\tRecursive 下層のファイル/ディレクトリをすべて検索",
	)
	flag.BoolVar(&dFlag, "d", false, 
		"Directory ディレクトリ検索モード",
	)
	flag.BoolVar(&fFlag, "f", false, 
		"FullPath 絶対パス表示モード",
	)
	flag.BoolVar(&nFlag, "n", false, 
		"!!!表示過多注意!!!\n" +
		"\tnon-stop 行検索時にまとめて表示",
	)
	flag.BoolVar(&tFlag, "t", false, 
		"Time 更新日時表示モード(未実装です)",
	)
	flag.BoolVar(&usFlag, "unsafe", false,
		"アンセーフ バイナリファイル等を判断せず表示する",
	)
	direS := flag.String("dire", "", 
		"ディレクトリ名検索。ヒットしたディレクトリより下層しか表示しません。\n" +
		"\tディレクトリ名に対して正規表現で検索をかけます。\n" +
		"\tカレントディレクトリは検索、表示対象から外されます。\n" +
		"\t実行には必ず-rが必要です。\n" +
		"\tスペースは必ず\\sを指定してください。スペース区切りはAND",
	)
	direnS := flag.String("diren", "", 
		"-direの否定検索版",
	)
	nameS := flag.String("name", "", 
		"ファイル名検索。正規表現で検索します。\n" +
		"\tスペースは必ず\\sを指定してください。スペース区切りはAND",
	)
	namenS := flag.String("namen", "", 
		"-nameの否定検索版",
	)
	lineS := flag.String("line", "", 
		"!!!表示・処理過多注意!!!\n" +
		"\t行検索。正規表現で検索します。\n" +
		"\tスペースは必ず\\sを指定してください。スペース区切りはAND\n" +
		"\t-nameを指定することで読み込むファイルを減らそう！",
	)
	linenS := flag.String("linen", "", 
		"-lineの否定検索版",
	)
	fromS := flag.String("from", ".", 
		"検索を開始するディレクトリを指定する",
	)
	
	//-hでフラグ詳細表示
	flag.Parse()
	if *hF || len(flag.Args()) != 0{
		fmt.Println("<ファイル検索ツール>")
		fmt.Println("例:下層のjavaファイルからpublic要素を検索する")
		fmt.Println("\tlf -r -name=\"java$\" -line=\"public\"")
		fmt.Println()
		flag.PrintDefaults()
		return
	}
	
	//フラグの受け取り
	reg, _ := regexp.Compile("\\s+")
	diresub = reglist{
		contain:	my.CompileAll(reg.Split(*direS, -1)),
		not:		my.CompileAll(reg.Split(*direnS, -1)),
	}
	namesub = reglist{
		contain:	my.CompileAll(reg.Split(*nameS, -1)),
		not:		my.CompileAll(reg.Split(*namenS, -1)),
	}
	linesub = reglist{
		contain:	my.CompileAll(reg.Split(*lineS, -1)),
		not:		my.CompileAll(reg.Split(*linenS, -1)),
	}
	//フラグの有無判定
	direF = len(diresub.contain) != 0 || len(diresub.not) != 0
	nameF = len(namesub.contain) != 0 || len(namesub.not) != 0
	lineF = len(linesub.contain) != 0 || len(linesub.not) != 0
	
	//フラグ組み合わせなどのエラー
	switch{
	case dFlag && (lineF || nFlag):
		fmt.Println("<-d(ディレクトリモード)を使用した場合は、-line/-linen(行検索)及び-n(連続表示)はすべて無効です>")
		return
	case !lineF && nFlag:
		fmt.Println("<-line/-linen(行検索)を指定していないので-n(連続表示)は無効です>\n")
		return
	case !rFlag && direF:
		fmt.Println("<-r(再起モード)を指定していないので-dire(ディレクトリ名検索)は使用できません。>\n")
		return
	}
	//エラー以外のフラグ処理
	if lineF && !usFlag{				//ファイル内検索指定ありで拡張子許可
		namesub.contain = append(namesub.contain, regexp.MustCompile("(\\.txt$|\\.java$|\\.go$|\\.c$|\\.cpp$|\\.bat$|\\.html$|\\.css$|\\.js$)"))
		namesub.not = append(namesub.not, my.CompileAll([]string{"\\.exe$", "\\.dll$"})...)
	}
	if lineF && !nFlag{
		fmt.Println("次を表示しない>skip(s)")
		fmt.Println("中断>exit(e)")
		fmt.Println("まとめて表示>all(a)")
		fmt.Println("次の内容を表示>Enter")
		fmt.Println()
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
	
	if *fromS != "."{
		disp, _ := os.Getwd()
		fmt.Println("< from = " + disp + " >\n")
	}
	
	recu("./")
	fmt.Print("\ndone.\n")
	
	//個数表示
	if dFlag{
		fmt.Println(count, "Directorys")
		if bottomCount > count{
			count = bottomCount
		}
		if !(nameF || direF) && count != bottomCount{
			fmt.Println(bottomCount, "Bottom Directorys")
		}
	}else{
		fmt.Println(count, "Files")
	}
	
}

//再帰的にディレクトリを探索する関数
//引数1:カレントディレクトリ
//戻り値1:末端ディレクトリ判定(再帰時に使用)
func recu(path string) bool{
	bottom := true//末端ディレクトリを判定
	fds, _ := ioutil.ReadDir(path)
	for _, v := range fds{
		if v.IsDir(){
			bottom = false
			func(){
				switch false{
					case nameF || direF: return				//検索用フラグ使用
					case dFlag: return						//ディレクトリ検索
					case namesub.check(v.Name()): return	//名前と正規表現の一致を検証
					case diresub.checkContain(path): return	//パスと正規表現の一致を検証
				}
				fmt.Println(pathFormat(v, path))
				count++
			}()
			bt := func() (bt bool){
				switch false{
					case rFlag: return						//-rで再帰
					case diresub.checkNot(v.Name()): return	//-direnで指定されたディレクトリを除外
				}
				return recu(path + v.Name() + "/")
			}()
			func(){
				switch false{
					case !(nameF || direF): return	//検索用フラグ不使用
					case dFlag: return				//ディレクトリ検索
				}
				count++
				switch false{
					case !rFlag || bt: return		//末端のディレクトリか、-r未指定(通常の表示)
				}
				fmt.Println(pathFormat(v, path))
				bottomCount++
			}()
		}else{
			switch false{
				case !dFlag: continue						//ファイル検索
				case namesub.check(v.Name()): continue		//名前と正規表現の一致を検証
				case diresub.checkContain(path): continue	//パスと正規表現の一致を検証
			}
			disp := pathFormat(v, path)
			if lineF{										//-line等を使用してファイルの内部を参照
				fileCheck(pathFormat(v, path), disp)
			}else{
				fmt.Println(disp)
				count++
			}
		}
	}
	return bottom
}

//絶対パスにしたりする
//引数1:ファイル情報
//引数2:相対パス
//戻り値2:加工済み文字列
func pathFormat(v os.FileInfo, path string) (name string){
	name = path + v.Name()
	if fFlag{
		name, _ = filepath.Abs(name)
	}
	return
}

//ファイルを開いて内部を参照し、正規表現でチェックする処理
//ファイル内文字列を探索して、発見したらfiledispを呼び出す
//引数1:ファイル名
//引数2:表示テキスト
func fileCheck(file, disp string){
	filestr, _ := ioutil.ReadFile(file)
	str, _ := my.AutoEnc(string(filestr))
	arr := strings.Split(str, "\n")
	str = ""
	for n, v := range arr{
		switch false{
			case linesub.check(v): continue
		}
		line := strings.Replace(v, "\t", " ", -1)
		num := strconv.Itoa(n + 1)
		if len(line) <= 300{
			str += num + "\t" + line + "\n"
		}else{
			str += num + "\t" + line[:300] + "(ry" + "\n"
		}
	}
	filedisp(disp, str)
}

//ファイル内容を表示
//引数1:ファイルとして表示する文字列
//引数2:ファイル内文字列
func filedisp(name, str string){
	if str == ""{
		return
	}
	fmt.Print(name, " ->")
	if !nFlag{
		var s string
		fmt.Scanln(&s)
		switch(s){
		case "skip", "s":
			return
		case "exit", "e":
			os.Exit(1)
		case "all", "a":
			nFlag = true
		}
	}else{
		fmt.Println()
	}
	fmt.Println(str[:len(str) - 1])
	fmt.Println("[EOF]\n")
	count++
}