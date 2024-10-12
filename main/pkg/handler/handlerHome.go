package handler

import (
	"net/http"

	"github.com/Mamvriyskiy/database_course/main/logger"
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createHome(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		logger.Log("Warning", "Get", "Error get userID from context", nil, id)
		return
	}

	userID, ok := id.(string)
	if !ok {
		logger.Log("Warning", "Get", "userID is not a string", nil, "userID")
		return
	}

	var input pkg.HomeHandler
	if err := c.BindJSON(&input); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	home, err := h.services.IHome.CreateHome(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"errors": "Ошибка создания дома",
		})
		logger.Log("Error", "CreateHome", "Error create home:", err, userID, input)
		return
	}

	_, err = h.services.IAccessHome.AddOwner(userID, home.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"errors": "Ошибка добавления хозяина дома",
		})
		logger.Log("Error", "AddOwner", "Error add owner:", err, userID, home)
		return
	}

	c.JSON(http.StatusOK, home)

	logger.Log("Info", "", "A home has been created", nil)
}

func (h *Handler) deleteHome(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		logger.Log("Warning", "Get", "Error get userID from context", nil, id)
		return
	}

	userID, ok := id.(string)
	if !ok {
		logger.Log("Warning", "Get", "userID is not a string", nil, "userID")
		return
	}

	homeID := c.Param("homeID")

	accessLevel, err := h.services.IUser.GetAccessLevel(userID, homeID)
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
		logger.Log("Warning", "Get", "Error get userID from context", nil, id)
		return
	}

	userID, ok := id.(string)
	if !ok {
		logger.Log("Warning", "Get", "userID is not a string", nil, "userID")
		return
	}

	homeListUser, err := h.services.IHome.ListUserHome(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"errors": "Ошибка получения списка домов",
		})
		logger.Log("Error", "ListUserHome", "Error get user:", err, userID)
		return
	}

	c.JSON(http.StatusOK, homeListUser) //вернуть пустой массив

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
		logger.Log("Warning", "Get", "Error get userID from context", nil, id)
		return
	}

	userID, ok := id.(string)
	if !ok {
		logger.Log("Warning", "Get", "userID is not a string", nil, "userID")
		return
	}

	homeID := c.Param("homeID")

	var input pkg.HomeHandler
	err := c.BindJSON(&input)
	if err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	accessLevel, err := h.services.IUser.GetAccessLevel(userID, homeID)
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
