package domain

import (
	"context"
	"time"
)

// Beer Data of Beer
// swagger:domain
type Beer struct {
	ID uint `json:"id,omitempty"`
	// Required: true
	Name string `json:"name,omitempty"`
	// Required: true
	Brewery string `json:"brewery,omitempty"`
	// Required: true
	Country string `json:"country,omitempty"`
	// Required: true
	Price float64 `json:"price,omitempty"`
	// Required: true
	Currency  string    `json:"currency,omitempty"`
	CreatedAt time.Time `json:"-"`
}

type BeerUsecase interface {
	GetAllBeers(ctx context.Context) ([]Beer, error)
	GetBeerById(ctx context.Context, id uint) (Beer, error)
	CreateBeerWithId(ctx context.Context, beers *Beer) error
	GetOneBoxPrice(ctx context.Context, id uint, currency string, quantity int) (float64, error)
}

type BeerRepository interface {
	GetAllBeers(ctx context.Context) ([]Beer, error)
	GetBeerById(ctx context.Context, id uint) (Beer, error)
	CreateBeerWithId(ctx context.Context, beers *Beer) error
}

// SwaggerBeerID struct for swagger
// swagger:parameters idBeerPath
type SwaggerBeerID struct {
	// in: path
	// Required: true
	BeerID string `json:"beerID"`
}

// SwaggerBeerBoxPrice struct for swagger
// swagger:parameters idBeerBoxPricePath
type SwaggerBeerBoxPrice struct {
	// in: path
	// Required: true
	BeerID string `json:"beerID"`

	// in: query
	// Required: true
	Currency string `json:"currency"`

	// in: query
	Quantity string `json:"quantity"`
}

// SwaggerBeersRequest Information from Beer
// swagger:parameters beersRequest
type SwaggerBeersRequest struct {
	// in: body
	Body Beer
}

// SwaggerAllBeersResponse Beer It is the response of the all beers information
// swagger:response SwaggerAllBeersResponse
type SwaggerAllBeersResponse struct {
	// in: body
	Body []struct {
		ID       uint    `json:"id,omitempty"`
		Name     string  `json:"name,omitempty"`
		Brewery  string  `json:"brewery,omitempty"`
		Country  string  `json:"country,omitempty"`
		Price    float64 `json:"price,omitempty"`
		Currency string  `json:"currency,omitempty"`
	}
}

// SwaggerBeersResponse Beer It is the response of the beer information
// swagger:response SwaggerBeersResponse
type SwaggerBeersResponse struct {
	// in: body
	Body struct {
		ID       uint    `json:"id,omitempty"`
		Name     string  `json:"name,omitempty"`
		Brewery  string  `json:"brewery,omitempty"`
		Country  string  `json:"country,omitempty"`
		Price    float64 `json:"price,omitempty"`
		Currency string  `json:"currency,omitempty"`
	}
}

// Validate method to perform field validations
func (b *Beer) Validate() map[string]string {
	var errorMessages = make(map[string]string)

	if b.ID == 0 {
		errorMessages["id_required"] = "Id is required or id invalid"
	}

	if b.Name == "" {
		errorMessages["name_required"] = "names is required"
	}

	if b.Brewery == "" {
		errorMessages["brewery_required"] = "brewery is required"
	}

	if b.Country == "" {
		errorMessages["country_required"] = "country is required"
	}

	if b.Price == 0 {
		errorMessages["price_password"] = "price is required and different of zero"
	}

	if b.Currency == "" || len(b.Currency) < 3 {
		errorMessages["currency_required"] = "currency is required and it has to be a valid currency"
	}

	return errorMessages
}
