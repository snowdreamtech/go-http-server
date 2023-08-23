package i18n

import (
	"fmt"
	"testing"

	"golang.org/x/text/language"
)

var acceptLanguages = []string{
	"zh-Hans",
	"zh-Hant",
	"zh-CN,zh;q=0.8,en;q=0.2",
	"zh-CN",
	"zh-TW",
	"zh-HK",
	"zh-SG",
	"zh-MO",
	"en",
	"en-US",
	"fr-CH, fr;q=0.9, en;q=0.8, de;q=0.7, *;q=0.5",
	"zh-Hans",    // 简体中文
	"zh-Hans-CN", // 大陆地区使用的简体中文
	"zh-Hans-HK", // 香港地区使用的简体中文
	"zh-Hans-MO", // 澳门使用的简体中文
	"zh-Hans-SG", // 新加坡使用的简体中文
	"zh-Hans-TW", // 台湾使用的简体中文
	"zh-Hant",    // 繁体中文
	"zh-Hant-CN", // 大陆地区使用的繁体中文
	"zh-Hant-HK", // 香港地区使用的繁体中文
	"zh-Hant-MO", // 澳门使用的繁体中文
	"zh-Hant-SG", // 新加坡使用的繁体中文
	"zh-Hant-TW", // 台湾使用的繁体中文
}

func TestParseAcceptLanguage(t *testing.T) {
	for _, acceptLanguage := range acceptLanguages {
		testname := "Accept-Language:" + acceptLanguage

		t.Run(testname, func(t *testing.T) {
			tags, _, _ := language.ParseAcceptLanguage(acceptLanguage)

			for _, tag := range tags {
				base, _ := tag.Base()
				script, _ := tag.Script()
				region, _ := tag.Region()

				fmt.Println("tag:")
				fmt.Println(tag)

				fmt.Println("base:")
				fmt.Println(base)
				fmt.Println("script:")
				fmt.Println(script)
				fmt.Println("region:")
				fmt.Println(region)
				fmt.Println("")
				fmt.Println("")
			}
		})
	}
}
