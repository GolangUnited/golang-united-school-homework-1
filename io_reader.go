package hard

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
func SeekTillHalfOfString(strReader *strings.Reader) string {
	var (
		err error
	)

	defer func() {
		if b := recover(); b != nil {
			fmt.Println("Panic: ", b)
		}
	}()

	_, err = strReader.Seek(strReader.Size()/2, 0)
	if err != nil {
		panic(err)
	}

	rSize := strReader.Size() - strReader.Size()/2
	rBuf := make([]byte, rSize)

	_, err = strReader.Read(rBuf)
	if err != nil {
		panic(err)
	}

	return string(rBuf)
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
				break
			}
		}
		chunks = append(chunks, string(buf[:n]))
	}
	return chunks
}
