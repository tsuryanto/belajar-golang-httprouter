package main

import (
	"embed"
	_ "embed"
	"io/fs"
	"net/http"
	"testing"

	"github.com/tsuryanto/belajar-golang-httprouter/helper"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T) {
	router := httprouter.New()
	directory, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(directory))

	_, body := helper.RunHttpRouterTest(router, "GET", localhost+"/files/hello.txt")
	assert.Equal(t, "Hello Httprouter Serve File", body)
}
