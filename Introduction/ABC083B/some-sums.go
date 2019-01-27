/*
ABC083B - Some Sums
https://atcoder.jp/contests/abs/tasks/abc083_b
*/

package main

import (
	"bufio"
	"fmt"
	"log"
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
	var sums int
	input := strings.Split(readLine(), " ")
	n, _ := strconv.Atoi(input[0])
	min, _ := strconv.Atoi(input[1])
	max, _ := strconv.Atoi(input[2])

	for i := 0; i <= n; i++ {
		var sum int
		split := strings.Split(strconv.Itoa(i), "")

		for _, s := range split {
			num, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			sum += num
		}

		if min <= sum && sum <= max {
			sums += i
		}
	}

	fmt.Printf("%d\n", sums)
}
