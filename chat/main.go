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
		ReverseIntersect().Render(r.Context(), w)
	})

	dsRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		sse := datastar.NewSSE(w, r)
		sse.PatchElementTempl(LoadMore(), datastar.WithSelectorID("intersect"), datastar.WithModePrepend())
	})

	log.Printf("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
