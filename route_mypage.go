package main

import (
	"./data"
	"net/http"
)

// GET /mypage
// Show the mypage
func mypage(writer http.ResponeWriter, request *http.Request) {
	t := parseTemplateFile("layout", "private.navber", "login")
	t.Execute(writer, nil)
}
