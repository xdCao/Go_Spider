package parser

import (
	"Go_Spider/crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9a-z]+)"[^>]+>([^<]+)</a>`

func ParseCity(content []byte) engine.ParserResult {

	re := regexp.MustCompile(cityRe)
	match := re.FindAllSubmatch(content, -1)

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

	return result
}
