package utils

import (
	"regexp"
	"strings"
)

func GetDetailColumnSQL(data string) (columnName, value string) {
	var res []string
	re := regexp.MustCompile(`\((.*?)\)`)
	submatchall := re.FindAllString(data, -1)
	for _, element := range submatchall {
		element = strings.Trim(element, "(")
		element = strings.Trim(element, ")")
		res = append(res, element)
	}
	return res[0], res[1]
}
