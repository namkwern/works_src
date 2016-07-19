package my

import (
	"bufio"
	"net/textproto"
	"golang.org/x/text/encoding/japanese"
    "golang.org/x/text/transform"
    "io"
    "regexp"
    "fmt"
    "io/ioutil"
    "strings"
    "golang.org/x/net/html/charset"
    "bytes"
)

//表示可能なパイプを作成する ReadLine()で行を取得できる
func SjisReader(stdout io.ReadCloser) *textproto.Reader{
	rio := transform.NewReader(stdout, japanese.ShiftJIS.NewDecoder())
	return textproto.NewReader(bufio.NewReader(rio))
}
//stringのsjis→utf8変換
func FromSjis(str string) (string, error) {
    rio := transform.NewReader(strings.NewReader(str), japanese.ShiftJIS.NewDecoder())
    b, err := ioutil.ReadAll(rio)
    return string(b), err
}

//
func ReadAssets(str string) (string, error) {
	body, err := ioutil.ReadFile(str)
	if err != nil {
		return "", err
	}

	var f []byte
	encodings := []string{"sjis", "utf-8"}
	for _, enc := range encodings {
		if enc != "" {
			ee, _ := charset.Lookup(enc)
			if ee == nil {
				continue
			}
			var buf bytes.Buffer
			ic := transform.NewWriter(&buf, ee.NewDecoder())
			_, err := ic.Write(body)
			if err != nil {
				continue
			}
			err = ic.Close()
			if err != nil {
				continue
			}
			f = buf.Bytes()
			break
		}
	}
	return string(f), nil
}

//すべての正規表現殿完全一致判定 (文字列,正規表現配列)
func MatchAll(str string, arr []string) bool{
	for _, v := range arr{
		reg := regexp.MustCompile(v)
		if !reg.MatchString(str){
			return false
		}
	}
	return true
}

//すべての正規表現殿完全不一致判定 (文字列,正規表現配列)
func NotMatchAll(str string, arr []string) bool{
	for _, v := range arr{
		reg := regexp.MustCompile(v)
		if v == ""{
			continue
		}
		if reg.MatchString(str){
			return false
		}
	}
	return true
}

//すべての正規表現殿完全一致判定、判定が通ったらそのまま出力する (文字列,正規表現配列)
func MatchAllPrintln(str string, arr []string){
	findF := true
	for _, v := range arr{
		reg := regexp.MustCompile(v)
		if !reg.MatchString(str){
			findF = false
		}
	}
	if findF{
		fmt.Println(str)
	}
}