package parser

import (
	"crawler-single-task/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParserCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		name:=string(m[2])
		result.Items = append(result.Items, "User "+string(m[2]))

		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			// ParserFunc: engine.NilParser,
			ParserFunc: func(contents []byte) engine.ParseResult {
				// return ParseProfile(contents, string(m[2]))
				return ParseProfile(contents, name)
			},
		})
	}
	return result
}
