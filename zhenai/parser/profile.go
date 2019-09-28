package parser

import (
	//"fmt"
	"go-crawler/engine"
	"go-crawler/model"
	"regexp"
	"strconv"
)

var marriageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>(离异|未婚|已婚)</div>`)
var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>(\d+)岁</div>`)
var xinzuoRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<]+座)[^<]+</div>`)
var addressRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>工作地:([^<]+)</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>月收入:([^<]+)</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d]+)cm</div>`)
var weightRe = regexp.MustCompile(` <div class="m-btn purple" data-v-8b1eac0c>([\d]+)kg</div>`)
var educationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>(高中及以下|高中|大学本科)</div>`)

func ParseProfile(contents []byte, name string, gender string) engine.ParseResult {

	profile := model.Profile{}

	profile.Name = name
	profile.Gender = gender
	age, _ := strconv.Atoi(extractString(contents, ageRe))
	profile.Age = age
	height, _ := strconv.Atoi(extractString(contents, heightRe))
	profile.Height = height
	weight, _ := strconv.Atoi(extractString(contents, weightRe))
	profile.Weight = weight
	profile.Marriage = extractString(contents, marriageRe)
	profile.Xinzuo = extractString(contents, xinzuoRe)
	profile.Address = extractString(contents, addressRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Education = extractString(contents, educationRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
