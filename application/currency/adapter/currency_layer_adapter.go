package adapter

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/samuskitchen/beer-api-clean-arch/domain"
)

type currencyAdapter struct {
	currencyClient *http.Client
}

func NewCurrencyAdapter(Connection *http.Client) domain.CurrencyLayerRepository {
	return &currencyAdapter{
		currencyClient: Connection,
	}
}

func (c *currencyAdapter) GetCurrency(currencyPay, currencyBeer string) ([]float64, error) {
	var valueEmpty []float64
	accessKey := os.Getenv("ACCESS_KEY_CURRENCY")

	responseCurrency, err := c.currencyClient.Get(fmt.Sprintf("http://apilayer.net/api/live?access_key=%s&currencies=%s,%s&source=USD&format=1", accessKey, currencyPay, currencyBeer))
	if err != nil {
		return valueEmpty, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("panic occurred:", err)
		}
	}(responseCurrency.Body)

	if responseCurrency.StatusCode != 200 {
		return valueEmpty, fmt.Errorf("status code error: %d %s", responseCurrency.StatusCode, responseCurrency.Status)
	}

	responseData, err := ioutil.ReadAll(responseCurrency.Body)
	if err != nil {
		return valueEmpty, err
	}

	var currencyLayer domain.CurrencyLayer
	err = json.Unmarshal(responseData, &currencyLayer)
	if err != nil {
		return valueEmpty, err
	}

	values := make([]float64, 0)

	valueCurrencyPay, ok := currencyLayer.Quotes["USD"+currencyPay].(float64)
	if !ok {
		return valueEmpty, errors.New("error get currency to pay")
	}

	values = append(values, valueCurrencyPay)

	valueCurrencyBeer, ok := currencyLayer.Quotes["USD"+currencyBeer].(float64)
	if !ok {
		return valueEmpty, errors.New("error get currency of the beer")
	}

	values = append(values, valueCurrencyBeer)

	return values, err
}
