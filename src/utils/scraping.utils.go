package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"scrapper/src/models"
	"strings"
)

func GetProductURL(r *http.Request) (string, error) {
	productToScraped := &models.ScrapProductRequest{}
	err := json.NewDecoder(r.Body).Decode(productToScraped)
	if err != nil {
		log.Println("Error decoding the product URL", err)
		return "", err
	}
	return productToScraped.URL, nil
}

func GetBrandIdByPLU(PLU string) (brand, id string) {
	result := strings.Split(PLU, "<!-- -->-PLU: <!-- -->")
	brand = result[0]
	id = result[1]
	return
}

func SendScrapedErrorResponse(requestResponse *models.ScrapProductResponse, err error, errorCode int) []byte {
	// Create the Response object
	requestResponse.ScrapedProduct = models.ScrapedProduct{}
	requestResponse.ErrorCode = errorCode
	requestResponse.ErrorMessage = err.Error()
	requestResponse.Success = false

	// Serialize the Object
	jsonData, _ := json.Marshal(requestResponse)

	return jsonData
}

func PrintResult(htmlContent string) {
	// Print Text with the hole HTML
	file, _ := os.Create("response.html")
	_, err := file.WriteString(htmlContent)
	if err != nil {
		log.Println("Error creating HTML document")
		return
	}
	file.Close()
}
