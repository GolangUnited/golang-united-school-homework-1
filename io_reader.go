package hard

import (
	"fmt"
	"strings"
)

/*
You own a Goal Parser that can interpret a string command.
The command consists of an alphabet of "G", "()" and/or "(al)" in some order.
The Goal Parser will interpret
"G" as the string "G", "()" as the string "o", and "(al)" as the string "al".
The interpreted strings are then concatenated in the original order.
For example:
input: "G()(al)
output: Goal
input: G()()()()(al)
output: Gooooal
input: (al)G(al)()()G
output: alGalooG
*/
func GoalParsers(strReader *strings.Reader) string {

	var (
		readerBuf = make([]byte, 1024)
	)

	_, err := strReader.Read(readerBuf)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(readerBuf))
	parsedString := parse(string(readerBuf))
	return parsedString
}

func parse(word string) string {

	var (
		output string
	)

	if word == "" {
		return ""
	}

	for i := 0; i < len(word); i++ {

		switch b := word[i]; string(b) {
		case "G":
			output = "G"
		case "(":
			output += "("
		case ")":
			output += ")"
		case "a":
			output += "a"
		case "l":
			output += "l"

		}

		switch output {
		case "G":
			return "G" + parse(word[i+1:])
		case "()":
			return "o" + parse(word[i+1:])
		case "(al)":
			return "al" + parse(word[i+1:])
		}
	}

	return ""
}
