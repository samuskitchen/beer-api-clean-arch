package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/samuskitchen/beer-api-clean-arch/application/beer/usecase"
	"github.com/samuskitchen/beer-api-clean-arch/domain"
	repoMock "github.com/samuskitchen/beer-api-clean-arch/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func Test_beerUsecase_CreateBeerWithId(t *testing.T) {

	t.Run("Error SQL With ID Create Usecase", func(tt *testing.T) {

		mockUsecase := &repoMock.BeerUsecase{}
		mockBeerRepository := &repoMock.BeerRepository{}
		mockCurrencyLayer := &repoMock.CurrencyLayerRepository{}

		testBeerUsecase := usecase.NewBeerUsecase(mockBeerRepository, mockCurrencyLayer)
		mockBeerRepository.On("CreateBeerWithId", mock.Anything, mock.AnythingOfType("*domain.Beer")).Return(errors.New("error trace test"))

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := testBeerUsecase.CreateBeerWithId(ctx, &dataBeers()[0])
		assert.Error(tt, err)
		mockUsecase.AssertExpectations(tt)
	})

	t.Run("Create With ID Usecase", func(tt *testing.T) {

		mockUsecase := &repoMock.BeerUsecase{}
		mockBeerRepository := &repoMock.BeerRepository{}
		mockCurrencyLayer := &repoMock.CurrencyLayerRepository{}

		testBeerUsecase := usecase.NewBeerUsecase(mockBeerRepository, mockCurrencyLayer)
		mockBeerRepository.On("CreateBeerWithId", mock.Anything, mock.AnythingOfType("*domain.Beer")).Return(nil)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := testBeerUsecase.CreateBeerWithId(ctx, &dataBeers()[0])
		assert.NoError(tt, err)
		mockUsecase.AssertExpectations(tt)
	})

}

func Test_beerUsecase_GetAllBeers(t *testing.T) {

	t.Run("Error Get All Beers Usecase", func(tt *testing.T) {

		mockUsecase := &repoMock.BeerUsecase{}
		mockBeerRepository := &repoMock.BeerRepository{}
		mockCurrencyLayer := &repoMock.CurrencyLayerRepository{}

		testBeerUsecase := usecase.NewBeerUsecase(mockBeerRepository, mockCurrencyLayer)
		mockBeerRepository.On("GetAllBeers", mock.Anything).Return(nil, errors.New("error trace test"))

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		beers, err := testBeerUsecase.GetAllBeers(ctx)
		assert.Error(tt, err)
		assert.Empty(tt, beers)
		mockUsecase.AssertExpectations(tt)
	})

	t.Run("(Not Found) Get All Beers Usecase", func(tt *testing.T) {

		mockUsecase := &repoMock.BeerUsecase{}
		mockBeerRepository := &repoMock.BeerRepository{}
		mockCurrencyLayer := &repoMock.CurrencyLayerRepository{}

		testBeerUsecase := usecase.NewBeerUsecase(mockBeerRepository, mockCurrencyLayer)
		mockBeerRepository.On("GetAllBeers", mock.Anything).Return(nil, nil)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		beers, err := testBeerUsecase.GetAllBeers(ctx)
		assert.Error(tt, err)
		assert.Empty(tt, beers)
		mockUsecase.AssertExpectations(tt)
	})

	t.Run("Get All Beers Usecase", func(tt *testing.T) {

		mockUsecase := &repoMock.BeerUsecase{}
		mockBeerRepository := &repoMock.BeerRepository{}
		mockCurrencyLayer := &repoMock.CurrencyLayerRepository{}

		testBeerUsecase := usecase.NewBeerUsecase(mockBeerRepository, mockCurrencyLayer)
		mockBeerRepository.On("GetAllBeers", mock.Anything).Return(dataBeers(), nil)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		beers, err := testBeerUsecase.GetAllBeers(ctx)
		assert.NoError(tt, err)
		assert.NotEmpty(tt, beers)
		mockUsecase.AssertExpectations(tt)
	})
}

