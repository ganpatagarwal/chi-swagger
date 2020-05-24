package handlers

import (
	"net/http"

	"github.com/ganpatagarwal/chi-swagger/utils"
)

type response struct {
	Status     string
	StatusCode int
}

// RootHandler - Returns all the available APIs
// @Summary This API can be used as health check for this application.
// @Description Tells if the chi-swagger APIs are working or not.
// @Tags chi-swagger
// @Accept  json
// @Produce  json
// @Success 200 {string} response "api response"
// @Router / [get]
func RootHandler(w http.ResponseWriter, r *http.Request) {

	utils.WriteResponse(w, utils.SuccessResponse, &response{
		Status:     "Working OK",
		StatusCode: 200,
	})
}
