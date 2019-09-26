package engine

import (
	"go-crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching %s", r.URL)
		body, err := fetcher.Fetche(r.URL)
		if err != nil {
			log.Printf("Fetcher:error fecher url %s: %v", r.URL, err)
			continue
		}
		parseResult := r.ParseFunc(body)
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got Item %v", item)
		}
	}
}
