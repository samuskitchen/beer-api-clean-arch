package adapter

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/samuskitchen/beer-api-clean-arch/application/currency/adapter"
	repoMock "github.com/samuskitchen/beer-api-clean-arch/domain/mocks"
	"github.com/stretchr/testify/assert"
)

// roundTripFunc .
type roundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// httpClientWithRoundTripper returns *http.Client with Transport replaced to avoid making real calls
func httpClientWithRoundTripper(statusCode int, response string) *http.Client {
	return &http.Client{
		Transport: roundTripFunc(func(req *http.Request) *http.Response {
			return &http.Response{
				StatusCode: statusCode,
				Body:       ioutil.NopCloser(bytes.NewBufferString(response)),
			}
		}),
	}
}

func Test_currencyClientData_GetCurrency(t *testing.T) {

	t.Run("Error status code", func(tt *testing.T) {

		mockCurrencyLayer := &repoMock.CurrencyLayerRepository{}

		client := httpClientWithRoundTripper(http.StatusBadRequest, "")

		testCurrencyAdapter := adapter.NewCurrencyAdapter(client)

		result, err := testCurrencyAdapter.GetCurrency("USD", "COP")
		assert.Error(tt, err)
		assert.Empty(tt, result)
		mockCurrencyLayer.AssertExpectations(tt)
	})

	t.Run("Error Unmarshal", func(tt *testing.T) {

		mockCurrencyLayer := &repoMock.CurrencyLayerRepository{}

		client := httpClientWithRoundTripper(http.StatusOK, "Ok")

		testCurrencyAdapter := adapter.NewCurrencyAdapter(client)

		result, err := testCurrencyAdapter.GetCurrency("USD", "COP")
		assert.Error(tt, err)
		assert.Empty(tt, result)
		mockCurrencyLayer.AssertExpectations(tt)
	})

	t.Run("Error get currency to pay", func(tt *testing.T) {

		mockCurrencyLayer := &repoMock.CurrencyLayerRepository{}

		client := httpClientWithRoundTripper(http.StatusOK, `
					{
						"success": true,
						"terms": "https://currencylayer.com/terms",
						"privacy": "https://currencylayer.com/privacy",
						"timestamp": 1620958384,
						"source": "USD",
						"quotes": {
							"USDCOP": 3721.11
						}
					}`)

		testCurrencyAdapter := adapter.NewCurrencyAdapter(client)

		result, err := testCurrencyAdapter.GetCurrency("USD", "COP")
		assert.Error(tt, err)
		assert.Empty(tt, result)
		mockCurrencyLayer.AssertExpectations(tt)
	})

	t.Run("Error get currency of the beer", func(tt *testing.T) {

		mockCurrencyLayer := &repoMock.CurrencyLayerRepository{}

		client := httpClientWithRoundTripper(http.StatusOK, `
					{
						"success": true,
						"terms": "https://currencylayer.com/terms",
						"privacy": "https://currencylayer.com/privacy",
						"timestamp": 1620958384,
						"source": "USD",
						"quotes": {
							"USDUSD": 1
						}
					}`)

		testCurrencyAdapter := adapter.NewCurrencyAdapter(client)

		result, err := testCurrencyAdapter.GetCurrency("USD", "COP")
		assert.Error(tt, err)
		assert.Empty(tt, result)
		mockCurrencyLayer.AssertExpectations(tt)
	})

	t.Run("Get Currency Successful", func(tt *testing.T) {

		mockCurrencyLayer := &repoMock.CurrencyLayerRepository{}

		client := httpClientWithRoundTripper(http.StatusOK, `
					{
						"success": true,
						"terms": "https://currencylayer.com/terms",
						"privacy": "https://currencylayer.com/privacy",
						"timestamp": 1620958384,
						"source": "USD",
						"quotes": {
							"USDCOP": 3721.11,
							"USDUSD": 1
						}
					}`)

		testCurrencyAdapter := adapter.NewCurrencyAdapter(client)

		result, err := testCurrencyAdapter.GetCurrency("USD", "COP")
		assert.NoError(tt, err)
		assert.NotEmpty(tt, result)
		mockCurrencyLayer.AssertExpectations(tt)
	})
}
