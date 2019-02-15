package parser

import (
	"imooc.com/u2pppw/crawler/engine"
	"imooc.com/u2pppw/crawler/model"
	"regexp"
)

var profileRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([^<]+)</div>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	profileStringList := model.ProfileStringList{}
	profileStringList.Info = append(profileStringList.Info, name)
	for _, m := range matches {
		profileStringList.Info = append(profileStringList.Info, string(m[1]))
	}

	result.Items = []interface{}{profileStringList}
	return result
}
