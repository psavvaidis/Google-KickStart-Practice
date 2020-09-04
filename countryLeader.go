package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func countLetters(name string) int {
	lettersFound := make(map[string]bool)

	for _, char := range name {
		if string(char) == " " {
			continue
		}
		lettersFound[string(char)] = true
	}
	return len(lettersFound)
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func compareNames(prevName, newName *[]byte) {
	println("Leader: ", string(*prevName), ", New Name: ", string(*newName))

	curNameLetterNum := countLetters(string(*newName))
	currentLeaderLetterNum := countLetters(string(*prevName))

	if curNameLetterNum > currentLeaderLetterNum {
		*prevName = *newName
	}
	if curNameLetterNum == currentLeaderLetterNum {
		for i, char := range *newName {
			if string(string(*prevName)[i]) > string(char) {
				*prevName = *newName
				break
			}
		}
	}
}

func main() {

	// Read input File from arguments
	cmdArgs := os.Args
	inputFileName := cmdArgs[1]
	inputFile, err := os.Open(inputFileName)
	check(err)

	// Read the number of cases
	reader := bufio.NewReader(inputFile)
	cases, _, err := reader.ReadLine()
	check(err)
	casesNum, err := strconv.Atoi(string(cases))
	check(err)

	for i := 0; i < casesNum; i++ {
		currentLeader := []byte{}

		// Read Number of Inputs for this case
		inputs, _, err := reader.ReadLine()
		check(err)
		inputsNum, err := strconv.Atoi(string(inputs))
		check(err)

		for j := 0; j < inputsNum; j++ {
			// Read Input Name
			curName, _, err := reader.ReadLine()
			check(err)

			// Check conditions for name to be a leader
			compareNames(&currentLeader, &curName)
		}
		// Print the case result
		fmt.Println("Case #", i+1, ": ", string(currentLeader))

	}
}
