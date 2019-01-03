// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// var sc = bufio.NewScanner(os.Stdin)

// func readLine() string {
// 	if !sc.Scan() {
// 		return ""
// 	}
// 	return sc.Text()
// }

// func main() {
// 	t1 := readLine()
// 	t2 := strings.Split(readLine(), " ")
// 	t3 := readLine()

// 	n1, _ := strconv.Atoi(t1)
// 	n2, _ := strconv.Atoi(t2[0])
// 	n3, _ := strconv.Atoi(t2[1])
// 	total := n1 + n2 + n3

// 	fmt.Printf("%d %s\n", total, t3)
// }
