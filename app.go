package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/restfulapi/models"
	"github.com/restfulapi/routes"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(static.Serve("/view", static.LocalFile("client", true)))

	db := models.GetDB()
	models.Migrate(db)

	apiroute := r.Group("/api")

	routes.ApiRoute(apiroute, db)

	r.Run(":8080")
	db.Close()
}
