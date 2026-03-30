package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"micro-go/model"
	"micro-go/pkg/logger"

	"go.uber.org/zap"
)

// UserHandler 用户 API 处理器
type UserHandler struct {
}

// NewUserHandler 创建新的用户 API 处理器
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// CreateUser 创建用户
// @Summary 创建用户
// @Description 创建新用户
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.UserRequest true "用户信息"
// @Success 200 {object} model.Response{data=model.User}
// @Failure 400 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req model.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Invalid request body",
			zap.Error(err),
		)
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, "Invalid request body"))
		return
	}
	
	logger.Info("User created successfully",
		zap.String("name", req.Name),
		zap.String("email", req.Email),
		zap.Int("age", req.Age),
	)

	user := &model.User{
		ID:    1,
		Name:  req.Name,
		Email: req.Email,
		Age:   req.Age,
	}

	c.JSON(http.StatusOK, model.SuccessResponse(user))
}

// GetUserByID 根据ID获取用户
// @Summary 根据ID获取用户
// @Description 根据用户ID获取用户信息
// @Tags users
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} model.Response{data=model.User}
// @Failure 400 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /users/{id} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error("Invalid user ID",
			zap.Error(err),
			zap.String("id", idStr),
		)
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, "Invalid user ID"))
		return
	}
	
	logger.Info("User retrieved successfully",
		zap.Int("id", id),
	)

	user := &model.User{
		ID:    id,
		Name:  "John Doe",
		Email: "john@example.com",
		Age:   30,
	}

	c.JSON(http.StatusOK, model.SuccessResponse(user))
}

// GetAllUsers 获取所有用户
// @Summary 获取所有用户
// @Description 获取所有用户信息
// @Tags users
// @Produce json
// @Success 200 {object} model.Response{data=[]model.User}
// @Failure 500 {object} model.Response
// @Router /users [get]
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	logger.Info("Getting all users")
	
	users := []*model.User{
		{
			ID:    1,
			Name:  "John Doe",
			Email: "john@example.com",
			Age:   30,
		},
		{
			ID:    2,
			Name:  "Jane Smith",
			Email: "jane@example.com",
			Age:   25,
		},
	}
	
	logger.Info("Users retrieved successfully",
		zap.Int("count", len(users)),
	)

	c.JSON(http.StatusOK, model.SuccessResponse(users))
}

// UpdateUser 更新用户
// @Summary 更新用户
// @Description 更新用户信息
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Param user body model.UserRequest true "用户信息"
// @Success 200 {object} model.Response{data=model.User}
// @Failure 400 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error("Invalid user ID",
			zap.Error(err),
			zap.String("id", idStr),
		)
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, "Invalid user ID"))
		return
	}

	var req model.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Invalid request body",
			zap.Error(err),
		)
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, "Invalid request body"))
		return
	}
	
	logger.Info("User updated successfully",
		zap.Int("id", id),
		zap.String("name", req.Name),
		zap.String("email", req.Email),
		zap.Int("age", req.Age),
	)

	user := &model.User{
		ID:    id,
		Name:  req.Name,
		Email: req.Email,
		Age:   req.Age,
	}

	c.JSON(http.StatusOK, model.SuccessResponse(user))
}

// DeleteUser 删除用户
// @Summary 删除用户
// @Description 删除用户
// @Tags users
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error("Invalid user ID",
			zap.Error(err),
			zap.String("id", idStr),
		)
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, "Invalid user ID"))
		return
	}
	
	logger.Info("User deleted successfully",
		zap.Int("id", id),
	)

	c.JSON(http.StatusOK, model.SuccessResponse(nil))
}
