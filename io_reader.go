package hard

import (
	"fmt"
	"io"
	"strings"
)

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

func SeekTillHalfOfString(strReader *strings.Reader) string {
    strLength := strReader.Len()

    _, err := strReader.Seek(int64(strLength/2), io.SeekStart)
    if err != nil {
        panic(err)
    }

    buf := make([]byte, (strLength-strLength/2))
    _, err = strReader.Read(buf)
    if err != nil {
        if err == io.EOF {
            return string(buf)
        }
        panic(err)
    }

    return string(buf)
}
