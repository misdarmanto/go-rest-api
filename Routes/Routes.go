package Routes

import (
	"go-rest-api/Controllers"
	"go-rest-api/Middlewares"

	"github.com/gin-gonic/gin"
)

func UserRouter() *gin.Engine {
	route := gin.Default()

	route.GET("/", Controllers.Index)
	route.POST("/user/login", Controllers.Login)
	route.POST("/user/register", Controllers.Register)
	route.PUT("/user/:userID", Middlewares.Authorization, Controllers.UpdateUser)
	route.DELETE("/user/:userID", Middlewares.Authorization, Controllers.DeleteUser)

	route.POST("/photos", Middlewares.Authorization, Controllers.CreatePhoto)
	route.GET("/photos/:photosID", Middlewares.Authorization, Controllers.GetPhotoByID)
	route.PUT("/photos/:photosID", Middlewares.Authorization, Controllers.UpdatePhoto)
	route.DELETE("/photos/:photosID", Middlewares.Authorization, Controllers.DeletePhoto)

	return route
}
