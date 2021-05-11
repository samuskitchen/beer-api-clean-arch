package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi"
	"github.com/samuskitchen/beer-api-clean-arch/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	v1 "github.com/samuskitchen/beer-api-clean-arch/beer/delivery/v1"
	repoMock "github.com/samuskitchen/beer-api-clean-arch/domain/mocks"
)

// dataBeers is data for test
func dataBeers() []domain.Beer {
	now := time.Now()

	return []domain.Beer{
		{
			ID:        uint(1),
			Name:      "Golden",
			Brewery:   "Kross",
			Country:   "Chile",
			Price:     10.5,
			Currency:  "CLP",
			CreatedAt: now,
		},
		{
			ID:        uint(2),
			Name:      "Club Colombia",
			Brewery:   "Bavaria",
			Country:   "Colombia",
			Price:     2550,
			Currency:  "COP",
			CreatedAt: now,
		},
	}
}

func TestBeersRouter_GetAllBeersHandler(t *testing.T) {

	t.Run("Error Get All Beers Handler", func(tt *testing.T) {

		newRequest := httptest.NewRequest(http.MethodGet, "/api/v1/beer/", nil)
		newRecorder := httptest.NewRecorder()
		mockUsecase := &repoMock.BeerUsecase{}

		testBeersHandler := &v1.BeerRouter{Usecase: mockUsecase}
		mockUsecase.On("GetAllBeers", mock.Anything).Return(nil, errors.New("error trace test"))

		testBeersHandler.GetAllBeersHandler(newRecorder, newRequest)
		mockUsecase.AssertExpectations(tt)
	})

	t.Run("(Not Found) Get All Beers Handler", func(tt *testing.T) {

		newRequest := httptest.NewRequest(http.MethodGet, "/api/v1/beer/", nil)
		newRecorder := httptest.NewRecorder()
		mockUsecase := &repoMock.BeerUsecase{}

		testBeersHandler := &v1.BeerRouter{Usecase: mockUsecase}
		mockUsecase.On("GetAllBeers", mock.Anything).Return(nil, nil)

		testBeersHandler.GetAllBeersHandler(newRecorder, newRequest)
		mockUsecase.AssertExpectations(tt)
	})

	t.Run("Get All Beers Handler", func(tt *testing.T) {

		newRequest := httptest.NewRequest(http.MethodGet, "/api/v1/beer/", nil)
		newRecorder := httptest.NewRecorder()
		mockUsecase := &repoMock.BeerUsecase{}

		testBeersHandler := &v1.BeerRouter{Usecase: mockUsecase}
		mockUsecase.On("GetAllBeers", mock.Anything).Return(dataBeers(), nil)

		testBeersHandler.GetAllBeersHandler(newRecorder, newRequest)
		mockUsecase.AssertExpectations(tt)
	})
}

