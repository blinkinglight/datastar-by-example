package main

import (
	"net/http"

	"github.com/blinkinglight/datastar-by-example/shared"
)

func main() {
	router, _ := shared.Routers()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		Main().Render(r.Context(), w)
	})

	shared.Serve(router)
}
