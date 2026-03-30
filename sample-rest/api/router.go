package api

import (
	"github.com/gin-gonic/gin"

	"micro-go/pkg/logger"
)

// SetupRouter 设置路由
func SetupRouter(logger *logger.Logger) *gin.Engine {
	r := gin.Default()

	// 创建用户处理器
	userHandler := NewUserHandler(logger)

	// 注册用户路由
	users := r.Group("/users")
	{
		users.POST("", userHandler.CreateUser)
		users.GET("", userHandler.GetAllUsers)
		users.GET("/:id", userHandler.GetUserByID)
		users.PUT("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)
	}

	return r
}
