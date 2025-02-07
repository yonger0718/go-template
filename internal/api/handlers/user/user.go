package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-template/internal/api/handlers"
	"go-template/internal/constants"
	"go-template/internal/models"
	"go-template/internal/services"
	"go-template/internal/utils/logutil"
	"go-template/internal/utils/response"
	"go-template/internal/validators"
)

// Handler struct，用於處理使用者相關的 HTTP 請求
type Handler struct {
	userService services.UserService
}

// NewUserHandler 建立一個新的 UserHandler 實例
func NewUserHandler(userService services.UserService) *Handler {
	return &Handler{userService: userService}
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RegisterUser 處理使用者註冊的請求
// @Summary 註冊使用者
// @Description 註冊一個新的使用者
// @Tags User
// @Accept  json
// @Produce  json
// @Param user body models.User true "使用者資料"
// @Success 201 {object} response.SuccessResponseData{data=models.User} "註冊成功"
// @Failure 400 {object} response.ErrorResponseData "錯誤的請求"
// @Failure 500 {object} response.ErrorResponseData "系統錯誤"
// @Router /user/register [post]
func (h *Handler) RegisterUser(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
	}
	// 解析請求的 JSON 數據到 user 變數
	if err := c.ShouldBindJSON(&input); err != nil {
		logutil.Logger.Debugf(handlers.ErrMsgInvalidRequestBody, err) // DEBUG 等級
		response.ErrorResponse(c, http.StatusBadRequest, handlers.ErrCodeInvalidRequest)
		return
	}

	// 驗證使用者資料
	if err := validators.ValidateNewUser(input.Username, input.Email, input.Password); err != nil {
		logutil.Logger.Debugf("Invalid user data: %v", err)
		response.ErrorResponse(c, http.StatusBadRequest, handlers.ErrCodeInvalidRequest)
		return
	}

	// 建立一個新的 User 實例
	user := models.User{
		Username: input.Username,
		Password: input.Password,
		Email:    input.Email,
	}

	// 呼叫 userService 建立使用者
	if err := h.userService.CreateUser(&user); err != nil {
		logutil.Logger.Errorf("Error creating user: %v", err) // ERROR 等級
		response.ErrorResponse(c, http.StatusInternalServerError, handlers.ErrCodeUnknown)
		return
	}

	// 回應註冊成功的訊息
	logutil.Logger.Infof("User created: %s", user.Username) // INFO 等級
	response.SuccessResponse(c, http.StatusCreated, "User created successfully", user)
}

// LoginUser 處理使用者登入的請求
// @Summary 登入使用者
// @Description 登入一個已註冊的使用者
// @Tags User
// @Accept  json
// @Produce  json
// @Param credentials body loginRequest true "使用者登入資訊"
// @Success 200 {object} response.SuccessResponseData{Data=string} "登入成功"
// @Failure 400 {object} response.ErrorResponseData "錯誤的請求"
// @Failure 401 {object} response.ErrorResponseData "使用者不存在或密碼錯誤"
// @Failure 500 {object} response.ErrorResponseData "系統錯誤"
// @Router /user/login [post]
func (h *Handler) LoginUser(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	// 解析請求的 JSON 數據到 input 變數
	if err := c.ShouldBindJSON(&input); err != nil {
		logutil.Logger.Debugf(handlers.ErrMsgInvalidRequestBody, err) // DEBUG 等級
		response.ErrorResponse(c, http.StatusBadRequest, handlers.ErrCodeInvalidRequest)
		return
	}

	// 呼叫 userService 進行使用者登入
	token, err := h.userService.Login(input.Username, input.Password)
	if err != nil {
		logutil.Logger.Errorf("Error logging in: %v", err) // ERROR 等級
		// 根據不同的錯誤類型回覆不同的錯誤碼
		switch err {
		case services.ErrUserNotFound:
			response.ErrorResponse(c, http.StatusUnauthorized, handlers.ErrCodeUserNotFound)
		case services.ErrInvalidCredentials:
			response.ErrorResponse(c, http.StatusUnauthorized, handlers.ErrCodeInvalidCredentials)
		default:
			response.ErrorResponse(c, http.StatusInternalServerError, handlers.ErrCodeUnknown)
		}
		return
	}

	// 回應登入成功的訊息和 JWT token
	logutil.Logger.Infof("User logged in: %s", input.Username) // INFO 等級
	response.SuccessResponse(c, http.StatusOK, "Login successful", gin.H{"token": token})
}

