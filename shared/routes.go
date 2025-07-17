package shared

import (
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

	return router, dsRouter
}
