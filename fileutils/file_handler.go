package fileutils

import (
	"bufio"
	"flag"
	"io"
	"os"
)

func GetInputFile(args []string) (*os.File, error) {
	if len(args) < 1 {
		return os.Stdin, nil
	}
	inputFile, err := os.OpenFile(flag.Arg(0), os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	return inputFile, nil

}

func GetOutputFile(args []string) (*os.File, error) {
	if len(args) < 2 {
		return os.Stdout, nil
	}
	outputFile, err := os.OpenFile(flag.Arg(1), os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	return outputFile, nil

}

func ReadLinesFromFile(ioReader io.Reader) ([]string, error) {
	reader := bufio.NewReader(ioReader)
	var lines []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} 
			return []string{}, err
			
		} else {
			line = line[:len(line)-1]
			lines = append(lines, line)
		}
	}
	return lines, nil
}

func WriteLinesIntoFile(ioWriter io.Writer, lines []string) error {
	writer := bufio.NewWriter(ioWriter)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	err := writer.Flush()
	return err
}
