package app

import "github.com/coderile/shortenurl-go/controller"

func mapUrl() {
	router.POST("/short_url", controller.CrateShortUrl)
	router.GET("/:shortid", controller.RedirectOriginalURL)
}
