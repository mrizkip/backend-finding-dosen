package main

import (
	"net/http"
	"os"

	r "github.com/mrizkip/backend-finding-dosen/routers"
)

func main() {
	router := r.NewRouter()
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
