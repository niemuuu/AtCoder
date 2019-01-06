/*
ABC088B - Card Game for Two
https://atcoder.jp/contests/abs/tasks/abc088_b
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	n, _ := strconv.Atoi(readLine())
	cards := strings.Split(readLine(), " ")
	nums := make([]int, 0, n)

	for _, card := range cards {
		num, _ := strconv.Atoi(card)
		nums = append(nums, num)
	}
	// sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	var a, b int
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			a += nums[i]
		} else {
			b += nums[i]
		}
	}

	fmt.Printf("%d\n", a-b)
}
