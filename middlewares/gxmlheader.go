package middlewares

import (
	"encoding/xml"

	"github.com/gin-gonic/gin"
	"snowdream.tech/http-server/pkg/net/http"
	"snowdream.tech/http-server/pkg/tools"
)

// XMLHeader XMLHeader
func XMLHeader() gin.HandlerFunc {
	tools.DebugPrintF("[INFO] Starting Middleware %s", "XMLHeader")

	return func(c *gin.Context) {
		if c.NegotiateFormat(http.OFFEREDALL...) == gin.MIMEXML {
			c.Writer.Write([]byte(xml.Header))
		}

		c.Next()

	}
}
