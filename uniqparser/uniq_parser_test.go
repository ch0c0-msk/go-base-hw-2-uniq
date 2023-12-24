package uniqparser

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"go-base-hw-2-uniq/cmdutils"
)

var testLines = []string{"i love", "i love", "we love", "we love", "we ll love", "you love"}

func TestUniqParserWithoutFlags(t *testing.T) {
	expected := []string{"i love", "we love", "we ll love", "you love"}
	opt := &cmdutils.Options{}
	actual := Uniq(testLines, opt)
	assert.Equal(t, expected, actual)
}

func TestUniqParserWithCFlag(t *testing.T) {
	expected := []string{"2 i love", "2 we love", "1 we ll love", "1 you love"}
	opt := &cmdutils.Options{CFlag: true}
	actual := Uniq(testLines, opt)
	assert.Equal(t, expected, actual)
}

func TestUniqParserWithDFlag(t *testing.T) {
	expected := []string{"i love", "we love"}
	opt := &cmdutils.Options{DFlag: true}
	actual := Uniq(testLines, opt)
	assert.Equal(t, expected, actual)
}

func TestUniqParserWithUFlag(t *testing.T) {
	expected := []string{"we ll love", "you love"}
	opt := &cmdutils.Options{UFlag: true}
	actual := Uniq(testLines, opt)
	assert.Equal(t, expected, actual)
}

func TestUniqParserWithIFlag(t *testing.T) {
	expected := []string{"i love", "we love"}
	opt := &cmdutils.Options{DFlag: true, IFlag: true}
	lines := testLines
	lines[1] = strings.ToUpper(lines[1])
	actual := Uniq(lines, opt)
	assert.Equal(t, expected, actual)
}

func TestUniqParserWithFFlagAndSFlag(t *testing.T) {
	expected := []string{"i llove"}
	opt := &cmdutils.Options{DFlag: true, FFlag: 1, SFlag: 1}
	lines := []string{"i llove", "we wlove"}
	actual := Uniq(lines, opt)
	assert.Equal(t, expected, actual)
}
