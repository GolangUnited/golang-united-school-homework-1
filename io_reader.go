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
	strLength := strReader.Len()

	_, err := strReader.Seek(int64(strLength/2), io.SeekStart)
	if err != nil {
		panic(err)
	}

	buf := make([]byte, (strLength - strLength/2))
	_, err = strReader.Read(buf)
	if err != nil {
		if err == io.EOF {
			return string(buf)
		}
		panic(err)
	}

	return string(buf)
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
	buf := make([]byte, n)
	for {
		_, err := strReader.Read(buf)
		if err == io.EOF {
			if len(buf) > 0 {
				chunks = append(chunks, string(buf[:]))
			}
			break
		}
		chunks = append(chunks, string(buf[:]))
	}
	return chunks
}


