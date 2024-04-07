package main

import (
	"net/http"

	"github.com/fabiov3105/westminsteripvc_front/drivers/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8001", nil)
}