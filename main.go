package main

import (
	"net/http"
	"tradeEngine/router"
)

func main() {
	routers := router.NewRouter()
	http.ListenAndServe(":80", routers)
}
