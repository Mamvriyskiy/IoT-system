package handler

import (
	"fmt"
	"net/http"
	"bytes"
	"io"
	"github.com/Mamvriyskiy/database_course/main/logger"
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/gin-gonic/gin"
	"net/mail"
)

func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
    return err == nil
}

func (h *Handler) checkCode(c *gin.Context) {
	var input pkg.VerifyCode
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{})
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	fmt.Println(input)
	err := h.services.CheckCode(input.Code, input.Token)
	if err != nil {
		logger.Log("Error", "h.services.CheckCode(codeID)", "Error CheckCode:", err, input)
		c.JSON(http.StatusBadRequest, map[string]interface{}{})
		return
	}  

	c.JSON(http.StatusOK, map[string]interface{}{})
}

func (h *Handler) code(c *gin.Context) {
	var input pkg.EmailHandler
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{})
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	err := h.services.IUser.SendCode(input)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{})
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{})
}

func (h *Handler) changePassword(c *gin.Context) {
	var input pkg.UpdatePassword
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{})
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	fmt.Println(input)
	err := h.services.IUser.ChangePassword(input.Password, input.Token)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{})
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	c.JSON(http.StatusOK, struct{}{})
}

func (h *Handler) SignUp(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read request body"})
        return
    }

    // Логируем тело запроса
    fmt.Println("Request Body:", string(body))

    // Восстанавливаем тело запроса, чтобы его можно было использовать позже
    c.Request.Body = io.NopCloser(bytes.NewReader(body))
	var input pkg.UserHandler
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{})
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	fmt.Println(input)
	if !isEmailValid(input.Email) {
		logger.Log("Info", "isEmailValid", fmt.Sprintf("Invalid email: %s", input.Email), nil)
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors": "Неверный формат почты",
		})
		return
	}

	if count, err := h.services.IUser.GetUserByEmail(input.Email); err != nil || count != 0 {
		logger.Log("Info", "CheckUser(user pkg.User)", fmt.Sprintf("User already register: %s", input.Email), nil)
		c.JSON(http.StatusConflict, map[string]interface{}{
			"errors": "Пользователь уже зарегистрирован",
		})
		return
	}
	fmt.Println("+")
	_, err = h.services.IUser.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, struct{}{})
		logger.Log("Error", "h.services.IUser.CreateUser(input)", "Error create user:", err, input)
		return
	}

	c.JSON(http.StatusOK, struct{}{})

	logger.Log("Info", "", fmt.Sprintf("User %s is registered", input.Username), nil)
}

func (h *Handler) signIn(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read request body"})
        return
    }

    // Логируем тело запроса
    fmt.Println("Request Body:", string(body))

    // Восстанавливаем тело запроса, чтобы его можно было использовать позже
    c.Request.Body = io.NopCloser(bytes.NewReader(body))

	var input pkg.UserHandler
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{})
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}	

	fmt.Println(input.Email)

	// if !isEmailValid(input.Email) {
	// 	fmt.Println("Invalid mail")
	// 	logger.Log("Info", "isEmailValid", fmt.Sprintf("Invalid email: %s", input.Email), nil)
	// 	c.JSON(http.StatusBadRequest, map[string]interface{}{
	// 		"errors": "Неверный формат почты",
	// 	})
	// 	return
	// }

	user, token, err := h.services.IUser.GenerateToken(input.Email, input.Password)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, map[string]interface{}{})
		logger.Log("Error", "GenerateToken", "Error GenerateToken:", err, input)
		return
	}

	fmt.Println(user, token, err)
	c.JSON(http.StatusOK, map[string]interface{}{
		"Token": token,
		"Login": user.Username,
		"Email": user.Email,
	})

	logger.Log("Info", "", fmt.Sprintf("User %s ganied access", input.Email), nil)
}
