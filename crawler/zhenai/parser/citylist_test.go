package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	// contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	// fmt.Printf("%s\n", contents)
	// ParserCityList(contents)
	result := ParserCityList(contents)
	const resultSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/hainan",
		"http://www.zhenai.com/zhenghun/yangjiang",
		"http://www.zhenai.com/zhenghun/guangzhou",
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d "+"requests; but had %d", resultSize, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but "+"was %s", i, url, result.Requests[i].Url)
		}
	}
}
