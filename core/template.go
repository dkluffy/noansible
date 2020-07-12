package core

import (
	"errors"
	"regexp"
	"strings"
)

var reFindPat *regexp.Regexp = regexp.MustCompile(`\{\{.*?\}\}`)
var reRep *regexp.Regexp = regexp.MustCompile(`[\{\} ]`)

func Render(oristr string) (string, error) {
	//如果某个全局变量PlaybookVars[nv] 不存在，则返回空
	var err error
	resultstr := oristr

	varlist := reFindPat.FindAllString(oristr, -1)
	for _, v := range varlist {
		nv := reRep.ReplaceAllString(v, "")
		pv, ok := PlaybookVars[nv]
		if !ok || reFindPat.MatchString(pv) {
			err = errors.New("RenderFailed: " + oristr)
			return "", err
		}
		resultstr = strings.ReplaceAll(resultstr, v, pv)
	}
	return resultstr, err

}
