package main

import (
	"fmt"
)

func main() {
	ini()
	for {
		list := canPutList()
		disp(list)
		if turn {
			player()
		} else {
			player()
		}
		turn = !turn
	}
}
func player() {
again:
	x, y := getPos()
	if !canPut(x, y) {
		goto again
	}
	put(x, y)
	reverseAll(x, y)
}

const (
	BLACK rune = '〇'
	WHITE rune = '●'
	NONE  rune = '・'
)

var (
	ban  [][]rune
	turn bool
)

func ini() {
	ban = [][]rune{
		{'・', '・', '・', '・', '・', '・', '・', '・'},
		{'・', '・', '・', '・', '・', '・', '・', '・'},
		{'・', '・', '・', '・', '・', '・', '・', '・'},
		{'・', '・', '・', '〇', '●', '・', '・', '・'},
		{'・', '・', '・', '●', '〇', '・', '・', '・'},
		{'・', '・', '・', '・', '・', '・', '・', '・'},
		{'・', '・', '・', '・', '・', '・', '・', '・'},
		{'・', '・', '・', '・', '・', '・', '・', '・'},
	}
}
func disp(list [][]bool) {
	for n, _ := range ban {
		fmt.Printf("%2d", n+1)
	}
	fmt.Println()
	for n, v := range ban {
		fmt.Print(n + 1)
		for n2, v2 := range v {
			if list[n][n2] && v2 == NONE {
				fmt.Print(string('＊'))
			} else {
				fmt.Print(string(v2))
			}
		}
		fmt.Println()
	}
}
func getPos() (int, int) {
	var x, y int
	fmt.Scan(&x, &y)
	return x - 1, y - 1
}
func put(x, y int) {
	if ban[x][y] != NONE {
		return
	}
	if turn {
		ban[x][y] = BLACK
	} else {
		ban[x][y] = WHITE
	}
}

var direction = [][]int{
	{-1, -1}, //左上
	{0, -1},  //上
	{1, -1},  //右上

	{-1, 0}, //左
	//{0, 0},//変化量00はありえない
	{1, 0}, //右

	{-1, 1}, //左下
	{0, 1},  //下
	{1, 1},  //右下
}

func canPutList() (list [][]bool) {
	list = make([][]bool, 8)
	for n, _ := range list {
		list[n] = make([]bool, 8)
	}
	for n := 0; n < len(ban); n++ {
		for m := 0; m < len(ban[n]); m++ {
			if canPut(n, m) {
				list[n][m] = true
			}
		}
	}
	return
}
func canPut(x, y int) bool {
	for _, v := range direction {
		dx := x
		dy := y
		flg := false
		for {
			dx += v[0]
			dy += v[1]
			if !inside(dx, dy) {
				break
			}
			if turn {
				if ban[dx][dy] == WHITE {
					flg = true
					continue
				} else if flg {
					if ban[dx][dy] == BLACK {
						return true
					}
				}
				break
			} else {
				if ban[dx][dy] == BLACK {
					flg = true
					continue
				} else if flg {
					if ban[dx][dy] == WHITE {
						return true
					}
				}
				break
			}
		}
	}
	return false
}

func reverseAll(x, y int) {
	for _, v := range direction {
		dx := x
		dy := y
		flg := false
		for {
			dx += v[0]
			dy += v[1]
			if !inside(dx, dy) {
				break
			}

			if turn {
				if ban[dx][dy] == WHITE {
					flg = true
				} else if flg {
					if ban[dx][dy] == BLACK {
						reverseLine(x+v[0], y+v[1], dx-v[0], dy-v[1])
					}
					break
				}
				if ban[dx][dy] != WHITE {
					break
				}
			} else {
				if ban[dx][dy] == BLACK {
					flg = true
				} else if flg {
					if ban[dx][dy] == WHITE {
						reverseLine(x+v[0], y+v[1], dx-v[0], dy-v[1])
					}
					break
				}
				if ban[dx][dy] != BLACK {
					break
				}
			}
		}
	}
}
func reverseLine(x, y, x2, y2 int) {
	if x > x2 {
		x, x2 = x2, x
	}
	if y > y2 {
		y, y2 = y2, y
	}
	for n := x; n <= x2; n++ {
		for m := y; m <= y2; m++ {
			reverse(n, m)
		}
	}
}
func reverse(x, y int) {
	if ban[x][y] == NONE {
		return
	}
	if ban[x][y] == BLACK {
		ban[x][y] = WHITE
	} else {
		ban[x][y] = BLACK
	}
}
func inside(x, y int) bool {
	return 0 <= x && x < 8 && 0 <= y && y < 8
}
