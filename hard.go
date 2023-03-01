package hard

import (
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
	stack := make([]string, 0)
	stack_element := ""
	length := strReader.Len()
	for i := 0; i < length; i++ {
		b, _ := strReader.ReadByte()
		char := string(b)

		if char == "(" {
			stack_element += char
		} else if char == ")" {
			stack_element += char
			stack = append(stack, stack_element)
			stack_element = ""
		} else if char == "G" {
			stack_element += char
			stack = append(stack, stack_element)
			stack_element = ""
		} else {
			stack_element += char
		}

	}
	output := ""
	for _, value := range stack {

		if value == "()" {
			output += "o"
		} else if value == "G" {
			output += "G"
		} else {
			output += value[1 : len(value)-1]
		}
	}
	return output

}
