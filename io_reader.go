package hard

import (
	"io"
	"strings"
)

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
