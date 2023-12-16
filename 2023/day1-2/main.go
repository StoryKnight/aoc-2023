package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
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
		var slice []int

		for index, c := range line {
			if num, err := strconv.Atoi(string(c)); err == nil {
				slice = append(slice, index)
				// fmt.Println("this is c: ", num)
				sb.WriteString(string(rune(num)))
			}
		}

		first := sb.String()[0]
		// fmt.Println("full: ", sb)
		// fmt.Println("first: ", first)
		last := sb.String()[len(sb.String())-1]
		// fmt.Println("last: ", last)

		// integer := byte(first)

		re := regexp.MustCompile(`one`)
		indices := re.FindAllStringIndex(line, -1)
		// [3,6], [7,12]
		// ja2js6sixone
		if indices != nil {
			for _, word := range indices {
				if word[1] > slices.Max(slice) {
					// "42"
					slice = append(slice, word[1])
					last = 1
				}

				if word[0] < slices.Min(slice) {
					slice = append(slice, word[0])
					first = 1
				}
			}

		}

		re = regexp.MustCompile(`two`)
		indices = re.FindAllStringIndex(line, -1)
		if indices != nil {
			for _, word := range indices {
				if word[1] > slices.Max(slice) {
					last = 2
					slice = append(slice, word[1])
				}
				if word[0] < slices.Min(slice) {
					first = 2
					slice = append(slice, word[0])
				}
			}
		}

		re = regexp.MustCompile(`three`)
		indices = re.FindAllStringIndex(line, -1)
		if indices != nil {
			for _, word := range indices {
				if word[1] > slices.Max(slice) {
					last = 3
					slice = append(slice, word[1])
				}
				if word[0] < slices.Min(slice) {
					first = 3
					slice = append(slice, word[0])
				}
			}
		}
		re = regexp.MustCompile(`four`)
		indices = re.FindAllStringIndex(line, -1)
		if indices != nil {
			for _, word := range indices {
				if word[1] > slices.Max(slice) {
					last = 4
					slice = append(slice, word[1])
				}
				if word[0] < slices.Min(slice) {
					first = 4
					slice = append(slice, word[0])
				}
			}
		}

		re = regexp.MustCompile(`five`)
		indices = re.FindAllStringIndex(line, -1)
		if indices != nil {
			for _, word := range indices {
				if word[1] > slices.Max(slice) {
					last = 5
					slice = append(slice, word[1])
				}
				if word[0] < slices.Min(slice) {
					first = 5
					slice = append(slice, word[0])
				}
			}
		}
		re = regexp.MustCompile(`six`)
		indices = re.FindAllStringIndex(line, -1)
		if indices != nil {
			for _, word := range indices {
				if word[1] > slices.Max(slice) {
					last = 6
					slice = append(slice, word[1])
				}
				if word[0] < slices.Min(slice) {
					first = 6
					slice = append(slice, word[0])
				}
			}
		}
		re = regexp.MustCompile(`seven`)
		indices = re.FindAllStringIndex(line, -1)
		if indices != nil {
			for _, word := range indices {
				if word[1] > slices.Max(slice) {
					last = 7
					slice = append(slice, word[1])
				}
				if word[0] < slices.Min(slice) {
					first = 7
					slice = append(slice, word[0])
				}
			}
		}

		re = regexp.MustCompile(`eight`)
		indices = re.FindAllStringIndex(line, -1)
		if indices != nil {
			for _, word := range indices {
				if word[1] > slices.Max(slice) {
					last = 8
					slice = append(slice, word[1])
				}
				if word[0] < slices.Min(slice) {
					first = 8
					slice = append(slice, word[0])
				}
			}
		}

		re = regexp.MustCompile(`nine`)
		indices = re.FindAllStringIndex(line, -1)
		if indices != nil {
			for _, word := range indices {
				if word[1] > slices.Max(slice) {
					last = 9
					slice = append(slice, word[1])
				}
				if word[0] < slices.Min(slice) {
					first = 9
					slice = append(slice, word[0])
				}
			}
		}

		total += (int(first) * 10) + int(last)
		slice = nil
		sb.Reset()
	}
	readFile.Close()
	fmt.Println("total: ", total)
}
