package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/tsuryanto/belajar-golang-httprouter/helper"
)

func TestRouterPatternNamedParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/product/:id/items/:itemid", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")
		itemId := p.ByName("itemid")
		fmt.Fprintf(w, "Product %s Items %s", id, itemId)
	})

	_, body := helper.RunHttpRouterTest(router, "GET", localhost+"/product/1/items/12")
	assert.Equal(t, "Product 1 Items 12", body)
}

func TestRouterPatternCatchAllParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		image := p.ByName("image")
		fmt.Fprintf(w, "Image : %s", image)
	})

	_, body := helper.RunHttpRouterTest(router, "GET", localhost+"/images/small/profile.png")
	assert.Equal(t, "Image : /small/profile.png", body)
}
