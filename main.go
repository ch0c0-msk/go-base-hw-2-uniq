package main

import (
	"log"

	"go-base-hw-2-uniq/cmdutils"
	"go-base-hw-2-uniq/fileutils"
	"go-base-hw-2-uniq/uniqparser"
)

func main() {
	args, opt, err := cmdutils.GetArgsAndOptions()
	if err != nil {
		log.Fatal(err)
	}

	inputFile, err := fileutils.GetInputFile(args)
	if err != nil {
		log.Fatal(err)
	}

	inputLines, err := fileutils.ReadLinesFromFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	outputLines := uniqparser.Uniq(inputLines, opt)

	outputFile, err := fileutils.GetOutputFile(args)
	if err != nil {
		log.Fatal(err)
	}

	err = fileutils.WriteLinesIntoFile(outputFile, outputLines)
	if err != nil {
		log.Fatal(err)
	}
}
