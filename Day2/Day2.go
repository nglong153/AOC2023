package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	numOfRed   int
	numOfBlue  int
	numOfGreen int
}

func Part1(scanner *bufio.Scanner) {
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			fmt.Println("Empty line")
			continue
		}

		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			fmt.Println("Wrong game format")
			return
		}

		gameIdStr := strings.TrimPrefix(parts[0], "Game ")
		gameValid := true

		sets := strings.Split(parts[1], ";")
		for _, set := range sets {
			colorWithNumbers := strings.Split(set, ",")
			for _, turn := range colorWithNumbers {
				tmp := strings.Split(strings.TrimSpace(turn), " ")
				number, err := strconv.Atoi(tmp[0])
				if err != nil {
					return
				}
				switch tmp[1] {
				case "blue":
					if number > 14 {
						gameValid = false
					}
				case "red":
					if number > 12 {
						gameValid = false
					}

				case "green":
					if number > 13 {
						gameValid = false
					}

				}

				if !gameValid {
					break
				}
			}

			if !gameValid {
				break
			}
		}

		if gameValid {
			gameIdNumber, err := strconv.Atoi(gameIdStr)
			if err != nil {
				return
			}
			sum += gameIdNumber
			fmt.Println(gameIdNumber)
		}
	}
	fmt.Println(sum)
}

func Part2(scanner *bufio.Scanner) {
	sum := 0
	for scanner.Scan() {

		maxRed, maxBlue, maxGreen := 0, 0, 0
		line := scanner.Text()
		if line == "" {
			fmt.Println("Empty line")
			continue
		}

		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			fmt.Println("Wrong game format")
			return
		}

		sets := strings.Split(parts[1], ";")
		for _, set := range sets {
			colorWithNumbers := strings.Split(set, ",")
			for _, turn := range colorWithNumbers {
				tmp := strings.Split(strings.TrimSpace(turn), " ")
				number, err := strconv.Atoi(tmp[0])
				if err != nil {
					return
				}
				switch tmp[1] {
				case "blue":
					if number > maxBlue {
						maxBlue = number
					}
				case "red":
					if number > maxRed {
						maxRed = number
					}

				case "green":
					if number > maxGreen {
						maxGreen = number
					}

				}

			}

		}
		sum += maxBlue * maxGreen * maxRed
	}
	fmt.Println(sum)
}

func parseFile(filePath string) (*bufio.Scanner, *os.File, error) {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Printf("Error when opening file with path %s", filePath)
		return nil, file, err
	}

	scanner := bufio.NewScanner((file))
	return scanner, file, nil

}

func main() {

	scanner, file, err := parseFile("day2input.txt")
	if err != nil {
		fmt.Println("Error when read file")
		return
	}
	defer file.Close()
	Part1(scanner)
	Part2(scanner)
}
