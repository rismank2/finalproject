package controllers

import (
	"finalproject/database"
	"finalproject/helpers"
	"finalproject/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var appJSON = "application/json"

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	User := models.User{}

	if appJSON == helpers.GetContentType(c) {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"age":      User.Age,
		"email":    User.Email,
		"id":       User.ID,
		"username": User.Username,
	})
}

func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := c.Request.Header.Get("Content_Type")
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid Email/Password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid Email/Password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func UpdateUser(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	User := models.User{}

	paramId, _ := strconv.Atoi(c.Param("userId"))
	userId := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	User.ID = userId

	err := db.Model(&User).Where("id = ?", paramId).Updates(models.User{Username: User.Username, Email: User.Email, Password: User.Password, Age: User.Age}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         User.ID,
		"email":      User.Email,
		"username":   User.Username,
		"age":        User.Age,
		"updated_at": User.UpdatedAt,
	})
}

func DeleteUser(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	User := models.User{}

	userID := uint(userData["id"].(float64))

	User.ID = userID

	err := db.Model(&User).Where("id = ?", userID).Delete(models.User{}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})

}
