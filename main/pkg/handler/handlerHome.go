package handler

import (
	"net/http"

	"github.com/Mamvriyskiy/database_course/main/logger"
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) createHome(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		logger.Log("Warning", "Get", "Error get userID from context", nil, "userId")
		return
	}

	userID, ok := id.(float64)
	if !ok {
		c.JSON(http.StatusOK, map[string]interface{}{
			"errors": "Ошибка создания дома",
		})
		logger.Log("Error", "userID.(float64)", "Error:", ErrNoFloat64Interface, "")
		return
	}

	var input pkg.Home
	if err := c.BindJSON(&input); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	homeID, err := h.services.IHome.CreateHome(input)
	if err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"errors": "Ошибка создания дома",
		})
		logger.Log("Error", "CreateHome", "Error create home:", err, userID, input)
		return
	}

	homeIDStr := strconv.Itoa(homeID)

	_, err = h.services.IAccessHome.AddOwner(int(userID), homeIDStr)
	if err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"errors": "Ошибка добавления хозяина дома",
		})
		logger.Log("Error", "AddOwner", "Error add owner:", err, userID, homeID)
		return
	}

	input.ID = homeID

	c.JSON(http.StatusOK, input)

	logger.Log("Info", "", "A home has been created", nil)
}

func (h *Handler) deleteHome(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		logger.Log("Warning", "Get", "Error get userID from context", nil, "userID")
		return
	}

	userID, ok := id.(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"errors": "Ошибка обновления статуса",
		})
		logger.Log("Error", "userID.(float64)", "Error:", ErrNoFloat64Interface, "")
		return
	}

	homeID := c.Param("homeID")

	accessLevel, err := h.services.IUser.GetAccessLevel(int(userID), homeID)
	if accessLevel != 4 || err != nil {
		c.JSON(http.StatusForbidden, map[string]string{
			"errors": "Недостаточно прав для удаления",
		})
		logger.Log("Error", "GetAccessLevel", "Error GetAccessLevel home:", err, accessLevel)
		return
	}

	err = h.services.IHome.DeleteHome(homeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"errors": "дом не найден",
		})
		logger.Log("Error", "DeleteHome", "Error delete home:", err, homeID)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{})

	logger.Log("Info", "", "A home has been deleted", nil)
}

func (h *Handler) listHome(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		logger.Log("Warning", "Get", "Error get userID from context", nil, "userID")
		return
	}

	userID, ok := id.(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"errors": "Ошибка обновления статуса",
		})
		logger.Log("Error", "userID.(float64)", "Error:", ErrNoFloat64Interface, "")
		return
	}

	homeListUser, err := h.services.IHome.ListUserHome(int(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"errors": "Ошибка получения списка домов",
		})
		logger.Log("Error", "ListUserHome", "Error get user:", err, id.(int))
		return
	}

	c.JSON(http.StatusOK, homeListUser)

	logger.Log("Info", "", "The list of users has been received", nil)
}

func (h *Handler) infoHome(c *gin.Context) {
	homeID := c.Param("homeID")

	home, err := h.services.IHome.GetHomeByID(homeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": "дом не найден",
		})
		logger.Log("Error", "UpdateHome", "Error update home:", err, "")
		return
	}

	c.JSON(http.StatusOK, home)

	logger.Log("Info", "", "A home has been update", nil)
}

func (h *Handler) updateHome(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		logger.Log("Warning", "Get", "Error get userID from context", nil, "userID")
		return
	}

	userID, ok := id.(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"errors": "Ошибка обновления статуса",
		})
		logger.Log("Error", "userID.(float64)", "Error:", ErrNoFloat64Interface, "")
		return
	}
	
	homeID := c.Param("homeID")

	var input pkg.Home
	err := c.BindJSON(&input)
	if err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	accessLevel, err := h.services.IUser.GetAccessLevel(int(userID), homeID)
	if accessLevel != 4 || err != nil {
		c.JSON(http.StatusForbidden, map[string]string{
			"errors": "Недостаточно прав для удаления",
		})
		logger.Log("Error", "GetAccessLevel", "Error GetAccessLevel home:", err, accessLevel)
		return
	}

	err = h.services.IHome.UpdateHome(homeID, input.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": "дом не найден",
		})
		logger.Log("Error", "UpdateHome", "Error update home:", err, "")
		return
	}

	home, err := h.services.IHome.GetHomeByID(homeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": "дом не найден",
		})
		logger.Log("Error", "UpdateHome", "Error update home:", err, "")
		return
	}

	c.JSON(http.StatusOK, home)

	logger.Log("Info", "", "A home has been update", nil)
}
