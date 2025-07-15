package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/starfederation/datastar-go/datastar"
)

func main() {
	router := chi.NewMux()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	dsRouter := chi.NewRouter()
	router.Use(middleware.RouteHeaders().Route("Datastar-Request", "true", middleware.New(dsRouter)).Handler)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		Intersect().Render(r.Context(), w)
	})

	dsRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		sse := datastar.NewSSE(w, r)
		sse.PatchElementTempl(LoadMore(), datastar.WithSelectorID("intersect"), datastar.WithModeAppend())
	})

	log.Printf("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
