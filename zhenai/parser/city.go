package parser

import (
	"go-crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
const genderRe = `<td width="180"><span class="grayL">性别：</span>([^<]+)</td>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	reGender := regexp.MustCompile(genderRe)
	matchesGender := reGender.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for i, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "User: "+ name)
		result.Requests = append(result.Requests, engine.Request{
			URL: string(m[1]),
			ParseFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name, string(matchesGender[i][1]))
			},
		})
	}
	return result
}
