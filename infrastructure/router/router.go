package infrastructure

import (
	"net/http"

	"github.com/go-chi/chi"
	v1 "github.com/samuskitchen/beer-api-clean-arch/application/beer/delivery/v1"
	"github.com/samuskitchen/beer-api-clean-arch/infrastructure/database"
)

// Routes returns the API V1 Handler with configuration.
func Routes(conn *database.Data) http.Handler {
	router := chi.NewRouter()

	br := v1.NewBeerHandler(conn, http.DefaultClient)
	router.Mount("/beer", routesBeer(br))

	return router
}

// routesUser returns beer router with each endpoint.
func routesBeer(handler *v1.BeerRouter) http.Handler {
	router := chi.NewRouter()

	router.Get("/", handler.GetAllBeersHandler)
	router.Get("/{beerID}", handler.GetOneHandler)
	router.Get("/{beerID}/boxprice", handler.GetOneBoxPriceHandler)
	router.Post("/", handler.CreateHandler)

	return router
}
