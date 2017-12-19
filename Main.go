package main

import (
	"net/http"

	r "github.com/mrizkip/backend-finding-dosen/routers"
)

func main() {
	router := r.NewRouter()
	http.ListenAndServe(":8080", router)
}
