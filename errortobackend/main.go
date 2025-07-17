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

	dsRouter.Patch("/error", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		w.WriteHeader(200)
		sse := datastar.NewSSE(w, r)
		_ = sse
		log.Printf("%+v", r.Form)
		sse.PatchElementTempl(Error())
	})

	dsRouter.Get("/error", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "This is an error page", http.StatusInternalServerError)
	})

	dsRouter.Post("/error", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		w.WriteHeader(200)
		var signals struct {
			Error string `json:"error"`
		}
		if err := datastar.ReadSignals(r, &signals); err != nil {
			log.Printf("Error reading signals: %v", err)
			http.Error(w, "Failed to read signals", http.StatusBadRequest)
			return
		}
		log.Printf("Received error signal: %s", signals.Error)
		sse := datastar.NewSSE(w, r)
		sse.PatchElementTempl(Results("Got error from FE: " + signals.Error))

	})
	log.Printf("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
