package routes

import (
	"api/controllers"
)

func init() {
	register("GET", "/api/language/{id}", controllers.GetLanguage, nil)
}
