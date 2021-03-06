package main

import (
	"net/http"
	"time"
)

func main() {

	// handle static assets
	mux := http.NewServeMux()

	files := http.FileServer(http.Dir(config.Static))
	// Deleate Prefix
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	//
	// all route patterns matched here
	// route handler functions defined in other files
	//

	// index
	mux.HandleFunc("/", index)
	// error
	mux.HandleFunc("/err", err)

	// defined in route_auth.go
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	// defined in route_thread.go
//	mux.HandleFunc("/thread/new", newThread)
//	mux.HandleFunc("/thread/create", createThread)
//	mux.HandleFunc("/thread/post", postThread)
//	mux.HandleFunc("/thread/read", readThread)

  // defined in route_mypage.go
  mux.HandleFunc("/mypage", mypage)

	// defined in route_wiki.go
	mux.HandleFunc("/wiki/new", newWiki)
	mux.HandleFunc("/wiki/create", createWiki)
	mux.HandleFunc("/wiki/post", postWiki)
	mux.HandleFunc("/wiki/read", readWiki)

	// starting up the server
	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
