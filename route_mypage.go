package main

import (
//	"./data"
	"net/http"
)

// GET /mypage
// Show the mypage
func mypage(writer http.ResponseWriter, request *http.Request) {
	t := parseTemplateFiles("layout", "private.navbar", "mypage")
	t.Execute(writer, nil)
}
