package main

import (
	"log"
	"os"

	c "github.com/DaurenCS/Online-Store/controllers"
	d "github.com/DaurenCS/Online-Store/database"
	"github.com/DaurenCS/Online-Store/middleware"
	"github.com/DaurenCS/Online-Store/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := c.NewApplication(d.ProductData(d.Client, "Products"), d.UserData(d.Client, "Users"))
	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())
	router.GET("/addtocart", app.AddToCart())
	router.GET("/remove", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}

//part 3
