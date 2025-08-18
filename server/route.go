package route

import (
	"prf-manager/interfaces/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
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
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "api_key"},
		ExposedHeaders:   []string{},
		AllowCredentials: true,
		MaxAge:           300,
		Debug:            true,
	}).Handler)

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
			r.Patch("/{id}/chuyen-luu-tru", route.RecordHandler.TransferToArchive)
			r.Patch("/{id}/huy-luu-tru", route.RecordHandler.UndoTransfer)
			r.Post("/search", route.RecordHandler.Search)
		})
		r.Post("/sign-in", route.UserHandler.Login)
	})
	return r
}
