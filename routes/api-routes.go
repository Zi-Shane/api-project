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
	register("GET", "/api/GetLanguage/:id", controllers.GetLanguage)
	register("GET", "/api/GetLanguageRange/:start/:end", controllers.GetLanguages)
	register("GET", "/api/GetCountryUselanguage", controllers.GetCountryUesdLanguages)
	register("POST", "/api/AddLanguage", controllers.AddLanguage)
	register("DELETE", "/api/RemoveLanguage/:language", controllers.RemoveLanguage)
	register("PUT", "/api/UpdateLanguage", controllers.UpdateLanguage)
}