func TestBeersRouter_GetOneHandler(t *testing.T) {

	t.Run("Error Param Get One Handler", func(tt *testing.T) {

		newRecorder := httptest.NewRecorder()
		newRequest := httptest.NewRequest(http.MethodGet, "/api/v1/beer/{beerID}", nil)

		mockUsecase := &repoMock.BeerUsecase{}
		testBeersHandler := &v1.BeerRouter{Usecase: mockUsecase}

		testBeersHandler.GetOneHandler(newRecorder, newRequest)
		mockUsecase.AssertExpectations(tt)
	})

	t.Run("Error SQL Get One Handler", func(tt *testing.T) {

		newRecorder := httptest.NewRecorder()
		newRequest := httptest.NewRequest(http.MethodGet, "/api/v1/beer/{beerID}", nil)

		newRequestCtx := chi.NewRouteContext()
		newRequestCtx.URLParams.Add("beerID", "1")

		newRequest = newRequest.WithContext(context.WithValue(newRequest.Context(), chi.RouteCtxKey, newRequestCtx))
		mockUsecase := &repoMock.BeerUsecase{}

		testBeersHandler := &v1.BeerRouter{Usecase: mockUsecase}
		mockUsecase.On("GetBeerById", mock.Anything, mock.Anything).Return(domain.Beer{}, errors.New("error sql")).Once()

		testBeersHandler.GetOneHandler(newRecorder, newRequest)
		mockUsecase.AssertExpectations(tt)
	})

	t.Run("Error Parse Get One Handler", func(tt *testing.T) {

		newRecorder := httptest.NewRecorder()
		newRequest := httptest.NewRequest(http.MethodGet, "/api/v1/beer/{beerID}", nil)

		newRequestCtx := chi.NewRouteContext()
		newRequestCtx.URLParams.Add("beerID", "no")

		newRequest = newRequest.WithContext(context.WithValue(newRequest.Context(), chi.RouteCtxKey, newRequestCtx))
		mockUsecase := &repoMock.BeerUsecase{}

		testBeersHandler := &v1.BeerRouter{Usecase: mockUsecase}

		testBeersHandler.GetOneHandler(newRecorder, newRequest)
		mockUsecase.AssertExpectations(tt)
	})

	t.Run("(Not Found) Get One Handler", func(tt *testing.T) {

		newRecorder := httptest.NewRecorder()
		newRequest := httptest.NewRequest(http.MethodGet, "/api/v1/beer/{beerID}", nil)

		newRequestCtx := chi.NewRouteContext()
		newRequestCtx.URLParams.Add("beerID", "1")

		newRequest = newRequest.WithContext(context.WithValue(newRequest.Context(), chi.RouteCtxKey, newRequestCtx))
		mockUsecase := &repoMock.BeerUsecase{}

		testBeersHandler := &v1.BeerRouter{Usecase: mockUsecase}
		mockUsecase.On("GetBeerById", mock.Anything, mock.Anything).Return(domain.Beer{}, nil).Once()

		testBeersHandler.GetOneHandler(newRecorder, newRequest)
		mockUsecase.AssertExpectations(tt)
	})

	t.Run("Get One Handler", func(tt *testing.T) {

		newRecorder := httptest.NewRecorder()
		newRequest := httptest.NewRequest(http.MethodGet, "/api/v1/beer/{beerID}", nil)

		newRequestCtx := chi.NewRouteContext()
		newRequestCtx.URLParams.Add("beerID", "1")

		newRequest = newRequest.WithContext(context.WithValue(newRequest.Context(), chi.RouteCtxKey, newRequestCtx))
		mockUsecase := &repoMock.BeerUsecase{}

		testBeersHandler := &v1.BeerRouter{Usecase: mockUsecase}
		mockUsecase.On("GetBeerById", mock.Anything, mock.Anything).Return(dataBeers()[0], nil).Once()

		testBeersHandler.GetOneHandler(newRecorder, newRequest)
		mockUsecase.AssertExpectations(tt)
	})

}

