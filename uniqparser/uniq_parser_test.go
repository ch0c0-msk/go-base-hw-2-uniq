package uniqparser

import (
	"go-base-hw-2-uniq/cmdutils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlags(t *testing.T) {
	testCases := []struct {
		name   string
		input  []string
		opt    *cmdutils.Options
		output []string
	}{
		{
			"NoneFlag",
			[]string{"i love", "i love", "we love", "we love", "we ll love", "you love"},
			&cmdutils.Options{},
			[]string{"i love", "we love", "we ll love", "you love"},
		},
		{
			"CFlag",
			[]string{"i love", "i love", "we love", "we love", "we ll love", "you love"},
			&cmdutils.Options{CFlag: true},
			[]string{"2 i love", "2 we love", "1 we ll love", "1 you love"},
		},
		{
			"DFlag",
			[]string{"i love", "i love", "we love", "we love", "we ll love", "you love"},
			&cmdutils.Options{DFlag: true},
			[]string{"i love", "we love"},
		},
		{
			"UFlag",
			[]string{"i love", "i love", "we love", "we love", "we ll love", "you love"},
			&cmdutils.Options{UFlag: true},
			[]string{"we ll love", "you love"},
		},
		{
			"IFlag",
			[]string{"i love", "I LOVE", "we love", "we love", "we ll love", "you love"},
			&cmdutils.Options{IFlag: true, DFlag: true},
			[]string{"i love", "we love"},
		},
		{
			"FFlagAndSFlag",
			[]string{"i llove", "we wlove"},
			&cmdutils.Options{FFlag: 1, SFlag: 1, DFlag: true},
			[]string{"i llove"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Uniq(tc.input, tc.opt)
			assert.Equal(t, tc.output, actual)
		})
	}
}
