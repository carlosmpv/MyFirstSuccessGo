package routes

import (
	"crypto/md5"
	"encoding/hex"
	"html"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/restfulapi/models"
)

func ApiRoute(r *gin.RouterGroup, db *gorm.DB) {
	r.GET("/users", func(c *gin.Context) {
		var users []models.User
		db.Find(&users)
		c.JSON(200, users)
	})

	r.POST("/login", func(c *gin.Context) {
		var json models.UserLogin
		var hash = md5.New()
		if err := c.ShouldBindJSON(&json); err == nil {
			hash.Write([]byte(json.Password))
			email := html.EscapeString(json.Email)
			pass := hex.EncodeToString(hash.Sum(nil))
			var user models.User

			db.Where("email = ? AND pass_hash = ?", email, pass).First(&user)
			c.JSON(200, user)
		}
	})

	r.POST("/register", func(c *gin.Context) {
		var json models.UserRegister
		var hash = md5.New()
		if err := c.ShouldBindJSON(&json); err == nil {
			var exist models.User
			db.Where("email = ?", json.Email).First(&exist)
			if exist.ID != 0 {
				c.JSON(200, gin.H{"Error": "User already registered"})
				return
			}

			hash.Write([]byte(json.Password))
			user := &models.User{
				Name:     html.EscapeString(json.Name),
				Email:    html.EscapeString(json.Email),
				PassHash: hex.EncodeToString(hash.Sum(nil)),
			}
			db.Create(&user)
			c.JSON(200, json)
		} else {
			c.JSON(200, gin.H{"Error": err})
		}
	})
}
