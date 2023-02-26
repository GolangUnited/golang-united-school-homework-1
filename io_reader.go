package homework

import (
	"fmt"
	"io"
	"strings"
)

/*
SeekTillHalfOfString -  contains a code snippet in Go that defines a function called
"SeekTillHalfOfString". The function takes a string reader as input,
seeks to the middle of the string, reads
half of the remaining string, and returns it as a string.
*/
//func strReader(str string) {
//
//}

func SeekTillHalfOfString(strReader *strings.Reader) string {
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

	var b []string
	reader := io.Reader(strReader)
	buf := make([]byte, n)
	for {
		n, err := reader.Read(buf)
		if err == io.EOP {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		if n > 0 {
			b = append(b, string(buf[:n]))
		}
	}
	return b
}
