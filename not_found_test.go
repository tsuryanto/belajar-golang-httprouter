package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/tsuryanto/belajar-golang-httprouter/helper"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestNotFound(t *testing.T) {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Gak Ketemu")
	})

	_, body := helper.RunHttpRouterTest(router, "GET", localhost+"/404")
	assert.Equal(t, "Gak Ketemu", body)
}
