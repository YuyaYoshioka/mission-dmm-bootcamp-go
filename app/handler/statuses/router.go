package statuses

import (
	"net/http"

	"yatter-backend-go/app/app"
	"yatter-backend-go/app/handler/auth"
	// "yatter-backend-go/app/handler/auth"

	"github.com/go-chi/chi"
)

// Implementation of handler
type handler struct {
	app *app.App
}

// Create Handler for `/v1/statuses/`
func NewRouter(app *app.App) http.Handler {
	router := chi.NewRouter()
	h := &handler{app: app}
	router.Route("/", func(r chi.Router) {
		r.Use(auth.Middleware(app))
		r.Post("/", h.Create)
	})
	router.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.Get)
	})

	return router
}
