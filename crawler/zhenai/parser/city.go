package parser

import (
	"Go_Spider/crawler/engine"
	"regexp"
)

var profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9a-z]+)"[^>]+>([^<]+)</a>`)
var cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)

func ParseCity(content []byte) engine.ParserResult {

	match := profileRe.FindAllSubmatch(content, -1)

	result := engine.ParserResult{}

	for _, m := range match {
		name := string(m[2])
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParserResult {
				return ParseProfile(c, name)
			},
		})
	}

	match = cityUrlRe.FindAllSubmatch(content, -1)
	for _, m := range match {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