// GetUser 處理取得使用者資訊的請求
// @Summary 取得使用者資訊
// @Description 取得指定使用者的資訊
// @Tags User
// @Produce  json
// @Param id path int true "User ID"
// @Security BearerAuth
// @Success 200 {object} response.SuccessResponseData{Data=models.User} "取得成功"
// @Failure 404 {object} response.ErrorResponseData "使用者不存在"
// @Failure 500 {object} response.ErrorResponseData "系統錯誤"
// @Router /user/{id} [get]
func (h *Handler) GetUser(c *gin.Context) {
	// 從 gin.Context 中取得 userID
	userID, exists := c.Get(constants.CtxUserIDKey)
	if !exists {
		logutil.Logger.Debugf(handlers.ErrMsgUserIDNotInContext) // DEBUG 等級
		response.ErrorResponse(c, http.StatusInternalServerError, handlers.ErrCodeUserIDNotInContext)
		return
	}

	// 將 userID 轉成 uint 型別
	id, ok := userID.(uint)
	if !ok {
		logutil.Logger.Debugf(handlers.ErrMsgUserIDFormatInvalid) // DEBUG 等級
		response.ErrorResponse(c, http.StatusInternalServerError, handlers.ErrCodeUserIDFormatInvalid)
		return
	}

	// 呼叫 userService 取得使用者資訊
	user, err := h.userService.GetUserByID(id)
	if err != nil {
		logutil.Logger.Errorf("Error getting user: %v", err) // ERROR 等級
		// 根據不同的錯誤類型回覆不同的錯誤碼
		if err == services.ErrUserNotFound {
			response.ErrorResponse(c, http.StatusNotFound, handlers.ErrCodeUserNotFound)
		} else {
			response.ErrorResponse(c, http.StatusInternalServerError, handlers.ErrCodeUnknown)
		}
		return
	}

	// 回應使用者資訊
	logutil.Logger.Debugf("User found: %s", user.Username) // DEBUG 等級
	response.SuccessResponse(c, http.StatusOK, "User found", user)
}

// UpdateUser 處理更新使用者資訊的請求
// @Summary 更新使用者資訊
// @Description 更新指定使用者的資訊
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path int true "使用者 ID"
// @Security BearerAuth
// @Param user body models.User true "使用者資料"
// @Success 200 {object} response.SuccessResponseData{Data=models.User} "更新成功"
// @Failure 400 {object} response.ErrorResponseData "錯誤的請求"
// @Failure 404 {object} response.ErrorResponseData "使用者不存在"
// @Failure 500 {object} response.ErrorResponseData "系統錯誤"
// @Router /user/{id} [put]
func (h *Handler) UpdateUser(c *gin.Context) {
	// 從 gin.Context 中取得 userID
	userID, exists := c.Get(constants.CtxUserIDKey)
	if !exists {
		logutil.Logger.Debugf(handlers.ErrMsgUserIDNotInContext) // DEBUG 等級
		response.ErrorResponse(c, http.StatusInternalServerError, handlers.ErrCodeUserIDNotInContext)
		return
	}

	// 將 userID 轉成 uint 型別
	id, ok := userID.(uint)
	if !ok {
		logutil.Logger.Debugf(handlers.ErrMsgUserIDFormatInvalid) // DEBUG 等級
		response.ErrorResponse(c, http.StatusInternalServerError, handlers.ErrCodeUserIDFormatInvalid)
		return
	}

	var user models.User
	// 解析請求的 JSON 數據到 user 變數
	if err := c.ShouldBindJSON(&user); err != nil {
		logutil.Logger.Debugf(handlers.ErrMsgInvalidRequestBody, err) // DEBUG 等級
		response.ErrorResponse(c, http.StatusBadRequest, handlers.ErrCodeInvalidRequest)
		return
	}

	// 將使用者 ID 設定為從 token 中取得的 ID
	user.ID = id
	// 呼叫 userService 更新使用者資訊
	if err := h.userService.UpdateUser(&user); err != nil {
		logutil.Logger.Errorf("Error updating user: %v", err) // ERROR 等級
		response.ErrorResponse(c, http.StatusInternalServerError, handlers.ErrCodeUnknown)
		return
	}

	// 回應更新成功的訊息
	logutil.Logger.Info("User updated") // INFO 等級
	response.SuccessResponse(c, http.StatusOK, "User updated successfully", user)
}

// DeleteUser 處理刪除使用者的請求
// @Summary 刪除使用者
// @Description 刪除指定使用者
// @Tags User
// @Produce  json
// @Param id path int true "使用者 ID"
// @Security BearerAuth
// @Success 200 {object} response.SuccessResponseData "刪除成功"
// @Failure 404 {object} response.ErrorResponseData "使用者不存在"
// @Failure 500 {object} response.ErrorResponseData "系統錯誤"
// @Router /user/{id} [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	// 從 gin.Context 中取得 userID
	userID, exists := c.Get(constants.CtxUserIDKey)
	if !exists {
		logutil.Logger.Debugf(handlers.ErrMsgUserIDNotInContext) // DEBUG 等級
		response.ErrorResponse(c, http.StatusInternalServerError, handlers.ErrCodeUserIDNotInContext)
		return
	}

	// 將 userID 轉成 uint 型別
	id, ok := userID.(uint)
	if !ok {
		logutil.Logger.Debugf(handlers.ErrMsgUserIDFormatInvalid) // DEBUG 等級
		response.ErrorResponse(c, http.StatusInternalServerError, handlers.ErrCodeUserIDFormatInvalid)
		return
	}

	// 呼叫 userService 刪除使用者
	if err := h.userService.DeleteUser(id); err != nil {
		logutil.Logger.Errorf("Error deleting user: %v", err) // ERROR 等級
		response.ErrorResponse(c, http.StatusInternalServerError, handlers.ErrCodeUnknown)
		return
	}

	// 回應刪除成功的訊息
	logutil.Logger.Info("User deleted") // INFO 等級
	response.SuccessResponse(c, http.StatusOK, "User deleted successfully", nil)
}
