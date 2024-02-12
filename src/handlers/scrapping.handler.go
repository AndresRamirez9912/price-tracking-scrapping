package handlers

import (
	"encoding/json"
	"net/http"
	"scrapper/src/constants"
	"scrapper/src/models"
	"scrapper/src/services"
	"scrapper/src/utils"
)

func ScrappingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(constants.CONTENT_TYPE, constants.APPICATION_JSON)
	requestResponse := &models.ScrapProductResponse{}

	// Get the product's URL
	url, err := utils.GetProductURL(r)
	if err != nil {
		errorResponse := utils.SendScrapedErrorResponse(requestResponse, err, 0)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorResponse)
		return
	}

	// Get the information product information
	result, err := services.ScrapProduct(url)
	if err != nil {
		errorResponse := utils.SendScrapedErrorResponse(requestResponse, err, 0)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(errorResponse)
		return
	}

	// Build the response
	requestResponse.ScrapedProduct = *result
	requestResponse.Success = true

	jsonData, err := json.Marshal(requestResponse)
	if err != nil {
		errorResponse := utils.SendScrapedErrorResponse(requestResponse, err, 0)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(errorResponse)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write(jsonData)
}
