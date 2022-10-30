package Controllers

import (
	"go-rest-api/Config"
	"go-rest-api/Models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CreatePhoto(c *gin.Context) {
	var body Models.PhotoModel
	c.BindJSON(&body)
	photo := Models.PhotoModel{Title: body.Title, Caption: body.Caption, PhotoUrl: body.PhotoUrl, UserID: body.UserID}
	result := Config.DB.Create(&photo)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func GetPhotoByID(c *gin.Context) {
	var photo Models.PhotoModel
	if err := Config.DB.First(&photo, "id = ?", c.Param("photosID")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": photo})
}

func UpdatePhoto(c *gin.Context) {
	var photo Models.PhotoModel
	if err := Config.DB.Where("id = ?", c.Param("photosID")).First(&photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input Models.PhotoModel
	if err := c.ShouldBindBodyWith(&input, binding.JSON); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	Config.DB.Model(&photo).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": photo})
}

func DeletePhoto(c *gin.Context) {
	var photo Models.PhotoModel
	if err := Config.DB.Delete(&photo, c.Param("photosID")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "photo deleted"})
}
