package scrapers

import "scrapper/src/models"

func GetScraperByURL(url string) models.Scraper {
	switch url {
	case "www.exito.com":
		return ExitoScraper{}
	case "www.jumbo.com":
		return ExitoScraper{}
	case "www.amazon.com":
		return ExitoScraper{}
	}
	return nil
}
