package main

import (
	"net/http"
)

// func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:

// 		switch r.URL.Path {
// 		case "/api/v1/ping":
// 			w.Write([]byte("we are here for you"))
// 			return
// 		case "/api/v2/ping":
// 			w.Write([]byte("You are alone"))
// 			return
// 		default:
// 			w.Write([]byte("We need you to stop"))
// 			return
// 		}

// }
func main() {
	api := &api{addr: ":8080"}

	mux := http.NewServeMux()

	serve := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)

	serve.ListenAndServe()
}
