package main

import (
	"fmt"
	"./data"
	"net/http"
)

// GET /wikis/new
// Show the new wiki form page
func newwiki(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		generateHTML(writer, nil, "layout", "private.navbar", "new.wiki")
	}
}

// POST /signup
// Create the user account
func createwiki(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			danger(err, "Cannot get user from session")
		}
		topic := request.PostFormValue("topic")
		if _, err := user.Createwiki(topic); err != nil {
			danger(err, "Cannot create wiki")
		}
		http.Redirect(writer, request, "/", 302)
	}
}

// GET /wiki/read
// Show the details of the wiki, including the posts and the form to write a post
func readwiki(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	uuid := vals.Get("id")
	wiki, err := data.wikiByUUID(uuid)
	if err != nil {
		error_message(writer, request, "Cannot read wiki")
	} else {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, &wiki, "layout", "public.navbar", "public.wiki")
		} else {
			generateHTML(writer, &wiki, "layout", "private.navbar", "private.wiki")
		}
	}
}

// POST /wiki/post
// Create the post
func postwiki(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			danger(err, "Cannot get user from session")
		}
		body := request.PostFormValue("body")
		uuid := request.PostFormValue("uuid")
		wiki, err := data.wikiByUUID(uuid)
		if err != nil {
			error_message(writer, request, "Cannot read wiki")
		}
		if _, err := user.CreatePost(wiki, body); err != nil {
			danger(err, "Cannot create post")
		}
		url := fmt.Sprint("/wiki/read?id=", uuid)
		http.Redirect(writer, request, url, 302)
	}
}
