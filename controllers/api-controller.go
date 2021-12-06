package controllers

import (
	"api/database"
	"api/services"
	"net/http"

	"github.com/gorilla/mux"
)

func GetLanguage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	queryId := vars["id"] //獲取url參數

	// intQueryId, _ := strconv.Atoi(queryId)
	var languages []database.Languages = services.ReadLanguage(queryId)

	response := ApiResponse{"200", languages}
	services.ResponseWithJson(w, http.StatusOK, response)
}
