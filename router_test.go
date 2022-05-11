package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/tsuryanto/belajar-golang-httprouter/helper"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello World")
	})

	_, body := helper.RunHttpRouterTest(router, "GET", localhost+"/")
	assert.Equal(t, "Hello World", body)
}
