package usecase

import (
	"context"
	"errors"
	"math"
	"time"

	"github.com/samuskitchen/beer-api-clean-arch/domain"
)

type beerUsecase struct {
	beerRepo     domain.BeerRepository
	currencyRepo domain.CurrencyLayerRepository
}

// NewBeerUsecase constructor
func NewBeerUsecase(b domain.BeerRepository, c domain.CurrencyLayerRepository) domain.BeerUsecase {
	return &beerUsecase{
		beerRepo:     b,
		currencyRepo: c,
	}
}

func (b *beerUsecase) GetAllBeers(ctx context.Context) ([]domain.Beer, error) {

	beers, err := b.beerRepo.GetAllBeers(ctx)
	if err != nil {
		return nil, err
	}

	if beers == nil {
		return nil, errors.New("beers not found")
	}

	return beers, nil
}

func (b *beerUsecase) GetBeerById(ctx context.Context, id uint) (domain.Beer, error) {
	beerResult, err := b.beerRepo.GetBeerById(ctx, id)
	if err != nil {
		return domain.Beer{}, err
	}

	if (domain.Beer{}) == beerResult {
		return domain.Beer{}, errors.New("the beer ID does not exist")
	}

	return beerResult, nil
}

func (b *beerUsecase) CreateBeerWithId(ctx context.Context, beers *domain.Beer) error {
	beers.CreatedAt = time.Now()
	err := b.beerRepo.CreateBeerWithId(ctx, beers)

	if err != nil {
		return err
	}

	return nil
}

// GetOneBoxPrice Method that returns the currency of payment and beer as the base currency of conversion to the dollar,
// the order of the returned values are: currencyPay, currencyBeer
func (b *beerUsecase) GetOneBoxPrice(ctx context.Context, id uint, currency string, quantity int) (float64, error) {

	beerResult, err := b.beerRepo.GetBeerById(ctx, id)
	if err != nil {
		return 0, err
	}

	if (domain.Beer{}) == beerResult {
		return 0, errors.New("the beer ID does not exist")
	}

	valueCurrency, err := b.currencyRepo.GetCurrency(currency, beerResult.Currency)
	if err != nil {
		return 0, err
	}

	valueTotalBeer := beerResult.Price * float64(quantity)
	total := valueCurrency[0] / valueCurrency[1] * valueTotalBeer
	totalFloat := math.Round(total*100) / 100

	return totalFloat, nil
}
