package v1

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/samuskitchen/beer-api-clean-arch/beer/delivery/v1/response"
	repository "github.com/samuskitchen/beer-api-clean-arch/beer/repository/postgres"
	"github.com/samuskitchen/beer-api-clean-arch/beer/usecase"
	"github.com/samuskitchen/beer-api-clean-arch/currency/adapter"
	"github.com/samuskitchen/beer-api-clean-arch/domain"
	"github.com/samuskitchen/beer-api-clean-arch/infrastructure/database"
	"github.com/samuskitchen/beer-api-clean-arch/infrastructure/middleware"
)

// BeerRouter struct handler beer
type BeerRouter struct {
	Usecase domain.BeerUsecase
}

// NewBeerHandler constructor
func NewBeerHandler(db *database.Data, client *http.Client) *BeerRouter {
	return &BeerRouter{
		Usecase: usecase.NewBeerUsecase(repository.NewBeerRepository(db), adapter.NewCurrencyAdapter(client)),
	}
}

// GetAllBeersHandler response all the beers.
// swagger:route GET /beers beer getAllBeers
//
// GetAllBeersHandler.
// List all the beers found in the database
//
//     produces:
//      - application/json
//
//	   schemes: http, https
//
//     responses:
//        200: SwaggerAllBeersResponse
//		  404: SwaggerErrorMessage
func (br *BeerRouter) GetAllBeersHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	beers, err := br.Usecase.GetAllBeers(ctx)
	if beers == nil && err != nil {
		_ = middleware.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	if err != nil {
		_ = middleware.HTTPError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	_ = middleware.JSON(w, r, http.StatusOK, beers)
}

// GetOneHandler response one beer by id.
// swagger:route GET /beers/{beerID} beer idBeerPath
//
// GetOneHandler.
// Search for a beer by its ID
//
//     produces:
//      - application/json
//
//	   schemes: http, https
//
//     responses:
//        200: SwaggerBeersResponse
//		  404: SwaggerErrorMessage
func (br *BeerRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "beerID")
	if idStr == "" {
		_ = middleware.HTTPError(w, r, http.StatusBadRequest, errors.New("cannot get id").Error())
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		_ = middleware.HTTPError(w, r, http.StatusBadRequest, errors.New("the id is not a numeric").Error())
		return
	}

	beerResult, err := br.Usecase.GetBeerById(ctx, uint(id))
	if (domain.Beer{}) == beerResult && err != nil {
		_ = middleware.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	if err != nil {
		_ = middleware.HTTPError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	_ = middleware.JSON(w, r, http.StatusOK, beerResult)
}

// CreateHandler Create a new beer.
// swagger:route POST /beers beer beersRequest
//
// CreateHandler.
// Enter a new beer
//
//     consumes:
//     - application/json
//
//     produces:
//      - application/json
//
//	   schemes: http, https
//
//     responses:
//        201: SwaggerSuccessfullyMessage
//		  400: SwaggerErrorMessage
//		  409: SwaggerErrorMessage
//		  422: SwaggerErrorMessage
func (br *BeerRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var beers domain.Beer

	err := json.NewDecoder(r.Body).Decode(&beers)
	if err != nil {
		_ = middleware.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("panic occurred:", err)
		}
	}(r.Body)

	userErrors := beers.Validate()
	if len(userErrors) > 0 {
		_ = middleware.HTTPErrors(w, r, http.StatusUnprocessableEntity, userErrors)
		return
	}

	err = br.Usecase.CreateBeerWithId(ctx, &beers)
	if err != nil {
		_ = middleware.HTTPError(w, r, http.StatusConflict, err.Error())
		return
	}

	_ = middleware.JSONMessages(w, r, http.StatusCreated, "Beer created")
}

// GetOneBoxPriceHandler get the price of a case of beer by its id
// swagger:route GET /beers/{beerID}/boxprice beer idBeerBoxPricePath
//
// GetOneBoxPriceHandler.
// Get the price of a case of beer by its ID
//
//     produces:
//      - application/json
//
//	   schemes: http, https
//
//     responses:
//        200: SwaggerPriceResponse
//		  404: SwaggerErrorMessage
func (br *BeerRouter) GetOneBoxPriceHandler(w http.ResponseWriter, r *http.Request) {
	quantity := 6
	ctx := r.Context()

	idStr := chi.URLParam(r, "beerID")
	currencyStr := r.URL.Query().Get("currency")
	quantityStr := r.URL.Query().Get("quantity")

	if idStr == "" {
		_ = middleware.HTTPError(w, r, http.StatusBadRequest, errors.New("cannot get id").Error())
		return
	}

	if currencyStr == "" || len(currencyStr) < 3 {
		_ = middleware.HTTPError(w, r, http.StatusBadRequest, errors.New("cannot get currency").Error())
		return
	}

	if quantityStr != "" || len(quantityStr) > 0 {
		quantityValue, err := strconv.Atoi(quantityStr)
		if err != nil {
			_ = middleware.HTTPError(w, r, http.StatusBadRequest, errors.New("the quantity is not a numeric").Error())
			return
		}

		if quantityValue != 0 {
			quantity = quantityValue
		}
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		_ = middleware.HTTPError(w, r, http.StatusBadRequest, errors.New("the id is not a numeric").Error())
		return
	}

	total, err := br.Usecase.GetOneBoxPrice(ctx, uint(id), currencyStr, quantity)
	if err != nil {
		_ = middleware.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	totalResponse := response.PriceResponse{
		PriceTotal: total,
	}

	_ = middleware.JSON(w, r, http.StatusOK, totalResponse)
}
