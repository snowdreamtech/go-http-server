package i18n

import (
	"github.com/gin-gonic/gin"
)

const (
	// GoTextKey GoTextKey
	GoTextKey = "snowdream.tech/http-server/pkg/i18n/gotextkey"

	// GoPlayGround GoPlayGround
	GoPlayGround = "snowdream.tech/http-server/pkg/i18n/goplayground"
)

// I18N I18N
type I18N interface {
	LoadFromDisk(dirRootPath string) error

	LoadwithDefaultLanguageFromDisk(dirRootPath string, lang string) error

	LoadFromEmbed() error

	LoadwithDefaultLanguageFromEmbed(lang string) error

	// goplayground
	// gotext
	T(c *gin.Context, key string, params ...any) string

	// goplayground
	// creates the cardinal translation for the locale given the 'key', 'num' and 'digit' arguments
	//  and param passed in
	C(c *gin.Context, key string, num float64, digits uint64) string

	// goplayground
	// creates the ordinal translation for the locale given the 'key', 'num' and 'digit' arguments
	// and param passed in
	O(c *gin.Context, key string, num float64, digits uint64) string

	// goplayground
	//  creates the range translation for the locale given the 'key', 'num1', 'digit1', 'num2' and
	//  'digit2' arguments and 'param1' and 'param2' passed in
	R(c *gin.Context, key string, num1 float64, digits1 uint64, num2 float64, digits2 uint64) string

	// goplayground
	// Translate translates all of the ValidationErrors
	E(c *gin.Context, err error) map[string]string

	// gotext
	TN(c *gin.Context, key string, plural string, n int, params ...any) string

	// gotext
	TD(c *gin.Context, domain, key string, params ...any) string

	// gotext
	TND(c *gin.Context, domain, key string, plural string, n int, params ...any) string

	// gotext
	TC(c *gin.Context, key string, ctx string, params ...any) string

	// gotext
	TNC(c *gin.Context, key string, plural string, n int, ctx string, params ...any) string

	// gotext
	TDC(c *gin.Context, domain, key string, ctx string, params ...any) string

	// gotext
	TNDC(c *gin.Context, domain, key string, plural string, n int, ctx string, params ...any) string
}

// Default Default
func Default(c *gin.Context) I18N {
	return c.MustGet(GoTextKey).(I18N)
}

// DefaultGoPlayGround DefaultGoPlayGround
func DefaultGoPlayGround(c *gin.Context) I18N {
	return c.MustGet(GoPlayGround).(I18N)
}
