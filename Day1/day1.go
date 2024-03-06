package main 

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("Day1Input.txt")
	check(err)

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		firstNum:=0
		lastNum:=0
		str := fileScanner.Text()
		for _,char := range str{
			if unicode.IsDigit(char){
firstNum1,_ = strconv.Atoi(string(char))
fistNum = fisrtNum1
			break
			}
		for i:=len(str)-1;i >= 0; i--{
			char = rune(str[i])
	if unicode.IsDigit(char){
			lastNumTemp,_ = strconv.Atoi(string(char))
lastNum = lastNumTemp
break;
		}

		fmt.Println(firstNum,lastNum)
	}}}

	check(err)

	fmt.Println("Hello world")
}
