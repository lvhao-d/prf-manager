package route

import (
	"prf-manager/interfaces/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Route struct {
	UserHandler   *handler.UserHandler
	CoQuanHandler *handler.CoQuanHandler
	KhoHandler    *handler.KhoHandler
	HoSoHandler   *handler.HoSoHandler
}

func (route *Route) NewRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/users", func(r chi.Router) {
		r.Post("/", route.UserHandler.Create)
	})
	r.Route("/coquans", func(r chi.Router) {
		r.Post("/", route.CoQuanHandler.Create)
		r.Get("/", route.CoQuanHandler.GetAll)
		r.Patch("/{id}", route.CoQuanHandler.Update)
		r.Delete("/{id}", route.CoQuanHandler.Delete)

	})
	return r
}
