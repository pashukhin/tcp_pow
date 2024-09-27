package quotes

import (
	"bufio"
	"os"

	"math/rand"
)

type quotes struct {
	quotes []string
}

func New(path string) (Quotes, error) {
	readFile, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	return &quotes{
		quotes: fileLines,
	}, nil
}

func (q *quotes) Quote() string {
	return q.quotes[rand.Intn(len(q.quotes))]
}
