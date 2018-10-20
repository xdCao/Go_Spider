package parser

import (
	"Go_Spider/crawler/engine"
	"Go_Spider/crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var marrigeRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)元</td>`)
var eduRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var workPlaceRe = regexp.MustCompile(`<td><span class="label">工作地：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var xinzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span>([^<]+)</td>`)
var jiguanRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([\d]+)KG</span></td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)

func ParseProfile(contents []byte, url string, name string) engine.ParserResult {

	profile := model.Profile{}
	profile.Url = url
	profile.Id = url
	profile.Name = name
	profile.Marriage = extractString(contents, marrigeRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Education = extractString(contents, eduRe)
	profile.WorkPlace = extractString(contents, workPlaceRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Xinzuo = extractString(contents, xinzuoRe)
	profile.Jiguan = extractString(contents, jiguanRe)
	profile.Gender = extractString(contents, genderRe)
	profile.House = extractString(contents, houseRe)
	profile.Car = extractString(contents, carRe)
	//profile.Name =

	if age, err := strconv.Atoi(extractString(contents, ageRe)); err == nil {
		profile.Age = age
	}

	if height, err := strconv.Atoi(extractString(contents, heightRe)); err == nil {
		profile.Height = height
	}

	if weight, err := strconv.Atoi(extractString(contents, weightRe)); err == nil {
		profile.Weight = weight
	}

	result := engine.ParserResult{
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
