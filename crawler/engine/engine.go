package engine

import (
	"crawler-single-task/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	// log.Printf("requests=%v; len(requests)=%v ", requests, len(requests))
	// return
	// /*requests=[{http://www.zhenai.com/zhenghun 0x1272b60}]; len(requests)=1 */

	/*
		requests[0]
		{http://www.zhenai.com/zhenghun 0x1272b60}
		Url	   							ParserFunc
	*/

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetch:error "+"fetching url %s: %v", r.Url, err)
			continue
		}

		// log.Printf("body=%s", body)

		parseResult := r.ParserFunc(body)
		requests = append(requests, parseResult.Requests...)

		// for _, item := range parseResult.Items {
		// 	// log.Printf("Got item %v", item)
		// 	log.Printf("Got item %s", item)
		// }

	}
}
