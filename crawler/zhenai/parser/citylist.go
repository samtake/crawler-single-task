package parser

import (
	"crawler-single-task/engine"
	"regexp"
	// "log"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z]+)"[^>]*>([^<]+)</a>`

// 城市列表 .
func ParserCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	// log.Printf("matches = %s", matches)
	/*
	[
		[<a href="http://www.zhenai.com/zhenghun/aba" data-v-1573aa7c>阿坝</a> http://www.zhenai.com/zhenghun/aba 阿坝] 
		[<a href="http://www.zhenai.com/zhenghun/akesu" data-v-1573aa7c>阿克苏</a> http://www.zhenai.com/zhenghun/akesu 阿克苏]
	]
	*/

	result := engine.ParseResult{}

	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			// ParserFunc: engine.NilParser,
			ParserFunc: ParserCity,
		})

		// log.Printf("m0 = %s", m[0])
		// log.Printf("m1 = %s", m[1])
		// log.Printf("m2 = %s", m[2])
		/*
		m0 = <a href="http://www.zhenai.com/zhenghun/aba" data-v-1573aa7c>阿坝</a>
		m1 = http://www.zhenai.com/zhenghun/aba
		m2 = 阿坝
		*/
	}
	return result
}