func TestBeersRouter_GetOneBoxPriceHandler(t *testing.T) {

	t.Run("(Error Param beerID) Get One Handler", func(tt *testing.T) {

		newRecorder := httptest.NewRecorder()
		newRequest := httptest.NewRequest(http.MethodGet, "/api/v1/beer/{beerID}/boxprice", nil)

		mockUsecase := &repoMock.BeerUsecase{}
		testBeersHandler := &v1.BeerRouter{Usecase: mockUsecase}

		testBeersHandler.GetOneBoxPriceHandler(newRecorder, newRequest)
		mockUsecase.AssertExpectations(tt)
	})

	t.Run("(Error Param Currency) Get One Handler", func(tt *testing.T) {

		newRecorder := httptest.NewRecorder()
		newRequest := httptest.NewRequest(http.MethodGet, "/api/v1/beer/{beerID}/boxprice", nil)

		newRequestCtx := chi.NewRouteContext()
		newRequestCtx.URLParams.Add("beerID", "1")

		newRequest = newRequest.WithContext(context.WithValue(newRequest.Context(), chi.RouteCtxKey, newRequestCtx))
		mockUsecase := &repoMock.BeerUsecase{}
		testBeersHandler := &v1.BeerRouter{Usecase: mockUsecase}

		testBeersHandler.GetOneBoxPriceHandler(newRecorder, newRequest)
		mockUsecase.AssertExpectations(tt)
	})

	t.Run("(Error Param Quantity) Get One Handler", func(tt *testing.T) {

		newRecorder := httptest.NewRecorder()
		newRequest := httptest.NewRequest(http.MethodGet, "/api/v1/beer/{beerID}/boxprice", nil)

		newRequestCtx := chi.NewRouteContext()
		newRequestCtx.URLParams.Add("beerID", "1")

		queryParam := newRequest.URL.Query()
		queryParam.Add("currency", "COP")
		queryParam.Add("quantity", "no")
		newRequest.URL.RawQuery = queryParam.Encode()

		newRequest = newRequest.WithContext(context.WithValue(newRequest.Context(), chi.RouteCtxKey, newRequestCtx))
		mockUsecase := &repoMock.BeerUsecase{}
		testBeersHandler := &v1.BeerRouter{Usecase: mockUsecase}

		testBeersHandler.GetOneBoxPriceHandler(newRecorder, newRequest)
		mockUsecase.AssertExpectations(tt)
	})

	t.Run("(Error Parse beerID) Get One Handler", func(tt *testing.T) {

		newRecorder := httptest.NewRecorder()
		newRequest := httptest.NewRequest(http.MethodGet, "/api/v1/beer/{beerID}/boxprice", nil)

		newRequestCtx := chi.NewRouteContext()
		newRequestCtx.URLParams.Add("beerID", "no")

		queryParam := newRequest.URL.Query()
		queryParam.Add("currency", "COP")
		queryParam.Add("quantity", "1")
		newRequest.URL.RawQuery = queryParam.Encode()

		newRequest = newRequest.WithContext(context.WithValue(newRequest.Context(), chi.RouteCtxKey, newRequestCtx))
		mockUsecase := &repoMock.BeerUsecase{}
		testBeersHandler := &v1.BeerRouter{Usecase: mockUsecase}

		testBeersHandler.GetOneBoxPriceHandler(newRecorder, newRequest)
		mockUsecase.AssertExpectations(tt)
	})

	t.Run("(Error Get Beer) Get One Handler", func(tt *testing.T) {

		newRecorder := httptest.NewRecorder()
		newRequest := httptest.NewRequest(http.MethodGet, "/api/v1/beer/{beerID}/boxprice", nil)

		newRequestCtx := chi.NewRouteContext()
		newRequestCtx.URLParams.Add("beerID", "1")

		queryParam := newRequest.URL.Query()
		queryParam.Add("currency", "COP")
		newRequest.URL.RawQuery = queryParam.Encode()

		newRequest = newRequest.WithContext(context.WithValue(newRequest.Context(), chi.RouteCtxKey, newRequestCtx))
		mockUsecase := &repoMock.BeerUsecase{}
		testBeersHandler := &v1.BeerRouter{Usecase: mockUsecase}

		mockUsecase.On("GetOneBoxPrice", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(float64(0), errors.New("error sql")).Once()

		testBeersHandler.GetOneBoxPriceHandler(newRecorder, newRequest)
		mockUsecase.AssertExpectations(tt)
	})

	t.Run("(No Found Beer) Get One Handler", func(tt *testing.T) {

		newRecorder := httptest.NewRecorder()
		newRequest := httptest.NewRequest(http.MethodGet, "/api/v1/beer/{beerID}/boxprice", nil)

		newRequestCtx := chi.NewRouteContext()
		newRequestCtx.URLParams.Add("beerID", "1")

		queryParam := newRequest.URL.Query()
		queryParam.Add("currency", "COP")
		newRequest.URL.RawQuery = queryParam.Encode()

		newRequest = newRequest.WithContext(context.WithValue(newRequest.Context(), chi.RouteCtxKey, newRequestCtx))
		mockUsecase := &repoMock.BeerUsecase{}
		testBeersHandler := &v1.BeerRouter{Usecase: mockUsecase}

		mockUsecase.On("GetOneBoxPrice", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(float64(1), nil).Once()

		testBeersHandler.GetOneBoxPriceHandler(newRecorder, newRequest)
		mockUsecase.AssertExpectations(tt)
	})

	t.Run("(Error Get Currency) Get One Handler", func(tt *testing.T) {

		newRecorder := httptest.NewRecorder()
		newRequest := httptest.NewRequest(http.MethodGet, "/api/v1/beer/{beerID}/boxprice", nil)

		newRequestCtx := chi.NewRouteContext()
		newRequestCtx.URLParams.Add("beerID", "1")

		queryParam := newRequest.URL.Query()
		queryParam.Add("currency", "EUR")
		newRequest.URL.RawQuery = queryParam.Encode()

		newRequest = newRequest.WithContext(context.WithValue(newRequest.Context(), chi.RouteCtxKey, newRequestCtx))
		mockUsecase := &repoMock.BeerUsecase{}
		mockCurrencyRepository := &repoMock.CurrencyLayerRepository{}
		testBeersHandler := &v1.BeerRouter{Usecase: mockUsecase}

		mockUsecase.On("GetOneBoxPrice", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(float64(1), nil).Once()
		mockCurrencyRepository.On("GetCurrency", mock.Anything, mock.Anything).Return([]float64{0, 0}, errors.New("error get currency")).Once()

		testBeersHandler.GetOneBoxPriceHandler(newRecorder, newRequest)
		mockUsecase.AssertExpectations(tt)
	})

	t.Run("Get One Box Price Handler", func(tt *testing.T) {

		newRecorder := httptest.NewRecorder()
		newRequest := httptest.NewRequest(http.MethodGet, "/api/v1/beer/{beerID}/boxprice", nil)

		newRequestCtx := chi.NewRouteContext()
		newRequestCtx.URLParams.Add("beerID", "1")

		queryParam := newRequest.URL.Query()
		queryParam.Add("currency", "EUR")
		newRequest.URL.RawQuery = queryParam.Encode()

		newRequest = newRequest.WithContext(context.WithValue(newRequest.Context(), chi.RouteCtxKey, newRequestCtx))
		mockUsecase := &repoMock.BeerUsecase{}
		mockCurrencyRepository := &repoMock.CurrencyLayerRepository{}
		testBeersHandler := &v1.BeerRouter{Usecase: mockUsecase}

		mockUsecase.On("GetOneBoxPrice", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(float64(1), nil).Once()
		mockCurrencyRepository.On("GetCurrency", mock.Anything, mock.Anything).Return([]float64{0.824798, 3420.45}, nil).Once()

		testBeersHandler.GetOneBoxPriceHandler(newRecorder, newRequest)
		mockUsecase.AssertExpectations(tt)
	})

}

func TestBeersRouter_CreateHandler(t *testing.T) {

	t.Run("Error Body Create Handler", func(tt *testing.T) {

		newRequest := httptest.NewRequest(http.MethodPost, "/api/v1/beer/", bytes.NewReader(nil))
		newRecorder := httptest.NewRecorder()
		mockUsecase := &repoMock.BeerUsecase{}

		testBeersHandler := &v1.BeerRouter{Usecase: mockUsecase}

		testBeersHandler.CreateHandler(newRecorder, newRequest)
		mockUsecase.AssertExpectations(tt)
	})

	t.Run("Validate Create Handler", func(tt *testing.T) {

		var userTest = dataBeers()[0]
		userTest.ID = 0
		userTest.Name = ""
		userTest.Brewery = ""
		userTest.Country = ""
		userTest.Price = 0
		userTest.Currency = ""

		marshal, err := json.Marshal(userTest)
		assert.NoError(tt, err)

		newRequest := httptest.NewRequest(http.MethodPost, "/api/v1/beer/", bytes.NewReader(marshal))
		newRecorder := httptest.NewRecorder()
		mockUsecase := &repoMock.BeerUsecase{}

		testBeersHandler := &v1.BeerRouter{Usecase: mockUsecase}

		testBeersHandler.CreateHandler(newRecorder, newRequest)
		mockUsecase.AssertExpectations(tt)

	})

	t.Run("Error SQL With ID Create Handler", func(tt *testing.T) {

		marshal, err := json.Marshal(dataBeers()[0])
		assert.NoError(tt, err)

		newRequest := httptest.NewRequest(http.MethodPost, "/api/v1/beer/", bytes.NewReader(marshal))
		newRecorder := httptest.NewRecorder()
		mockUsecase := &repoMock.BeerUsecase{}

		testBeersHandler := &v1.BeerRouter{Usecase: mockUsecase}
		mockUsecase.On("CreateBeerWithId", mock.Anything, mock.Anything).Return(errors.New("error sql"))

		testBeersHandler.CreateHandler(newRecorder, newRequest)
		mockUsecase.AssertExpectations(tt)
	})

	t.Run("Create With ID Handler", func(tt *testing.T) {

		marshal, err := json.Marshal(dataBeers()[0])
		assert.NoError(tt, err)

		newRequest := httptest.NewRequest(http.MethodPost, "/api/v1/beer/", bytes.NewReader(marshal))
		newRecorder := httptest.NewRecorder()
		mockUsecase := &repoMock.BeerUsecase{}

		testBeersHandler := &v1.BeerRouter{Usecase: mockUsecase}
		mockUsecase.On("CreateBeerWithId", mock.Anything, mock.Anything).Return(nil)

		testBeersHandler.CreateHandler(newRecorder, newRequest)
		mockUsecase.AssertExpectations(tt)
	})

}
