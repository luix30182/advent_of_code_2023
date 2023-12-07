package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var gameDelimiter = ":"
var setDelimiter = ";"
var valueDelimiter = ","

var redCubes = 12
var greenCubes = 13
var blueCubes = 14

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

		gameAndSets := strings.Split(line, gameDelimiter)
		sets := strings.Split(gameAndSets[1], setDelimiter)

		var red int
		var green int
		var blue int

		for _, set := range sets {
			colors := strings.Split(set, valueDelimiter)

			for _, color := range colors {
				nameValue := strings.Fields(color)
				value, _ := strconv.Atoi(nameValue[0])
				name := nameValue[1]

				if name == "red" && value > red {
					red = value
				}

				if name == "green" && value > green {
					green = value
				}

				if name == "blue" && value > blue {
					blue = value
				}
			}

		}

		power := red * green * blue

		total += power

	}

	fmt.Println(total)
}
