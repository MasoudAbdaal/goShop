package routes

import "github.com/gin-gonic/gin"

func RegisterUserRoutes(group *gin.RouterGroup) {

	group.GET("/:userId", getUserById)

}

func getUserById(c *gin.Context) {
	c.Redirect(305, "/hello")
}