func Test_beerUsecase_GetBeerById(t *testing.T) {

	t.Run("Error Get Beer by ID Usecase", func(tt *testing.T) {

		mockUsecase := &repoMock.BeerUsecase{}
		mockBeerRepository := &repoMock.BeerRepository{}
		mockCurrencyLayer := &repoMock.CurrencyLayerRepository{}

		testBeerUsecase := usecase.NewBeerUsecase(mockBeerRepository, mockCurrencyLayer)
		mockBeerRepository.On("GetBeerById", mock.Anything, mock.AnythingOfType("uint")).Return(domain.Beer{}, errors.New("error trace test"))

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		beer, err := testBeerUsecase.GetBeerById(ctx, uint(1))
		assert.Error(tt, err)
		assert.Empty(tt, beer)
		mockUsecase.AssertExpectations(tt)
	})

	t.Run("(Not Found) Get Beer by ID Usecase", func(tt *testing.T) {

		mockUsecase := &repoMock.BeerUsecase{}
		mockBeerRepository := &repoMock.BeerRepository{}
		mockCurrencyLayer := &repoMock.CurrencyLayerRepository{}

		testBeerUsecase := usecase.NewBeerUsecase(mockBeerRepository, mockCurrencyLayer)
		mockBeerRepository.On("GetBeerById", mock.Anything, mock.AnythingOfType("uint")).Return(domain.Beer{}, nil)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		beer, err := testBeerUsecase.GetBeerById(ctx, uint(1))
		assert.Error(tt, err)
		assert.Empty(tt, beer)
		mockUsecase.AssertExpectations(tt)
	})

	t.Run("Get Beer by ID Usecase", func(tt *testing.T) {

		mockUsecase := &repoMock.BeerUsecase{}
		mockBeerRepository := &repoMock.BeerRepository{}
		mockCurrencyLayer := &repoMock.CurrencyLayerRepository{}

		testBeerUsecase := usecase.NewBeerUsecase(mockBeerRepository, mockCurrencyLayer)
		mockBeerRepository.On("GetBeerById", mock.Anything, mock.AnythingOfType("uint")).Return(dataBeers()[0], nil)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		beer, err := testBeerUsecase.GetBeerById(ctx, uint(1))
		assert.NoError(tt, err)
		assert.NotEmpty(tt, beer)
		mockUsecase.AssertExpectations(tt)
	})

}

func Test_beerUsecase_GetOneBoxPrice(t *testing.T) {

	t.Run("(Error Get Beer By Id) Get One Box Price Usecase", func(tt *testing.T) {

		mockUsecase := &repoMock.BeerUsecase{}
		mockBeerRepository := &repoMock.BeerRepository{}
		mockCurrencyLayer := &repoMock.CurrencyLayerRepository{}

		testBeerUsecase := usecase.NewBeerUsecase(mockBeerRepository, mockCurrencyLayer)
		mockBeerRepository.On("GetBeerById", mock.Anything, mock.AnythingOfType("uint")).Return(domain.Beer{}, errors.New("error trace test"))

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		price, err := testBeerUsecase.GetOneBoxPrice(ctx, 1, "USD", 6)
		assert.Error(tt, err)
		assert.Empty(tt, price)
		mockUsecase.AssertExpectations(tt)
	})

	t.Run("(Not Found) Get One Box Price Usecase", func(tt *testing.T) {

		mockUsecase := &repoMock.BeerUsecase{}
		mockBeerRepository := &repoMock.BeerRepository{}
		mockCurrencyLayer := &repoMock.CurrencyLayerRepository{}

		testBeerUsecase := usecase.NewBeerUsecase(mockBeerRepository, mockCurrencyLayer)
		mockBeerRepository.On("GetBeerById", mock.Anything, mock.AnythingOfType("uint")).Return(domain.Beer{}, nil)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		price, err := testBeerUsecase.GetOneBoxPrice(ctx, 1, "USD", 6)
		assert.Error(tt, err)
		assert.Empty(tt, price)
		mockUsecase.AssertExpectations(tt)
	})

	t.Run("(Error Get Currency) Get One Box Price Usecase", func(tt *testing.T) {

		mockUsecase := &repoMock.BeerUsecase{}
		mockBeerRepository := &repoMock.BeerRepository{}
		mockCurrencyLayer := &repoMock.CurrencyLayerRepository{}

		testBeerUsecase := usecase.NewBeerUsecase(mockBeerRepository, mockCurrencyLayer)
		mockBeerRepository.On("GetBeerById", mock.Anything, mock.AnythingOfType("uint")).Return(dataBeers()[0], nil)
		mockCurrencyLayer.On("GetCurrency", mock.Anything, mock.Anything).Return([]float64{0, 0}, errors.New("error get currency")).Once()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		price, err := testBeerUsecase.GetOneBoxPrice(ctx, 1, "USD", 6)
		assert.Error(tt, err)
		assert.Empty(tt, price)
		mockUsecase.AssertExpectations(tt)
	})

	t.Run("Get One Box Price Usecase", func(tt *testing.T) {

		mockUsecase := &repoMock.BeerUsecase{}
		mockBeerRepository := &repoMock.BeerRepository{}
		mockCurrencyLayer := &repoMock.CurrencyLayerRepository{}

		testBeerUsecase := usecase.NewBeerUsecase(mockBeerRepository, mockCurrencyLayer)
		mockBeerRepository.On("GetBeerById", mock.Anything, mock.AnythingOfType("uint")).Return(dataBeers()[0], nil)
		mockCurrencyLayer.On("GetCurrency", mock.Anything, mock.Anything).Return([]float64{1, 2}, nil).Once()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		price, err := testBeerUsecase.GetOneBoxPrice(ctx, 1, "USD", 6)
		assert.NoError(tt, err)
		assert.NotEmpty(tt, price)
		mockUsecase.AssertExpectations(tt)
	})

}
