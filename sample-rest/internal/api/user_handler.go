package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"micro-go/internal/model"
	"micro-go/internal/mock"
	"micro-go/pkg/logger"
)

// UserHandler 用户 API 处理器
type UserHandler struct {
	userMock *mock.UserMock
	logger   *logger.Logger
}

// NewUserHandler 创建新的用户 API 处理器
func NewUserHandler(userMock *mock.UserMock, logger *logger.Logger) *UserHandler {
	return &UserHandler{
		userMock: userMock,
		logger:   logger,
	}
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
		h.logger.Error("Invalid request body", map[string]interface{}{"error": err.Error()})
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, "Invalid request body"))
		return
	}

	user := &model.User{
		Name:  req.Name,
		Email: req.Email,
		Age:   req.Age,
	}

	err := h.userMock.Create(user)
	if err != nil {
		h.logger.Error("Failed to create user", map[string]interface{}{"error": err.Error()})
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, "Failed to create user"))
		return
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
		h.logger.Error("Invalid user ID", map[string]interface{}{"error": err.Error()})
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, "Invalid user ID"))
		return
	}

	user, err := h.userMock.GetByID(id)
	if err != nil {
		h.logger.Error("Failed to get user", map[string]interface{}{"error": err.Error(), "user_id": id})
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, "Failed to get user"))
		return
	}

	if user == nil {
		h.logger.Error("User not found", map[string]interface{}{"user_id": id})
		c.JSON(http.StatusNotFound, model.ErrorResponse(http.StatusNotFound, "User not found"))
		return
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
	users, err := h.userMock.GetAll()
	if err != nil {
		h.logger.Error("Failed to get users", map[string]interface{}{"error": err.Error()})
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, "Failed to get users"))
		return
	}

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
		h.logger.Error("Invalid user ID", map[string]interface{}{"error": err.Error()})
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, "Invalid user ID"))
		return
	}

	var req model.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Invalid request body", map[string]interface{}{"error": err.Error()})
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, "Invalid request body"))
		return
	}

	user := &model.User{
		Name:  req.Name,
		Email: req.Email,
		Age:   req.Age,
	}

	err = h.userMock.Update(id, user)
	if err != nil {
		h.logger.Error("Failed to update user", map[string]interface{}{"error": err.Error(), "user_id": id})
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, "Failed to update user"))
		return
	}

	// 检查用户是否存在
	existingUser, err := h.userMock.GetByID(id)
	if err != nil {
		h.logger.Error("Failed to get user", map[string]interface{}{"error": err.Error(), "user_id": id})
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, "Failed to get user"))
		return
	}

	if existingUser == nil {
		h.logger.Error("User not found", map[string]interface{}{"user_id": id})
		c.JSON(http.StatusNotFound, model.ErrorResponse(http.StatusNotFound, "User not found"))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(existingUser))
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
		h.logger.Error("Invalid user ID", map[string]interface{}{"error": err.Error()})
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, "Invalid user ID"))
		return
	}

	err = h.userMock.Delete(id)
	if err != nil {
		h.logger.Error("Failed to delete user", map[string]interface{}{"error": err.Error(), "user_id": id})
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, "Failed to delete user"))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(nil))
}
