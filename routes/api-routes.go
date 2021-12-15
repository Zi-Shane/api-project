package routes

import (
	"api/controllers"
)

// Configure of API Route
func init() {
	setupHostInfoRoute()
	setupLanguageRoute()
}

func setupHostInfoRoute() {
	register("GET", "/api/hostname", controllers.GetHostname)
}

func setupLanguageRoute() {
	register("GET", "/api/language/:id", controllers.GetLanguage)
	register("GET", "/api/languageRange/:start/:end", controllers.GetLanguages)
	register("POST", "/api/language", controllers.AddLanguage)
	register("DELETE", "/api/language/:language", controllers.RemoveLanguage)
	register("PUT", "/api/language", controllers.UpdateLanguage)
}
