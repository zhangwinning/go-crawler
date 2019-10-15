package engine

import (
	"go-crawler/fetcher"
	"log"
)

type SimpleEngine struct { }



func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching %s", r.URL)
		// 并发版爬虫就是把 worker 搞成并发。
		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got Item %v", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.URL)
	body, err := fetcher.Fetcher(r.URL)
	if err != nil {
		log.Printf("Fetcher:error fecher url %s: %v", r.URL, err)
		return ParseResult{}, err
	}
	return r.ParseFunc(body), nil
}
