package main

import (
	"go-base-hw-2-uniq/cmdUtils"
	"go-base-hw-2-uniq/fileUtils"
	"go-base-hw-2-uniq/uniqParser"
	"log"
)

func main() {
	args, opt, err := cmdUtils.GetArgsAndOptions()
	if err != nil {
		log.Fatal(err)
	}

	inputFile, err := fileUtils.GetInputFile(args)
	if err != nil {
		log.Fatal(err)
	}

	inputLines, err := fileUtils.ReadLinesFromFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	outputLines := uniqParser.Uniq(inputLines, opt)

	outputFile, err := fileUtils.GetOutputFile(args)
	if err != nil {
		log.Fatal(err)
	}

	err = fileUtils.WriteLinesIntoFile(outputFile, outputLines)
	if err != nil {
		log.Fatal(err)
	}
}
