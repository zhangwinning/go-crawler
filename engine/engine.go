package engine

import (
	"go-crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var request []Request

	for _, r := range seeds {
		request = append(request, r)
	}

	for len(request) > 0 {
		r := request[0]
		request = request[1:]
		body, err := fetcher.Fetche(r.URL)
		if err != nil {
			log.Printf("Fetcher:error fecher url %s: %v", r.URL, err)
			continue
		}
		parseResult := r.ParseFunc(body)
	}
}
