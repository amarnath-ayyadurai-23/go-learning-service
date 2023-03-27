package main

import (
	"fmt"
	"net/http"

	httptreemux "github.com/dimfeld/httptreemux/v5"
)

var (
	router *httptreemux.ContextMux
)

func main() {

	router := httptreemux.NewContextMux()

	router.GET("/:page", func(w http.ResponseWriter, r *http.Request) {
		params := httptreemux.ContextParams(r.Context())
		fmt.Fprintf(w, "GET /%s", params["page"])
	})

	router.Handle(http.MethodGet, "/v2/:page", func(w http.ResponseWriter, r *http.Request) {
		params := httptreemux.ContextParams(r.Context())
		fmt.Fprintf(w, "GET v2 = /%s", params["page"])
	})

	group := router.NewGroup("/api")
	group.GET("/v1/:id", func(w http.ResponseWriter, r *http.Request) {
		ctxData := httptreemux.ContextData(r.Context())
		params := ctxData.Params()
		id := params["id"]

		// Useful for middleware to see which route was hit without dealing with wildcards
		routePath := ctxData.Route()

		// Prints GET /api/v1/:id id=...
		fmt.Fprintf(w, "GET %s id=%s", routePath, id)
	})
	fmt.Println("Server Listening on Port 8080...")
	fmt.Println("Ctrl + C to shutdown ")
	http.ListenAndServe(":8080", router)

}
