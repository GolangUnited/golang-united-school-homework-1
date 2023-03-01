package homework

import (
	"io"
	"strings"
)

/*
SeekTillHalfOfString -  contains a code snippet in Go that defines a function called
"SeekTillHalfOfString". The function takes a string reader as input,
seeks to the middle of the string, reads
half of the remaining string, and returns it as a string.
*/

func SeekTillHalfOfString(strReader *strings.Reader) string {
	bytesArray := make([]byte, strReader.Len())
	stringLength, _ := strReader.Read(bytesArray)
	return string(bytesArray[stringLength/2:])
}

/*
ReaderSplit - contains a code snippet written in Go that
defines a function called ReaderSplit.
The function takes a strings.Reader and an integer n as input,
and splits the contents of the reader into chunks of size n.
The function returns a slice of strings containing the chunks
*/

func ReaderSplit(strReader *strings.Reader, n int) []string {
	var chunks []string
	var bytes = make([]byte, n)
	for {
		n, err := strReader.Read(bytes)
		if err == io.EOF {
			break
		}
		if n > 0 {
			chunks = append(chunks, string(bytes[:n]))
		}
	}
	return chunks
}
