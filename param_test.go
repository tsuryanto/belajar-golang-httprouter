package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/tsuryanto/belajar-golang-httprouter/helper"
)

func TestParam(t *testing.T) {
	router := httprouter.New()
	router.GET("/product/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")
		fmt.Fprint(w, "Product ", id)
	})

	_, body := helper.RunHttpRouterTest(router, "GET", localhost+"/product/1")
	assert.Equal(t, "Product 1", body)
}
