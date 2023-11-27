package utils

import (
	"errors"
	"flag"
)

type Options struct {
	CFlag bool
	DFlag bool
	UFlag bool
	FFlag int
	SFlag int
	IFlag bool
}

func GetArgsAndOptions() ([]string, *Options, error) {
	opt := new(Options)
	flag.BoolVar(&opt.CFlag, "c", false, "count the number of occurrences")
	flag.BoolVar(&opt.DFlag, "d", false, "output only non-unique strings")
	flag.BoolVar(&opt.UFlag, "u", false, "output only unique strings")
	flag.IntVar(&opt.FFlag, "f", 0, "ignore the first num_fields of fields in a row")
	flag.IntVar(&opt.SFlag, "s", 0, "ignore the first num_chars characters in the string")
	flag.BoolVar(&opt.IFlag, "i", false, "ignore case of letters")
	flag.Parse()

	if !isOptionsCorrect(opt) {
		return []string{}, nil, errors.New("invalid flag values")
	}

	args := flag.Args()

	if len(args) > 2 {
		return []string{}, nil, errors.New("invalid number of arguments")
	} else {
		return args, opt, nil
	}
}

func isOptionsCorrect(opt *Options) bool {
	result := true
	if (opt.CFlag && opt.DFlag) || (opt.CFlag && opt.UFlag) || (opt.DFlag && opt.UFlag) {
		result = false
	}
	if (opt.FFlag < 0) || (opt.SFlag < 0) {
		result = false
	}
	return result
}
