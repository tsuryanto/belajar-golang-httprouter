package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tsuryanto/belajar-golang-httprouter/helper"

	"github.com/julienschmidt/httprouter"
)

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		fmt.Fprint(w, "Panic : ", i)
	}
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.WriteHeader(http.StatusInternalServerError) // Write header : Error 500 (Internal Server Error).
		panic("Ups")
	})

	_, body := helper.RunHttpRouterTest(router, "GET", localhost+"/")
	assert.Equal(t, "Panic : Ups", body)
}
