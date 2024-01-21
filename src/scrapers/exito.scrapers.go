package scrapers

import (
	"log"
	"scrapper/src/constants"
	"scrapper/src/models"
	"scrapper/src/utils"

	"github.com/PuerkitoBio/goquery"
)

type ExitoScraper struct{}

func (exito ExitoScraper) GetVisibleSelector() string {
	return constants.EXITO_OTHER_PRICE_SELECTOR
}

func (exito ExitoScraper) ScrapInformation(doc *goquery.Document, url string) (*models.ScrapedProduct, error) {
	lowerPrice := doc.Find(constants.EXITO_LOWER_PRICE_SELECTOR).First().Text()
	otherLower, err := doc.Find(constants.EXITO_OTHER_PRICE_SELECTOR).Html()
	higherPrice, _ := doc.Find(constants.EXITO_HIGHER_PRICE_SELECTOR).Html()
	productName, _ := doc.Find(constants.EXITO_PRODUCT_NAME_SELECTOR).Html()
	plu, _ := doc.Find(constants.EXITO_PRODUCT_BRAND_SELECTOR).Html()
	productImage := doc.Find(constants.EXITO_PRODUCT_IMAGE_SELECTOR).AttrOr("src", "")
	disccount, _ := doc.Find(constants.EXITO_PRODUCT_DISCOUNT_SELECTOR).Html()
	brand, id := utils.GetBrandIdByPLU(plu)
	store := "Exito"
	if err != nil {
		log.Println("Error sacraping the information from the Product", err)
		return nil, err
	}

	// Build the Product
	scrapedProduct := &models.ScrapedProduct{
		Name:                   productName,
		Id:                     id,
		Brand:                  brand,
		HigherPrice:            higherPrice,
		LowePrice:              lowerPrice,
		Discount:               disccount,
		ImageURL:               productImage,
		Store:                  store,
		ProductURL:             url,
		OtherPaymentLowerPrice: otherLower,
	}
	return scrapedProduct, nil
}
