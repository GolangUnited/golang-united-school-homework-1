## IO Reader and String Task

You own a Goal Parser that can interpret a string command.
The command consists of an alphabet of "G", "()" and/or "(al)" in some order.
The Goal Parser will interpret
"G" as the string "G", "()" as the string "o", and "(al)" as the string "al".
The interpreted strings are then concatenated in the original order.

```go
func GoalParsers(strReader *strings.Reader) string {
	panic("Not implemented yet...")
}
```

- Task is to implement the `GoalParsers` method based on the above requirements

Example Inputs:
```
input: G()(al)
output: Goal
```
```
input: G()()()()(al)
output: Gooooal
```
```
input: (al)G(al)()()G
output: alGalooG
```


## IO Seeker Task

SeekTillHalfOfString -  contains a code snippet in Go that defines a function called
"SeekTillHalfOfString". The function takes a string reader as input,
seeks to the middle of the string, reads
half of the remaining string, and returns it as a string.

```go
func SeekTillHalfOfString(strReader *strings.Reader, n int) []string {
	panic("Not implemented yet...")
}
```

- Task is to implement the `SeekTillHalfOfString` method based on the above requirements

Example Inputs:
```
input: HELLO-WORLD
output: -WORLD
```
```
input: HELLOWORLD
output: WORLD
```
```
input: GOLANG-UNITED
output: -UNITED
```