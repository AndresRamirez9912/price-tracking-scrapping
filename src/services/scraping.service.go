package services

import (
	"context"
	"log"
	"scrapper/src/models"
	"scrapper/src/scrapers"
	"scrapper/src/utils"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

func ScrapProduct(url string) (*models.ScrapedProduct, error) {
	// Create Scrapper
	scraper := scrapers.GetScraperByURL(url)

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
		chromedp.WaitVisible(scraper.GetVisibleSelector()),
		chromedp.OuterHTML("html", &htmlContent),
	)
	if err != nil {
		log.Println("Error scraping HTML element", err)
		return nil, err
	}

	// Generate HTML
	utils.PrintResult(htmlContent)

	// Create the new goquery Document from the HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Println("Error creating the Goquery document from the HTML content", err)
		return nil, err
	}

	// Scrap the information
	scrapedProduct, err := scraper.ScrapInformation(doc, url)
	if err != nil {
		return nil, err
	}

	return scrapedProduct, nil
}
