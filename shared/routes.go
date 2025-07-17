package shared

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ituoga/toolbox/hotreload"
)

// returns Router and DataStar Router
func Routers() (chi.Router, chi.Router) {
	router := chi.NewMux()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	dsRouter := chi.NewRouter()
	router.Use(middleware.RouteHeaders().Route("Datastar-Request", "true", middleware.New(dsRouter)).Handler)

	dsRouter.Get("/hotreload", hotreload.Handler)

	router.Get("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))).ServeHTTP)

	return router, dsRouter
}

func Serve(router chi.Router) {
	log.Printf("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
