package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("/home/storyknight/dev/advent/2023/day1/data.txt")

	if err != nil {
		fmt.Println(err)
	}

	fs := bufio.NewScanner(readFile)
	fs.Split(bufio.ScanLines)

	fs.Split(bufio.ScanLines)

	total := 0

	for fs.Scan() {

		line := fs.Text()

		var sb strings.Builder

		for _, c := range line {
			if num, err := strconv.Atoi(string(c)); err == nil {
				fmt.Println("this is c: ", num)
				sb.WriteString(string(rune(num)))
			}

		}

		first := sb.String()[0]
		fmt.Println("full: ", sb)
		fmt.Println("first: ", first)
		last := sb.String()[len(sb.String())-1]
		fmt.Println("last: ", last)

		// integer := byte(first)

		total += (int(first) * 10) + int(last)

		sb.Reset()
	}
	readFile.Close()
	fmt.Println("total: ", total)
}
