package scrapers

import (
	"net/http"
	"scrapper/src/constants"
	"scrapper/src/models"
)

func GetScraperByURL(url string) models.Scraper {
	req, _ := http.Get(url)
	host := req.Request.URL.Host
	switch host {
	case constants.EXITO_HOST:
		return ExitoScraper{}
	case constants.JUMBO_HOST:
		return JumboScraper{}
	case constants.AMAZON_HOST:
		return ExitoScraper{}
	default:
		return nil
	}
}
