package handler

import (
	"fmt"
	"net/http"

	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/logger"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/gin-gonic/gin"
)


type verifyCode struct {
	Code  string    `db:"code" json:"verificationCode"`
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
	Password  string    `db:"password" json:"password"`
	Token string `db:"token" json:"token"`
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

	id, err := h.services.IUser.CreateUser(input)
	if err != nil {
		// *TODO log
		return
	}

	// c.Next()

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

	logger.Log("Info", "", fmt.Sprintf("User %s is registered", input.Username), nil)
}

type signInInput struct {
	Password string `json:"password"`
	Username string `json:"login"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	token, err := h.services.IUser.GenerateToken(input.Username, input.Password)
	if err != nil {
		// *TODO log
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

	logger.Log("Info", "", fmt.Sprintf("User %s ganied access", input.Username), nil)
}