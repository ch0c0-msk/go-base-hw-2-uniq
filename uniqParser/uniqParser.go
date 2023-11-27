package uniqParser

import (
	"go-base-hw-2-uniq/utils"
	"strconv"
	"strings"
)

type ChangeData struct {
	commonString  string
	compareString string
	count         int
}

func createCompareString(commonString string, opt *utils.Options) string {
	compareString := commonString
	if opt.IFlag {
		compareString = strings.ToLower(compareString)
	}
	if opt.FFlag > 0 {
		for i := 0; i < opt.FFlag; i++ {
			index := strings.IndexByte(compareString, ' ')
			if index == -1 {
				break
			}
			compareString = compareString[index+1:]
		}
	}

	if opt.SFlag > 0 {
		left := opt.SFlag
		right := len(compareString)
		if left > right {
			left = right
		}
		compareString = compareString[left:right]
	}

	return compareString
}

func createResultWithFormPattern(changeDataList []ChangeData, opt *utils.Options) []string {
	result := []string{}
	if opt.CFlag {
		for _, val := range changeDataList {
			if val.count > 0 {
				countAndString := strings.Join([]string{strconv.Itoa(val.count), val.commonString}, " ")
				result = append(result, countAndString)
			}
		}
	} else if opt.DFlag {
		for _, val := range changeDataList {
			if val.count > 1 {
				result = append(result, val.commonString)
			}
		}
	} else if opt.UFlag {
		for _, val := range changeDataList {
			if val.count == 1 {
				result = append(result, val.commonString)
			}
		}
	} else {
		for _, val := range changeDataList {
			if val.count > 0 {
				result = append(result, val.commonString)
			}
		}
	}

	return result
}

func Uniq(inputLines []string, options *utils.Options) []string {

	var changeDataList []ChangeData
	compareLinesMap := map[string]int{}

	for _, line := range inputLines {
		compareString := createCompareString(line, options)
		changeData := ChangeData{line, compareString, 0}
		changeDataList = append(changeDataList, changeData)
		if _, inMap := compareLinesMap[compareString]; !inMap {
			compareLinesMap[compareString] = 0
		}
		compareLinesMap[compareString]++
	}

	for i, val := range changeDataList {
		if count, inMap := compareLinesMap[val.compareString]; inMap {
			changeDataList[i].count = count
			delete(compareLinesMap, val.compareString)
		}
	}

	outputStrings := createResultWithFormPattern(changeDataList, options)
	return outputStrings

}
