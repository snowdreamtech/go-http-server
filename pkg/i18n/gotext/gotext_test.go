package gotext

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	i18n := NewGotextI18N()

	i18n.LoadwithDefaultLanguageFromDisk("./languages/", "zh-Hans")
	// i18n.Load("./languages/")

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("POST", "/", nil)
	c.Request.Header.Set("Accept-Language", "zh")

	str := i18n.T(c, "hello world!")
	assert.Equal(t, "你好 世界！[简体中文]", str)
}

func TestT(t *testing.T) {
	i18n := NewGotextI18N()

	i18n.LoadFromDisk("./languages/")

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("POST", "/", nil)
	c.Request.Header.Set("Accept-Language", "zh")
	str := i18n.T(c, "hello world!")
	assert.Equal(t, "你好 世界！[简体中文]", str)

	c.Request.Header.Set("Accept-Language", "zh-Hans")
	str = i18n.T(c, "hello world!")
	assert.Equal(t, "你好 世界！[简体中文]", str)

	c.Request.Header.Set("Accept-Language", "zh-Hant")
	str = i18n.T(c, "hello world!")
	assert.Equal(t, "你好 世界！[繁体中文]", str)
}

func TestTN(t *testing.T) {
	i18n := NewGotextI18N()

	i18n.LoadFromDisk("./languages/")

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("POST", "/", nil)
	c.Request.Header.Set("Accept-Language", "en")

	str := i18n.TN(c, "One with var: %s", "Several with vars: %s", 1, "VALUE")
	assert.Equal(t, "This one is the singular: VALUE", str)

	str = i18n.TN(c, "One with var: %s", "Several with vars: %s", 2, "VALUE")
	assert.Equal(t, "This one is the plural: VALUE", str)

	str = i18n.TN(c, "One with var: %s", "Several with vars: %s", 3, "VALUE")
	assert.Equal(t, "This one is the plural: VALUE", str)
}

func TestTD(t *testing.T) {
	i18n := NewGotextI18N()

	i18n.LoadFromDisk("./languages/")

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("POST", "/", nil)
	c.Request.Header.Set("Accept-Language", "zh")

	str := i18n.T(c, "hello world!")
	assert.Equal(t, "你好 世界！[简体中文]", str)

	str = i18n.TD(c, "test", "hello world!")
	assert.Equal(t, "你好 世界！Test!", str)
}

func TestTND(t *testing.T) {
	i18n := NewGotextI18N()

	i18n.LoadFromDisk("./languages/")

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("POST", "/", nil)
	c.Request.Header.Set("Accept-Language", "en")

	str := i18n.TND(c, "default", "One with var: %s", "Several with vars: %s", 1, "VALUE")
	assert.Equal(t, "This one is the singular: VALUE", str)

	str = i18n.TND(c, "default", "One with var: %s", "Several with vars: %s", 2, "VALUE")
	assert.Equal(t, "This one is the plural: VALUE", str)

	str = i18n.TND(c, "default", "One with var: %s", "Several with vars: %s", 3, "VALUE")
	assert.Equal(t, "This one is the plural: VALUE", str)
}

func TestTC(t *testing.T) {
	i18n := NewGotextI18N()

	i18n.LoadFromDisk("./languages/")

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("POST", "/", nil)
	c.Request.Header.Set("Accept-Language", "en")

	str := i18n.TC(c, "Some random in a context", "Ctx")
	assert.Equal(t, "Some random Translation in a context", str)
}

func TestTNC(t *testing.T) {
	i18n := NewGotextI18N()

	i18n.LoadFromDisk("./languages/")

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("POST", "/", nil)
	c.Request.Header.Set("Accept-Language", "en")

	str := i18n.TNC(c, "One with var: %s", "Several with vars: %s", 1, "Ctx", "VALUE")
	assert.Equal(t, "This one is the singular in a Ctx context: VALUE", str)

	str = i18n.TNC(c, "One with var: %s", "Several with vars: %s", 2, "Ctx", "VALUE")
	assert.Equal(t, "This one is the plural in a Ctx context: VALUE", str)

	str = i18n.TNC(c, "One with var: %s", "Several with vars: %s", 3, "Ctx", "VALUE")
	assert.Equal(t, "This one is the plural in a Ctx context: VALUE", str)
}

func TestTDC(t *testing.T) {
	i18n := NewGotextI18N()

	i18n.LoadFromDisk("./languages/")

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("POST", "/", nil)
	c.Request.Header.Set("Accept-Language", "en")

	str := i18n.TDC(c, "default", "Some random in a context", "Ctx")
	assert.Equal(t, "Some random Translation in a context", str)
}

func TestTNDC(t *testing.T) {
	i18n := NewGotextI18N()

	i18n.LoadFromDisk("./languages/")

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("POST", "/", nil)
	c.Request.Header.Set("Accept-Language", "en")

	str := i18n.TNDC(c, "default", "One with var: %s", "Several with vars: %s", 1, "Ctx", "VALUE")
	assert.Equal(t, "This one is the singular in a Ctx context: VALUE", str)

	str = i18n.TNDC(c, "default", "One with var: %s", "Several with vars: %s", 2, "Ctx", "VALUE")
	assert.Equal(t, "This one is the plural in a Ctx context: VALUE", str)

	str = i18n.TNDC(c, "default", "One with var: %s", "Several with vars: %s", 3, "Ctx", "VALUE")
	assert.Equal(t, "This one is the plural in a Ctx context: VALUE", str)
}
