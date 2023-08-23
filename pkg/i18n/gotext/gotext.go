package gotext

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/leonelquinteros/gotext"
	"golang.org/x/text/language"
	"snowdream.tech/http-server/pkg/i18n"
)

// I18N I18N
type I18N struct {
	// Supported Languages
	tags map[string]language.Tag

	matcher language.Matcher

	translators map[string]*gotext.Locale

	fallback    *gotext.Locale
	fallbacktag language.Tag
}

// NewGotextI18N NewGotextI18N
func NewGotextI18N() i18n.I18N {
	tags := make(map[string]language.Tag, 0)
	translators := make(map[string]*gotext.Locale, 0)

	return &I18N{
		tags:        tags,
		translators: translators,
		fallback:    nil,
	}
}

// LoadFromDisk LoadFromDisk
func (i18n *I18N) LoadFromDisk(dirRootPath string) error {
	return i18n.LoadwithDefaultLanguageFromDisk(dirRootPath, "en")
}

// LoadwithDefaultLanguageFromDisk LoadwithDefaultLanguageFromDisk
func (i18n *I18N) LoadwithDefaultLanguageFromDisk(dirRootPath string, lang string) error {
	_, err := os.Stat(dirRootPath)
	if err != nil {
		return err
	}

	// recursively go through directory
	walker := func(fileOrDirPath string, info os.FileInfo, err error) error {
		if fileOrDirPath == dirRootPath {
			return nil
		}

		if info.IsDir() {
			name := info.Name()

			_, ok := i18n.translators[name]

			if ok {
				return nil
			}

			tag := language.MustParse(name)

			i18n.tags[name] = tag

			local := gotext.NewLocale(dirRootPath, name)

			if lang == "" {
				if i18n.fallback == nil {
					i18n.fallback = local
					i18n.fallbacktag = tag
				}
			} else {
				if i18n.fallback == nil && name == lang {
					i18n.fallback = local
					i18n.fallbacktag = tag
				}
			}

			i18n.translators[name] = local

			return nil
		}

		base := filepath.Base(fileOrDirPath)
		extension := filepath.Ext(fileOrDirPath)
		domain := base[0 : len(base)-len(extension)]
		dirpath := filepath.Dir(fileOrDirPath)
		dir := filepath.Base(dirpath)

		// skip non po,mo files
		if extension != ".po" && extension != ".mo" {
			return nil
		}

		local, ok := i18n.translators[dir]

		if !ok {
			return nil
		}

		local.AddDomain(domain)

		// fmt.Println(dirpath)
		// fmt.Println(dir)
		// fmt.Println(domain)
		// fmt.Println(fileOrDirPath)
		// fmt.Println(info.Name())
		// fmt.Println()

		return nil
	}

	err = filepath.Walk(dirRootPath, walker)

	if err != nil {
		return err
	}

	tagsarray := make([]language.Tag, 0, len(i18n.tags))

	if i18n.fallback != nil {
		tagsarray = append(tagsarray, i18n.fallbacktag)

		for _, tag := range i18n.tags {
			if tag != i18n.fallbacktag {
				tagsarray = append(tagsarray, tag)
			}
		}
	} else {
		for _, tag := range i18n.tags {
			tagsarray = append(tagsarray, tag)
		}
	}

	if len(tagsarray) > 0 {
		i18n.matcher = language.NewMatcher(tagsarray)
	}

	return nil
}

// LoadFromEmbed LoadFromEmbed
func (i18n *I18N) LoadFromEmbed() error {
	return i18n.LoadwithDefaultLanguageFromEmbed("en")
}

// LoadwithDefaultLanguageFromEmbed LoadwithDefaultLanguageFromEmbed
func (i18n *I18N) LoadwithDefaultLanguageFromEmbed(lang string) error {
	_, err := languages.Open("languages")
	if err != nil {
		return err
	}

	walker := func(fileOrDirPath string, dirEntry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if fileOrDirPath == "languages" || fileOrDirPath == "." || fileOrDirPath == ".." {
			return nil
		}

		if dirEntry.IsDir() {
			name := dirEntry.Name()

			_, ok := i18n.translators[name]

			if ok {
				return nil
			}

			tag := language.MustParse(name)

			i18n.tags[name] = tag

			local := gotext.NewLocale("languages", name)

			if lang == "" {
				if i18n.fallback == nil {
					i18n.fallback = local
					i18n.fallbacktag = tag
				}
			} else {
				if i18n.fallback == nil && name == lang {
					i18n.fallback = local
					i18n.fallbacktag = tag
				}
			}

			i18n.translators[name] = local

			return nil
		}

		base := filepath.Base(fileOrDirPath)
		extension := filepath.Ext(fileOrDirPath)
		domain := base[0 : len(base)-len(extension)]
		dirpath := filepath.Dir(fileOrDirPath)
		dir := filepath.Base(dirpath)

		// skip non po,mo files
		if extension != ".po" && extension != ".mo" {
			return nil
		}

		local, ok := i18n.translators[dir]

		if !ok {
			return nil
		}

		//local.AddDomain(domain)
		//AddDomain from embed files
		b, err := fs.ReadFile(languages, fileOrDirPath)

		if err != nil {
			return err
		}

		po := gotext.NewPo()
		po.Parse(b)

		local.AddTranslator(domain, po)

		// fmt.Println(dirpath)
		// fmt.Println(dir)
		// fmt.Println(domain)
		// fmt.Println(fileOrDirPath)
		// fmt.Println(info.Name())
		// fmt.Println()

		// fmt.Printf("path=%q, isDir=%v,dirEntryName=%v\n", fileOrDirPath, dirEntry.IsDir(), dirEntry.Name())
		return nil

	}

	fs.WalkDir(languages, "languages", walker)

	if err != nil {
		return err
	}

	tagsarray := make([]language.Tag, 0, len(i18n.tags))

	if i18n.fallback != nil {
		tagsarray = append(tagsarray, i18n.fallbacktag)

		for _, tag := range i18n.tags {
			if tag != i18n.fallbacktag {
				tagsarray = append(tagsarray, tag)
			}
		}
	} else {
		for _, tag := range i18n.tags {
			tagsarray = append(tagsarray, tag)
		}
	}

	if len(tagsarray) > 0 {
		i18n.matcher = language.NewMatcher(tagsarray)
	}

	return nil
}

