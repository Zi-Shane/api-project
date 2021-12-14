package routes

import (
	"api/controllers"
)

// Configure of API Route
func init() {
	register("GET", "/api/hostname", controllers.GetHostname)
	register("GET", "/api/language/:id", controllers.GetLanguage)
	register("POST", "/api/language", controllers.AddLanguage)
	// register("UPDATE", "/api/language", controllers.UpdateLanguage)
	register("DELETE", "/api/language/:id", controllers.RemoveLanguage)
}
