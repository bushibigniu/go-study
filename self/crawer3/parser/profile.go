package parser

import (
	_ "fmt"
	"go-study/self/crawer3/engine"
	"go-study/self/crawer3/model"
	"regexp"
	"strconv"
)
//使用这个正则，使用 ParseUserProfile，因为实际情况正则不一样
//<div data-v-bff6f798="" class="m-btn purple">未婚</div>
const regexpInfo  = `<div (.*?)class="m-btn purple"(.*?)>([^<]+)</div>`
const regexpInfo2  = `<div (.*?)class="m-btn pink"(.*?)>([^<]+)</div>`

//在这边初始化时为了提高性能，多个的时候，不需要每个都去实例化，而是在初始化就做掉
const age  = `<td><span class="label">年龄：</span>([\d]+)岁</td>`
var ageRe = regexp.MustCompile(age)

//废弃 因为正则不一样
func ParseUserProfile2(contents []byte) engine.ParseResult {
	//re := regexp.MustCompile(age)
	//	//
	//	//fmt.Print("begin")
	//	//
	//	//age := re.FindSubmatch(contents)
	//	//
	//	//fmt.Println(age)
	//	//fmt.Printf("age %s",age)
	//	//
	//	//return engine.ParseResult{}


	// 转成 int strconv.Atoi
	//var profile = model.Profile{}
	profile := model.Profile{}

	age, err := strconv.Atoi(extractString(contents, ageRe))

	if err != nil {
		//age
		profile.Age = string(age)
	}

	//marrages 是字符串， 直接使用
	//profile.Marriage = extractString(contents, marrages)

	return engine.ParseResult{}
}

func extractString(content []byte, re *regexp.Regexp) string{
	match := re.FindSubmatch(content)

	if len(match) >2 {
		return string(match[1])
	} else {
		return ""
	}



}


func ParseUserProfile(content []byte, name string) engine.ParseResult {
	re := regexp.MustCompile(regexpInfo)
	//re2 := regexp.MustCompile(regexpInfo2)

	// info 是类似这样的集合
//[<div class="m-btn purple" data-v-bff6f798>高中及以下</div>   data-v-bff6f798 高中及以下]]
	info := re.FindAllSubmatch(content, -1)

	var profile model.Profile
	profile.Name = name
	for k, v := range info{
		//把 if 条件写在这，和 封装成 dealValue 调用结果是一样的
		//用指针方式， 和 不用指针 用return 回来 结果是一样的
		//profile = dealValue(k, string(v[3]), profile)
		dealValue(k, string(v[3]), &profile)
	}

	result := engine.ParseResult{
		//{profile} 是具体的值
		Items:[]interface{}{profile},
	}

	//将profile 返回去
	return result
}

//可以优化
//func dealValue(k int, value string, profile model.Profile) (model.Profile){
func dealValue(k int, value string, profile *model.Profile){

	if k ==0 {
		profile.Marriage = value
	}

	if k ==1 {
		profile.Age = value
	}
	if k ==2 {
		profile.Xinzuo = value
	}
	if k ==3 {
		profile.Height = value
	}
	if k ==4 {
		profile.Weight = value
	}
	if k ==7 {
		profile.Occupation = value
	}
	if k ==6 {
		profile.Income = value
	}

	//return profile
}
