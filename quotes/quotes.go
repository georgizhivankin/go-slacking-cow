package quotes

import (
	"encoding/json"
	"fmt"
	"github.com/georgizhivankin/go-slacking-cow/config"
	"github.com/georgizhivankin/go-slacking-cow/helpers"
	"github.com/georgizhivankin/go-slacking-cow/models"
	"net/http"
	"time"
)

func GetRandomQuote() *models.Quote {
	quoteStruct := new(models.Quote)
	var httpClient = &http.Client{Timeout: 10 * time.Second}
	config, err := config.LoadEnv()
	helpers.CheckAndLogError(err, "panic")

	requestUrl := fmt.Sprintf("%s?method=%s&key=%d&format=%s&lang=%s", config.QuotesApi.ApiUrl, config.QuotesApi.ApiMethod, config.QuotesApi.ApiKey, config.QuotesApi.ApiFormat, config.QuotesApi.ApiLanguage)
	helpers.CheckAndLogError(err, "panic")

	req, err := httpClient.Get(requestUrl)

	helpers.CheckAndLogError(err, "panic")

	json.NewDecoder(req.Body).Decode(quoteStruct)

	return quoteStruct
}
