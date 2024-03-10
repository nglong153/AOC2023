package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type WordNumber struct {
	code  string
	value string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var wordNumber = [...]WordNumber{
	{code: "one", value: "one1one"},
	{code: "two", value: "two2two"},
	{code: "three", value: "three3three"},
	{code: "four", value: "four4four"},
	{code: "five", value: "five5five"},
	{code: "six", value: "six6six"},
	{code: "seven", value: "seven7seven"},
	{code: "eight", value: "eight8eight"},
	{code: "nine", value: "nine9nine"},
}

func main() {
	FindFistAndLastNumber()
}

func FindFistAndLastNumber() {
	f, err := os.Open("Day1Input.txt")
	sum := 0
	check(err)

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		firstNum := 0
		lastNum := 0
		str := modifileTextInput(fileScanner.Text())
		for _, char := range str {
			if unicode.IsDigit(char) {
				firstNum1, _ := strconv.Atoi(string(char))
				firstNum = firstNum1
				break
			}

		}
		for i := len(str) - 1; i >= 0; i-- {
			char := rune(str[i])
			if unicode.IsDigit(char) {
				lastNumTemp, _ := strconv.Atoi(string(char))
				lastNum = lastNumTemp
				break
			}
		}
		sum += firstNum*10 + lastNum
		fmt.Println(str, fileScanner.Text(), firstNum, lastNum)
	}

	check(err)

	fmt.Println(sum)
}

func modifileTextInput(text string) string {

	retString := text

	for _, value := range wordNumber {
		retString = strings.Replace(retString, value.code, value.value, -1)
	}

	return retString
}
