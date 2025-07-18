package main

import (
	"log"
	"net/http"

	"github.com/blinkinglight/datastar-by-example/shared"
	"github.com/starfederation/datastar-go/datastar"
)

func main() {
	router, dsRouter := shared.Routers()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		Form().Render(r.Context(), w)
	})

	dsRouter.Post("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		w.WriteHeader(200)
		sse := datastar.NewSSE(w, r)
		sse.PatchElementTempl(Results(r.Form.Get("name")))
	})

	log.Printf("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
