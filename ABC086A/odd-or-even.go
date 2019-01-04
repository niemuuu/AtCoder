/*
ABC086A - Product
https://atcoder.jp/contests/abs/tasks/abc086_a
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func readLine() string {
	if !sc.Scan() {
		return ""
	}
	return sc.Text()
}

func main() {
	txt := strings.Split(readLine(), " ")
	n1, _ := strconv.Atoi(txt[0])
	n2, _ := strconv.Atoi(txt[1])

	var result string
	sum := n1 * n2
	if sum%2 == 0 {
		result = "Even"
	} else {
		result = "Odd"
	}

	fmt.Printf("%s\n", result)
}
