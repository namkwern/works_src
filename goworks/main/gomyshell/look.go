package main

import(
	"fmt"
	"io/ioutil"
	"flag"
	"./my"
	"strings"
	"strconv"
)

var isdir bool
var filesub []string
var textsub []string
var filesubnot []string
var textsubnot []string
var inF bool
var innF bool
var filenum bool
var rec bool
func main(){
	rF := flag.Bool("r", false, "Recursive 下層のファイルすべて")
	dF := flag.Bool("d", false, "Directory ディレクトリ")
	nF := flag.Bool("n", false, "Number ファイル内の表示に行数を追加する")
	forS := flag.String("for", "", "ファイル名表示、正規表現で検索\n\tスペースは必ず\\sを指定してください。スペース区切りはAND")
	fornS := flag.String("forn", "", "forの否定検索版")
	inS := flag.String("in", "", "ファイルの中身表示、正規表現で検索\n\tスペースは必ず\\sを指定してください。スペース区切りはAND\n\t-forを指定することで読み込むファイルを減らそう！")
	innS := flag.String("inn", "", "-inの否定検索版")
	fromS := flag.String("from", ".", "検索を開始するディレクトリ")
	
	flag.Parse()
	
	if *dF && (*inS != "" || *innS != "" || *nF){
		fmt.Println("-d(ディレクトリ指定)を使用した場合は\n-in(ファイル内検索)及び-nは指定できません")
	}
	
	isdir = *dF
	filenum = * nF
	filesub = strings.Split(*forS, " ")
	textsub = strings.Split(*inS, " ")
	filesubnot = strings.Split(*fornS, " ")
	textsubnot = strings.Split(*innS, " ")
	inF = (*inS != "")
	innF = (*innS != "")
	if inF || innF{
		filesub = append(filesub, "(\\.txt$|\\.java$|\\.go$|\\.c$|\\.cpp$|\\.bat$|\\.html$|\\.css$|\\.js$)")
	}
	rec = *rF
	
	recu(*fromS, "")
}

//カレントディレクトリ、ファイルパス
func recu(cur string, path string){
	fds, _ := ioutil.ReadDir(cur + "\\" + path)
	for _, v := range fds{
		if v.IsDir(){
			if isdir{
				if my.MatchAll(v.Name(), filesub) && my.NotMatchAll(v.Name(), filesubnot){
					fmt.Println(path + v.Name())
				}
			}
			if rec{
				p := path + v.Name()
				if p != ""{
					p = p + "\\"
				}
				recu(cur, p)
			}
		}else{
			if !isdir{
				if my.MatchAll(v.Name(), filesub) && my.NotMatchAll(v.Name(), filesubnot){
					if !inF && !innF{
						fmt.Println(path + v.Name())
					}else{
						fileCheck(cur + "\\" +  path + v.Name())
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
			num := ""
			if filenum{
				num = strconv.Itoa(n + 1)
			}
			str =  str + num + "\t" + line + "\n"
		}
	}
	if str != ""{
		fmt.Println("▼[" + name + "]")
		fmt.Println(str)
		fmt.Println("▲[" + name + "]")
	}
}