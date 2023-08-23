package middlewares

import (
	"github.com/gin-gonic/gin"
	"snowdream.tech/http-server/pkg/i18n"
	"snowdream.tech/http-server/pkg/i18n/gotext"
	"snowdream.tech/http-server/pkg/tools"
)

// I18N I18N
func I18N() gin.HandlerFunc {
	tools.DebugPrintF("[INFO] Starting Middleware %s", "I18N")

	i18ngotext := gotext.NewGotextI18N()
	// i18n.LoadwithDefaultLanguageFromDisk("./languages/", "zh-Hans")
	i18ngotext.LoadFromEmbed()

	const defaultAcceptLanguage = "en"

	return func(c *gin.Context) {
		c.Set(i18n.GoTextKey, i18ngotext)

		var langAccept string

		if langQuery := c.Query("lang"); langQuery != "" {
			langAccept = langQuery
		} else if langCookie, _ := c.Cookie("lang"); langCookie != "" {
			langAccept = langCookie
		} else if langHeader := c.GetHeader("Accept-Language"); langHeader != "" {
			langAccept = langHeader
		} else {
			langAccept = defaultAcceptLanguage
		}

		c.Request.Header.Set("Accept-Language", langAccept)

		c.Next()
	}
}
