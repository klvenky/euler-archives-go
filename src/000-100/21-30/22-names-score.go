package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("p022_names.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
	names := strings.Split(string(data), ",")
	sort.Strings(names)
}
