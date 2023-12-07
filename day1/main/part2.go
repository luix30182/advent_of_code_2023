package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var tokens = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func leftNumber(line string, index int) string {
	n, err := strconv.Atoi(line[index : index+1])
	if err != nil {

	} else {
		return strconv.Itoa(n)
	}
	correct := "0"

	for token, value := range tokens {
		limit := index + len(token)
		if limit <= len(line) {
			txtNumber := line[index:limit]
			if txtNumber == token {
				return value
			}
		}
	}

	return correct
}

func rightNumber(line string, index int) string {
	n, err := strconv.Atoi(line[index-1 : index])
	if err != nil {

	} else {
		return strconv.Itoa(n)
	}
	correct := "0"

	for token, value := range tokens {
		limit := index - len(token)
		if limit >= 0 {
			txtNumber := line[limit:index]
			if txtNumber == token {
				return value
			}
		}
	}

	return correct
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no data file")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		var left string
		var right string

		for i := 0; i <= len(line); i++ {
			left = leftNumber(line, i)
			if left != "0" {
				break
			}
		}
		for i := len(line); i > 0; i-- {
			right = rightNumber(line, i)
			if right != "0" {
				break
			}
		}

		lineTotal, err := strconv.Atoi(left + right)
		if err != nil {

		}
		total += lineTotal
	}

	fmt.Println(total)

}
