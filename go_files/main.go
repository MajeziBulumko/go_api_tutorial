package main

import (
	"log"
	"net/http"
)

type server struct {
	addr string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		switch r.URL.Path {
		case "/api/v1/ping":
			w.Write([]byte("we are here for you"))
			return
		case "/api/v2/ping":
			w.Write([]byte("You are alone"))
			return
		default:
			w.Write([]byte("We need you to stop"))
			return
		}

	}
}
func main() {
	s := &server{addr: ":8080"}
	if err := http.ListenAndServe(s.addr, s); err != nil {
		log.Fatal(err)
	}
}
