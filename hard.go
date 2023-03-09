package homework

import "strings"

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
	panic("Not implemented yet...")
}
