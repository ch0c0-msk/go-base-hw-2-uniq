package main

import (
	"go-base-hw-2-uniq/uniqparser"
	"go-base-hw-2-uniq/utils"
	"log"
)

func main() {
	args, opt, err := utils.GetArgsAndOptions()
	if err != nil {
		log.Fatal(err)
	}

	inputFile, err := utils.GetInputFile(args)
	if err != nil {
		log.Fatal(err)
	}

	inputLines, err := utils.ReadLinesFromFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	outputLines := uniqparser.Uniq(inputLines, opt)

	outputFile, err := utils.GetOutputFile(args)
	if err != nil {
		log.Fatal(err)
	}

	err = utils.WriteLinesIntoFile(outputFile, outputLines)
	if err != nil {
		log.Fatal(err)
	}
}
