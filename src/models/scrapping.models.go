package models

type scraper interface {
	ScrapProduct()
}

type ScrapedProduct struct {
	Name                   string `json:"name"`
	Id                     string `json:"id"`
	Brand                  string `json:"brand"`
	HigherPrice            string `json:"higherPrice"`
	LowePrice              string `json:"lowePrice"`
	OtherPaymentLowerPrice string `json:"otherPaymentLowerPrice"`
	Discount               string `json:"discount"`
	ImageURL               string `json:"imageURL"`
	Store                  string `json:"store"`
	ProductURL             string `json:"productURL"`
}

type ScrapProductRequest struct {
	URL string `json:"url"`
}

type ScrapProductResponse struct {
	ScrapedProduct ScrapedProduct `json:"scrapedProduct"`
	Success        bool           `json:"success"`
	ErrorCode      int            `json:"errorCode"`
	ErrorMessage   string         `json:"errorMessage"`
}