// acceptLanguage Accept-Language from http header,such as zh-CN,zh;q=0.8,en;q=0.2
func (i18n *I18N) findLanguages(acceptLanguage string) []string {
	t, _, _ := language.ParseAcceptLanguage(acceptLanguage)
	tag, _, _ := i18n.matcher.Match(t...)
	base, _ := tag.Base()
	script, _ := tag.Script()
	region, _ := tag.Region()

	// fmt.Println(tag)
	// fmt.Println(fmt.Sprintf("%s-%s-%s", base, script, region))
	// fmt.Println(fmt.Sprintf("%s-%s", base, script))
	// fmt.Println(fmt.Sprintf("%s-%s", base, region))
	// fmt.Println(fmt.Sprintf("%s", base))

	return []string{
		fmt.Sprintf("%s-%s-%s", base, script, region),
		fmt.Sprintf("%s-%s", base, script),
		fmt.Sprintf("%s-%s", base, region),
		fmt.Sprintf("%s", base),
	}
}

func (i18n *I18N) findTranslator(acceptLanguage string) *gotext.Locale {
	var local *gotext.Locale
	var ok bool

	languages := i18n.findLanguages(acceptLanguage)

	for _, language := range languages {
		local, ok = i18n.translators[language]

		if ok {
			break
		}
	}

	if local == nil {
		local = i18n.fallback
	}

	return local
}

// C C
func (i18n *I18N) C(c *gin.Context, key string, num float64, digits uint64) string {
	panic("UnSupported Operation")
}

// O O
func (i18n *I18N) O(c *gin.Context, key string, num float64, digits uint64) string {
	panic("UnSupported Operation")
}

// R R
func (i18n *I18N) R(c *gin.Context, key string, num1 float64, digits1 uint64, num2 float64, digits2 uint64) string {
	panic("UnSupported Operation")
}

// E E
func (i18n *I18N) E(c *gin.Context, err error) map[string]string {
	panic("UnSupported Operation")
}

// T T
func (i18n *I18N) T(c *gin.Context, key string, params ...any) string {
	language := c.GetHeader("Accept-Language")

	local := i18n.findTranslator(language)

	return local.Get(key, params...)
}

// TN TN
func (i18n *I18N) TN(c *gin.Context, key string, plural string, n int, params ...any) string {
	language := c.GetHeader("Accept-Language")

	local := i18n.findTranslator(language)

	return local.GetN(key, plural, n, params...)
}

// TD TD
func (i18n *I18N) TD(c *gin.Context, domain, key string, params ...any) string {
	language := c.GetHeader("Accept-Language")

	local := i18n.findTranslator(language)

	return local.GetD(domain, key, params...)
}

// TND TND
func (i18n *I18N) TND(c *gin.Context, domain, key string, plural string, n int, params ...any) string {
	language := c.GetHeader("Accept-Language")

	local := i18n.findTranslator(language)

	return local.GetND(domain, key, plural, n, params...)
}

// TC TC
func (i18n *I18N) TC(c *gin.Context, key string, ctx string, params ...any) string {
	language := c.GetHeader("Accept-Language")

	local := i18n.findTranslator(language)

	return local.GetC(key, ctx, params...)
}

// TNC TNC
func (i18n *I18N) TNC(c *gin.Context, key string, plural string, n int, ctx string, params ...any) string {
	language := c.GetHeader("Accept-Language")

	local := i18n.findTranslator(language)

	return local.GetNC(key, plural, n, ctx, params...)
}

// TDC TDC
func (i18n *I18N) TDC(c *gin.Context, domain, key string, ctx string, params ...any) string {
	language := c.GetHeader("Accept-Language")

	local := i18n.findTranslator(language)

	return local.GetDC(domain, key, ctx, params...)
}

// TNDC TNDC
func (i18n *I18N) TNDC(c *gin.Context, domain, key string, plural string, n int, ctx string, params ...any) string {
	language := c.GetHeader("Accept-Language")

	local := i18n.findTranslator(language)

	return local.GetNDC(domain, key, plural, n, ctx, params...)
}
