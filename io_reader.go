package homework

import (
	"fmt"
	"strings"
)

/*
SeekTillHalfOfString -  contains a code snippet in Go that defines a function called
"SeekTillHalfOfString". The function takes a string reader as input,
seeks to the middle of the string, reads
half of the remaining string, and returns it as a string.
*/

func SeekTillHalfOfString(strReader *strings.Reader) string {

	length := strReader.Len()
	output := make([]byte, length-length/2)
	_, err := strReader.Read(make([]byte, length/2))
	_, err = strReader.Read(output)
	if err != nil {
		fmt.Println("EOF")
	}
	// fmt.Println("size: ", n)
	// fmt.Println("output: ", string(output))

	return ""
}

/*
ReaderSplit - contains a code snippet written in Go that
defines a function called ReaderSplit.
The function takes a strings.Reader and an integer n as input,
and splits the contents of the reader into chunks of size n.
The function returns a slice of strings containing the chunks
*/

func ReaderSplit(strReader *strings.Reader, n int) []string {

	output := []string{}

	number_of_chunks := strReader.Len()/n + 2

	for i := 0; i < number_of_chunks; i++ {
		chunk := make([]byte, n)
		_, err := strReader.Read(chunk)
		fmt.Println("chunk: ", string(chunk))
		output = append(output, string(chunk))
		if err != nil {
			fmt.Println("EOF")
			break
		}

	}

	return output
}
