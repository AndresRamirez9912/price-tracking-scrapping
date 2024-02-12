package scrapers

import (
	"log"
	"scrapper/src/constants"
	"scrapper/src/models"
	"sort"

	"github.com/PuerkitoBio/goquery"
)

type JumboScraper struct{}

func (jumbo JumboScraper) GetVisibleSelector() string {
	return constants.JUMBO_WAIT_UNTIL
}

func (jumbo JumboScraper) ScrapInformation(doc *goquery.Document, url string) (*models.ScrapedProduct, error) {
	prices := [3]string{"nil", "nil", "nil"}
	doc.Find(constants.JUMBO_PRICES_SELECTOR).Each(func(i int, s *goquery.Selection) {
		if s.Text() != "" {
			prices[i] = s.Text()
		}
	})
	sort.Strings(prices[:])
	productName, err := doc.Find(constants.JUMBO_PRODUCT_NAME_SELECTOR).Html()
	brand, _ := doc.Find(constants.JUMBO_PRODUCT_BRAND_SELECTOR).Html()
	productImage := doc.Find(constants.JUMBO_PRODUCT_IMAGE_SELECTOR).AttrOr("src", "")
	disccount, _ := doc.Find(constants.JUMBO_PRODUCT_DISCOUNT_SELECTOR).Html()
	id, _ := doc.Find(constants.JUMBO_PRODUCT_ID_SELECTOR).Html()
	store := constants.JUMBO_HOST
	if err != nil {
		log.Println("Error sacraping the information from the Product", err)
		return nil, err
	}

	// Build the Product
	scrapedProduct := &models.ScrapedProduct{
		Name:                   productName,
		Id:                     id,
		Brand:                  brand,
		HigherPrice:            prices[2],
		LowePrice:              prices[0],
		Discount:               disccount,
		ImageURL:               productImage,
		Store:                  store,
		ProductURL:             url,
		OtherPaymentLowerPrice: prices[1],
	}
	return scrapedProduct, nil
}
