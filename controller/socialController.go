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

func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	SocialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID

	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, SocialMedia)
}

func GetSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	SocialMedia := models.SocialMedia{}
	User := models.User{}

	userID := uint(userData["id"].(float64))

	SocialMedia.UserID = userID

	err := db.Where("user_id = ?", userID).Find(&SocialMedia).Error
	errUser := db.Where("id = ?", userID).Find(&User).Error

	if err != nil || errUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if SocialMedia.ID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"social_media": gin.H{
				"id":               SocialMedia.ID,
				"name":             SocialMedia.Name,
				"social_media_url": SocialMedia.SocialMediaURL,
				"user_id":          SocialMedia.UserID,
				"created_at":       SocialMedia.CreatedAt,
				"updated_at":       SocialMedia.UpdatedAt,
				"User": gin.H{
					"id":       User.ID,
					"username": User.Username,
				},
			},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "You don't have any data",
		})
	}
}

func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	SocialMedia := models.SocialMedia{}

	socialMediaID, _ := strconv.Atoi(c.Param("socialMediaId"))
	userId := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userId
	SocialMedia.ID = uint(socialMediaID)

	err := db.Model(&SocialMedia).Where("id = ?", socialMediaID).Updates(models.SocialMedia{Name: SocialMedia.Name, SocialMediaURL: SocialMedia.SocialMediaURL}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               SocialMedia.ID,
		"name":             SocialMedia.Name,
		"social_media_url": SocialMedia.SocialMediaURL,
		"user_id":          SocialMedia.UserID,
		"updated_at":       SocialMedia.UpdatedAt,
	})
}

func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	SocialMedia := models.SocialMedia{}

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	userId := uint(userData["id"].(float64))

	SocialMedia.ID = uint(socialMediaId)
	SocialMedia.UserID = userId

	err := db.Model(&SocialMedia).Where("id = ?", socialMediaId).Delete(models.SocialMedia{}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
