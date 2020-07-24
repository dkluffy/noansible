package extra

import (
	"regexp"
	"strconv"
	"strings"
)

// func SliceRemove2(s *[]interface{}, index int) {
// 	*s = append((*s)[:index], (*s)[index+1:]...)
// }

var reDigitalOnly = regexp.MustCompile("^[0-9.]*$")

//格式用`;`分割： left; [ >,<,=,!=,>=,<= ] ;right
//出错返回FALSE
func IfTester(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	if len(s) == 0 {
		return false
	}
	ruleStrings := strings.Split(s, ";")

	switch len(ruleStrings) {
	case 1:
		return true
	case 2: //长度不合法
		return false
	case 3:
		if reDigitalOnly.MatchString(ruleStrings[0]) && reDigitalOnly.MatchString(ruleStrings[2]) {
			return cmpNum(ruleStrings[0], ruleStrings[1], ruleStrings[2])
		} else {
			return cmpStr(ruleStrings[0], ruleStrings[1], ruleStrings[2])
		}

	}

	return false
}

//compare String
func cmpStr(s1 string, cmp string, s2 string) bool {
	if cmp == "=" {
		return s1 == s2
	}
	return false
}

//compare numbers - int only
func cmpNum(s1 string, cmp string, s2 string) bool {
	v1, err := strconv.Atoi(s1)
	if err != nil {
		return false
	}
	v2, err := strconv.Atoi(s2)
	if err != nil {
		return false
	}
	switch cmp {
	case ">":
		return v1 > v2
	case "<":
		return v1 < v2
	case "=":
		return v1 == v2
	case "!=":
		return v1 != v2
	case ">=":
		return v1 >= v2
	case "<=":
		return v1 <= v2

	}
	return false

}
