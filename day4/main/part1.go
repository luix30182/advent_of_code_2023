package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var cardDataDivider = ":"
var gameDelimiter = "|"

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

	var total float64

	for scanner.Scan() {
		line := scanner.Text()

		var points float64
		var t float64

		cardData := strings.Split(line, cardDataDivider)
		games := strings.Split(cardData[1], gameDelimiter)

		g1 := strings.Fields(games[0])
		g2 := strings.Fields(games[1])

		for _, x := range g1 {
			contains := slices.Contains(g2, x)
			if contains {
				points++

				if points > 1 {
					t = t * 2

				} else {
					t += points
				}
			}
		}

		total += t

	}

	fmt.Println(total)
}
