package Controllers

import (
	"go-rest-api/Config"
	"go-rest-api/Models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, "Welcome")
}

func Register(c *gin.Context) {
	var body Models.UserModel
	c.BindJSON(&body)
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	user := Models.UserModel{Email: body.Email, UserName: body.UserName, Password: string(hash)}
	result := Config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func Login(c *gin.Context) {
	var body Models.UserModel
	if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	var user Models.UserModel
	err := Config.DB.Where("user_name= ?", body.UserName).First(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "username not found",
		})
		return
	}

	isError := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if isError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email and password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.Id,
		"exp":    time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}

func UpdateUser(c *gin.Context) {
	var user Models.UserModel
	if err := Config.DB.Where("id = ?", c.Param("userID")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input Models.UserModel
	if err := c.ShouldBindBodyWith(&input, binding.JSON); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	Config.DB.Model(&user).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
	var user Models.UserModel
	if err := Config.DB.Delete(&user, c.Param("userID")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}
