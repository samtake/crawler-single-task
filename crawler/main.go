package main

import (
	"crawler-single-task/engine"
	"crawler-single-task/zhenai/parser"
)

func main() {

	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})

	// resp, err := http.Get("http://www.zhenai.com/zhenghun")
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()

	// if resp.StatusCode != http.StatusOK {
	// 	fmt.Println("Error:status code", resp.StatusCode)
	// 	return
	// }
	// //转码
	// e := determinEncoding(resp.Body)
	// utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	// all, err := ioutil.ReadAll(utf8Reader)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s\n", all)

	// printCityList(all)
}

// func printCityList(contents []byte) {
// 	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z]+)"[^>]*>([^<]+)</a>`)
// 	matches := re.FindAllSubmatch(contents, -1)
// 	//[][][]byte
// 	for _, m := range matches {
// 		// for _, subMatch := range m {
// 		// 	fmt.Printf("%s\n", subMatch)
// 		// }
// 		fmt.Printf("City: %s, URL:%s\n", m[2], m[1])
// 	}
// 	fmt.Printf("matches found:%d\n", len(matches))
// }
