package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/tsuryanto/belajar-golang-httprouter/helper"
)

func TestMethodNotAllowed(t *testing.T) {
	router := httprouter.New()
	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Gak Boleh")
	})

	router.POST("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "POST")
	})

	_, body := helper.RunHttpRouterTest(router, "GET", localhost+"/") // Gak Boleh
	// _, body := helper.RunHttpRouterTest(router, "POST", localhost+"/") // Gak Boleh
	assert.Equal(t, "Gak Boleh", body)
}
