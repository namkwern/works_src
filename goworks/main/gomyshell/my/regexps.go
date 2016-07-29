package my

import (
	"fmt"
	"regexp"
)

//すべての正規表現殿完全一致判定 (文字列,正規表現配列)
func MatchAll(str string, arr []string) bool {
	for _, v := range arr {
		reg := regexp.MustCompile(v)
		if !reg.MatchString(str) {
			return false
		}
	}
	return true
}

//すべての正規表現殿完全不一致判定 (文字列,正規表現配列)
func NotMatchAll(str string, arr []string) bool {
	for _, v := range arr {
		reg := regexp.MustCompile(v)
		if v == "" {
			continue
		}
		if reg.MatchString(str) {
			return false
		}
	}
	return true
}

//すべての正規表現殿完全一致判定、判定が通ったらそのまま出力する (文字列,正規表現配列)
func MatchAllPrintln(str string, arr []string) {
	findF := true
	for _, v := range arr {
		reg := regexp.MustCompile(v)
		if !reg.MatchString(str) {
			findF = false
		}
	}
	if findF {
		fmt.Println(str)
	}
}
