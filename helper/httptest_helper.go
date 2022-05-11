package helper

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/julienschmidt/httprouter"
)

// Buat versi struct
type HttpTest struct {
	Param  string
	Method string
}

func (test *HttpTest) Request() *http.Request {
	return httptest.NewRequest(test.Method, test.Param, nil)
}

func (test *HttpTest) Recorder() *httptest.ResponseRecorder {
	return httptest.NewRecorder()
}

func (test *HttpTest) ResultBody(recorder *httptest.ResponseRecorder) (*http.Response, []byte) {
	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)
	return result, body
}

func RunHttpTest(handler func(w http.ResponseWriter, r *http.Request), method string, param string) (*http.Response, string) {
	ht := HttpTest{
		Param:  param,
		Method: method,
	}
	recorder := ht.Recorder()
	handler(recorder, ht.Request())

	result, body := ht.ResultBody(recorder)
	return result, string(body)
}

func RunHttpRouterTest(router *httprouter.Router, method string, param string) (*http.Response, string) {
	ht := HttpTest{
		Param:  param,
		Method: method,
	}

	request := httptest.NewRequest(method, param, nil)
	recorder := ht.Recorder()
	router.ServeHTTP(recorder, request)

	result, body := ht.ResultBody(recorder)

	return result, string(body)
}

// func RunHttpRouterMiddlewareTest(middleware LogMiddleware, method string, param string) (*http.Response, string) {
// 	ht := HttpTest{
// 		Param:  param,
// 		Method: method,
// 	}

// 	request := httptest.NewRequest(method, param, nil)
// 	recorder := ht.Recorder()
// 	middleware.ServeHTTP(recorder, request)

// 	result, body := ht.ResultBody(recorder)

// 	return result, string(body)
// }
