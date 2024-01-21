package services

import (
	"context"
	"log"
	"scrapper/src/constants"
	"scrapper/src/models"
	"scrapper/src/utils"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

func ScrapProduct(url string) (*models.ScrapedProduct, error) {
	// Create the context
	ctxTimeOut, cancelTime := context.WithTimeout(context.Background(), 60*time.Second)
	ctx, cancel := chromedp.NewContext(ctxTimeOut)
	defer func() {
		cancel()
		cancelTime()
	}()

	// Catch the page after the JS selector loaded
	var htmlContent string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(constants.EXITO_OTHER_PRICE_SELECTOR),
		chromedp.OuterHTML("html", &htmlContent),
	)
	if err != nil {
		log.Println("Error scraping HTML element", err)
		return nil, err
	}

	// Create the new goquery Document from the HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Println("Error creating the Goquery document from the HTML content", err)
		return nil, err
	}

	// Scrap the information
	lowerPrice := doc.Find(constants.EXITO_LOWER_PRICE_SELECTOR).First().Text()
	if err != nil {
		log.Println("Error sacraping the information from the Product", err)
		return nil, err
	}
	otherLower, err := doc.Find(constants.EXITO_OTHER_PRICE_SELECTOR).Html()
	higherPrice, _ := doc.Find(constants.EXITO_HIGHER_PRICE_SELECTOR).Html()
	productName, _ := doc.Find(constants.EXITO_PRODUCT_NAME_SELECTOR).Html()
	plu, _ := doc.Find(constants.EXITO_PRODUCT_BRAND_SELECTOR).Html()
	productImage := doc.Find(constants.EXITO_PRODUCT_IMAGE_SELECTOR).AttrOr("src", "")
	disccount, _ := doc.Find(constants.EXITO_PRODUCT_DISCOUNT_SELECTOR).Html()
	brand, id := utils.GetBrandIdByPLU(plu)
	store := "Exito"

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
