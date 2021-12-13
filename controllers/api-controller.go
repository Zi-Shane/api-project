package controllers

import (
	"api/database"
	"api/services"
	"net/http"

	"github.com/gorilla/mux"
)

// HandlerFunc of `GET /api/language/{id}`
func GetLanguage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)   // 獲取url參數
	queryId := vars["id"] // 獲取{id}

	// Call function from `services` to get data
	var languages []database.Languages = services.ReadLanguage(queryId)

	// Response data to Client
	response := ApiResponse{"200", languages}
	ResponseWithJson(w, http.StatusOK, response)
}
