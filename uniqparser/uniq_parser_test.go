package uniqparser

import (
	"go-base-hw-2-uniq/cmdutils"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createEmptyOptions() *cmdutils.Options {
	opt := new(cmdutils.Options)
	return opt
}

func createInputLinesSlice() []string {
	lines := []string{}
	lines = append(lines, "i love", "i love", "we love", "we love", "we ll love", "you love")
	return lines
}

func TestUniqParserWithoutFlags(t *testing.T) {
	expected := []string{"i love", "we love", "we ll love", "you love"}
	opt := createEmptyOptions()
	actual := Uniq(createInputLinesSlice(), opt)
	assert.Equal(t, expected, actual)
}

func TestUniqParserWithCFlag(t *testing.T) {
	expected := []string{"2 i love", "2 we love", "1 we ll love", "1 you love"}
	opt := createEmptyOptions()
	opt.CFlag = true
	actual := Uniq(createInputLinesSlice(), opt)
	assert.Equal(t, expected, actual)
}

func TestUniqParserWithDFlag(t *testing.T) {
	expected := []string{"i love", "we love"}
	opt := createEmptyOptions()
	opt.DFlag = true
	actual := Uniq(createInputLinesSlice(), opt)
	assert.Equal(t, expected, actual)
}

func TestUniqParserWithUFlag(t *testing.T) {
	expected := []string{"we ll love", "you love"}
	opt := createEmptyOptions()
	opt.UFlag = true
	actual := Uniq(createInputLinesSlice(), opt)
	assert.Equal(t, expected, actual)
}

func TestUniqParserWithIFlag(t *testing.T) {
	expected := []string{"i love", "we love"}
	opt := createEmptyOptions()
	opt.DFlag = true
	opt.IFlag = true
	lines := createInputLinesSlice()
	lines[1] = strings.ToUpper(lines[1])
	actual := Uniq(lines, opt)
	assert.Equal(t, expected, actual)
}

func TestUniqParserWithFFlagAndSFlag(t *testing.T) {
	expected := []string{"i llove"}
	opt := createEmptyOptions()
	opt.DFlag = true
	opt.FFlag = 1
	opt.SFlag = 1
	lines := []string{"i llove", "we wlove"}
	actual := Uniq(lines, opt)
	assert.Equal(t, expected, actual)
}
