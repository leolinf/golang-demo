package parser

import (
	"golang-demo/crawler/engine"
	"golang-demo/crawler/model"
	"regexp"
)

var ageRe = regexp.MustCompile(`<div class="des f-cl" data-v-3c42fade>([^<]+)</div>`)

func ParseProfile(contents []byte) engine.ParserResult {
	profile := model.Profile{}
	gender := extractString(contents, ageRe)
	//age, err := strconv.Atoi(
	//	extractString(contents, ageRe))
	//if err != nil {
	//	profile.Age = age
	//}
	profile.Gender = gender
	result := engine.ParserResult{
		Item: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
