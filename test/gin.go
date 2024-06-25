package test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"glintecoTask/entity"
	"glintecoTask/utils"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
)

type MockGinContext struct {
	c *gin.Context
	w *httptest.ResponseRecorder
}

func NewMockGinContext() MockGinContext {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}
	ctx.Request.Header.Set("Content-Type", "application/json")

	return MockGinContext{
		c: ctx,
		w: w,
	}
}

func (m MockGinContext) SetMiddleware(actor entity.User) {
	m.c.Set(utils.MiddlewareUsernameKey, actor.Username)
	m.c.Set(utils.MiddlewareUserRoleKey, actor.IsAdmin)
	m.c.Set(utils.MiddlewareUserUuidKey, actor.Uuid)
}

func (m MockGinContext) Get(params gin.Params, u url.Values) {
	m.c.Request.Method = "GET"

	// set path params
	m.c.Params = params

	// set query params
	m.c.Request.URL.RawQuery = u.Encode()
}

func (m MockGinContext) Post(content any) {
	m.c.Request.Method = "POST"

	jsonBytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	// the request body must be an io.ReadCloser
	// the bytes buffer though doesn't implement io.Closer,
	// so you wrap it in a no-op closer
	m.c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))
}

func (m MockGinContext) Put(content any, params gin.Params) {
	m.c.Request.Method = "PUT"
	m.c.Params = params

	if !reflect.ValueOf(content).IsNil() {
		jsonBytes, err := json.Marshal(content)
		if err != nil {
			panic(err)
		}

		m.c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))
	}
}

func (m MockGinContext) Delete(params gin.Params) {
	m.c.Request.Method = "DELETE"
	m.c.Params = params
}

func (m MockGinContext) RunTest(f func(ctx *gin.Context)) {
	f(m.c)
}

func (m MockGinContext) ResponseStatus() int {
	return m.w.Code
}

func (m MockGinContext) ResponseBody() []byte {
	return m.w.Body.Bytes()
}
