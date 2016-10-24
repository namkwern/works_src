package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("start.")
	file, _ := os.Open("./data.txt")
	b, _ := ioutil.ReadAll(file)
	str := jsonParse(string(b), "class")
	fmt.Println(str)

	fmt.Println("done.")
}
func jsonParse(str string, key string) string {
	var data interface{}
	decoder := json.NewDecoder(strings.NewReader(str))
	decoder.Decode(&data)
	gr := map[string]string{}
	assert(data, key, gr)
	s := "["
	first := false
	for k, v := range gr {
		if first {
			s += ","
		}
		first = true
		s += "\n" + `"` + k + `":{`
		s += v + "\n}"
	}
	s += "\n]"
	return s
}

func assert(data interface{}, key string, gr map[string]string) (str string) {
	switch data.(type) {
	case string:
		str += `"` + fmt.Sprint(data) + `"` + "\n"
	case nil:
		str += fmt.Sprintln("null")
	case []interface{}:
		for _, v := range data.([]interface{}) {
			str += assert(v, key, gr)
		}
	case map[string]interface{}:
		curkey := fmt.Sprint(data.(map[string]interface{})[key])
		_, ok := gr[curkey]
		if ok {
			gr[curkey] += ","
		}
		gr[curkey] += "\n{\n"
		for k, v := range data.(map[string]interface{}) {
			if k == key {
				continue
			}
			gr[curkey] += `"` + k + `":` + assert(v, key, gr)
		}
		gr[curkey] += "}"
	default:
		str += fmt.Sprintln(data)
	}
	return str
}
