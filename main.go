package main

import (
	"api/routes"
)

func main() {
	router := routes.NewRouter()
	// router.TrustedPlatform = gin.PlatformCloudflare
	router.Run(":3000") // listen and serve on 0.0.0.0:8080
}
