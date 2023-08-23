package gin

import (
	"github.com/gin-gonic/gin"
)

// NewMiddleware return a new instance of a gin middleware.
func NewMiddleware(defaultAcceptLanguage string) gin.HandlerFunc {

	return func(c *gin.Context) {
		langAccept := c.GetHeader("Accept-Language")

		if langAccept == "" {

			if langQuery := c.Query("lang"); langQuery != "" {
				langAccept = langQuery
			} else if langCookie, _ := c.Cookie("lang"); langCookie != "" {
				langAccept = langCookie
			} else {
				langAccept = defaultAcceptLanguage
			}

			c.Request.Header.Set("Accept-Language", langAccept)
		}

		c.Next()
	}
}

// NewDefaultMiddleware return a new instance of a gin middleware.
func NewDefaultMiddleware() gin.HandlerFunc {
	const defaultAcceptLanguage = "en"

	return NewMiddleware(defaultAcceptLanguage)
}
