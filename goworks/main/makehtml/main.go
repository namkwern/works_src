package main
import(
	"fmt"
	"strings"
)

func main(){
	//mobileAct := html("Mobile Act NAGOYA#2に参加", cssTemp1, body())
	//benkyou := html("9月度勉強会に参加", cssTemp1, body2())
	benkyou := html("夏休みの感想", cssTemp1, body3())
	fmt.Print(benkyou)
}
//全体の作成　置換で作成する
func html(title, css, body string) string{
	h := htmlTemp
	h = strings.Replace(h, "---titleReplace---", title, 1)
	h = strings.Replace(h, "---cssReplace---", css, 1)
	h = strings.Replace(h, "---bodyReplace---", body, 1)
	return h
}
//ボディー作成
func body() string{
	body := header("Mobile Act NAGOYA #2")
	body += mainCont
	return body
}
func body2() string{
	body := header("9月度勉強会")
	body += mainCont2
	return body
}
func body3() string{
	body := header("夏休み 感想")
	body += mainCont3
	return body
}
//上部固定のタイトル要素を書く
func header(top string) string{
	head := `<div class="flex-container">`
	for _, v := range top{
		if v != ' '{
			head += `<div class="default">` + string(v) + `</div>` + "\n"
		}else{
			head += `<div class="space">` + " " + `</div>` + "\n"
		}
		
	}
	return head + `</div>`
}
//main content
func mainContent() string{
	return mainCont
}
/*
func allTag(ss ...string, start, end string) string{
	rv := ""
	for _, v = range ss{
		rv += start + v + end + "\n"
	}
}
*/
var htmlTemp = `
<html>
<head>
<title>---titleReplace---</title>
<style type="text/css">
---cssReplace---
</style>
</head>
<body>
---bodyReplace---
</body>
</html>
`

var cssTemp1 = `
body{
	background-color:#99ff00;
}
.flex-container{
	display:flex;
	flex-direction:row;
	flex-wrap:wrap;
}
.default{
	width:40px;
	height:40px;
	background-color:#ffff00;
	font-size:40px;
	text-align:center;
	margin-right:2px;
	margin-bottom:2px;

	-webkit-border-radius:2px;
	-moz-border-radius:2px;
	border-bottom-style:solid;
	border-right-style:solid;
}
.space{
	width:40px;
	height:40px;
	margin-right:2px;
	margin-bottom:2px;

	-webkit-border-radius:2px;
	-moz-border-radius:2px;
	border-top-style:solid;
	border-left-style:solid;
}
ul li{
	font-size:200%;
}
ul{
	padding:0;
}
.text{
	font-weight:bold;
	font-size:150%;

	background-color:#ffffdd;
	-webkit-border-radius:10px;
	-moz-border-radius:10px;
	padding:15px;
	margin:10px 0;
	border-style:solid;
}

:visited{
	color:#0000ff;
}
:hover{
	background-color:#cccccc;
	border-style:none;
}
`

var mainCont = `
<ul>
<li>Mobile Act NAGOYAとは？</li>
<pre class="text">モバイルアプリ開発について気軽に話し合いをする場です。
モバイル開発をしている企業が主催しており、モバイルに関する技術の調査・発表や、
エンジニア同士、社会人同士の交流を主に目的としています。
#2は9/30に開催されました。
</pre>
<hr>
<li>難航した準備</li>
<pre class="text">
PC組み立ての授業の関係で、前日にならないと当日参加可能かわからなかったので、
急きょ準備をしたため難航しました。
</pre>
<li>LTに挑戦!!</li>
<pre class="text">
LT(ライトニング・トーク)とは５分間ほどのスピーチのことです。
gomobileという技術に注目して、その利点や特徴などについて発表しました。
<a href="https://docs.google.com/presentation/d/1ELxNqwM_V1eKrCFioENc6tr_SFJVygrymdtPhL3o680/edit?usp=sharing">使用したスライド</a>
</pre>
<li>発表後の懇親会</li>
<pre class="text">
全員の発表後、懇親会が開かれました。
参加者は20人弱で、皆がモバイル技術について話し込んでいました。
発表者として参加したため、話題の中心になりやすかったことが発表してよかったことでしょうか。
</pre>
</ul>
`

var mainCont2 = `
<ul>
<li>9月度勉強会とは？</li>
<pre class="text">
技術者としての好奇心や向上心を満たすために技術的な挑戦をするふりをして遊びます。
類義語 : 技術の無駄遣い
</pre>
<hr>
<li>実態</li>
<pre class="text">
実態は、みんな黙々と勉強していました。
それでも、それぞれ好きな分野の勉強だったので実質遊んでいるようなものです。
</pre>
<li>感想</li>
<pre class="text">
時々わからないことを聞きあったり、得意な技術を紹介しあったりします。
新しい技術の風が入ってくるので、モチベーションを保つのにも一役買ってくれます。
</pre>
</ul>
`

var mainCont3 = `
<ul>
<li>夏休みの感想</li>
<pre class="text">
・<a href="MobileActNAGOYA.html">Mobile Act NAGOYA</a>
モバイルアプリ開発についてLTしました。

・<a href="9月度勉強会.html">9月度勉強会</a>
名古屋の勉強会に参加しました。
</pre>
</ul>
`