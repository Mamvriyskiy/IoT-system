package handler

import (
	"fmt"
	"net/http"

	"github.com/Mamvriyskiy/database_course/main/logger"
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/gin-gonic/gin"
)

type verifyCode struct {
	Code  string `db:"code" json:"verificationCode"`
	Token string `db:"token" json:"token"`
}

func (h *Handler) checkCode(c *gin.Context) {
	var input verifyCode
	if err := c.BindJSON(&input); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	err := h.services.CheckCode(input.Code, input.Token)
	if err != nil {
		logger.Log("Error", "h.services.CheckCode(codeID)", "Error CheckCode:", err, input)
		c.JSON(400, map[string]interface{}{})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{})
}

func (h *Handler) sendCode(c *gin.Context) {
	var input pkg.Email
	if err := c.BindJSON(&input); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	err := h.services.IUser.SendCode(input)
	if err != nil {
		// *TODO log
		return
	}
}

type newpassword struct {
	Password string `db:"password" json:"password"`
	Token    string `db:"token" json:"token"`
}

func (h *Handler) changePassword(c *gin.Context) {
	var input newpassword
	if err := c.BindJSON(&input); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	err := h.services.IUser.ChangePassword(input.Password, input.Token)
	if err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{})
}

func (h *Handler) signUp(c *gin.Context) {
	var input pkg.User
	if err := c.BindJSON(&input); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	if count, err := h.services.IUser.GetUserByEmail(input.Email); err != nil || count != 0 {
		logger.Log("Info", "CheckUser(user pkg.User)", fmt.Sprintf("User already register: %s", input.Email), nil)
		c.JSON(http.StatusOK, map[string]interface{}{
			"errors": "Пользователь уже зарегистрирован",
		})
		return
	}

	id, err := h.services.IUser.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{})
		logger.Log("Error", "h.services.IUser.CreateUser(input)", "Error create user:", err, input)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

	logger.Log("Info", "", fmt.Sprintf("User %s is registered", input.Username), nil)
}

type signInInput struct {
	Password string `json:"password"`
	Email string `json:"email"`
}

// Протаскивать ошибку из сервера и БД, ее номер
// Создать собственные ошибки
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	status := 0
	token, err := h.services.IUser.GenerateToken(input.Email, input.Password)
	if err != nil {
		status = http.StatusNotFound
		c.JSON(http.StatusNotFound, map[string]interface{}{})
		logger.Log("Error", "GenerateToken", "Error GenerateToken:", err, input)
		return
	}

	c.JSON(status, map[string]interface{}{
		"token": token,
	})

	logger.Log("Info", "", fmt.Sprintf("User %s ganied access", input.Email), nil)
}
