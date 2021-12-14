package routes

import (
	"api/controllers"
)

// Configure of API Route
func init() {
	register("GET", "/api/language/:id", controllers.GetLanguage)
	register("GET", "/api/hostname", controllers.GetHostname)
}
