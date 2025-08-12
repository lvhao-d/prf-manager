package route

import (
	"prf-manager/interfaces/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Route struct {
	UserHandler      *handler.UserHandler
	AgencyHandler    *handler.AgencyHandler
	WareHouseHandler *handler.WareHouseHandler
	RecordHandler    *handler.RecordHandler
}

func (route *Route) NewRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		r.Route("/agency", func(r chi.Router) {
			r.Post("/", route.AgencyHandler.Create)
			r.Get("/", route.AgencyHandler.GetAll)
			r.Patch("/{id}", route.AgencyHandler.Update)
			r.Delete("/{id}", route.AgencyHandler.Delete)

		})
		r.Route("/warehouse", func(r chi.Router) {
			r.Post("/", route.WareHouseHandler.Create)
			r.Get("/", route.WareHouseHandler.GetAll)
			r.Patch("/{id}", route.WareHouseHandler.Update)
			r.Delete("/{id}", route.WareHouseHandler.Delete)

		})
		r.Route("/record", func(r chi.Router) {
			r.Post("/", route.RecordHandler.Create)
			r.Get("/", route.RecordHandler.GetAll)
			r.Patch("/{id}", route.RecordHandler.Update)
			r.Delete("/{id}", route.RecordHandler.Delete)
		})
		r.Post("/login", route.UserHandler.Login)
	})
	return r
}
