package handler

import (
	"net/http"
	// "fmt"

	"github.com/Mamvriyskiy/database_course/main/logger"
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/gin-gonic/gin"
)

func (h *Handler) addUser(c *gin.Context) {
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

	
	var access pkg.Access
	if err := c.BindJSON(&access); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	if !isEmailValid(access.Email) {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors": "Неверный формат почты",
		})
		return
	}

	accessID, err := h.services.IAccessHome.AddUser(homeID, access)
	if err != nil {
		logger.Log("Error", "AddUser", "Error create access:",
			err, homeID, access.AccessLevel, access.Email)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"errors": "Ошибка добавления пользователя",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ID": accessID,
		"Email": access.Email,
		"AccessLevel": access.AccessLevel,
	})

	logger.Log("Info", "", "The user has been granted access", nil)
}

func (h *Handler) deleteUser(c *gin.Context) {
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

	accessID := c.Param("accessID")

	err = h.services.IAccessHome.DeleteUser(accessID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"errors": "Ошибка удаления пользователя",
		})
		logger.Log("Error", "DeleteUser", "Error delete access:", err, accessID)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{})

	logger.Log("Info", "", "The user's access was deleted", nil)
}

func (h *Handler) updateLevel(c *gin.Context) {
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

	accessID := c.Param("accessID")

	var access pkg.Access
	if err := c.BindJSON(&access); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	if !isEmailValid(access.Email) {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors": "Неверный формат почты",
		})
		return
	}

	err = h.services.IAccessHome.UpdateLevel(accessID, access)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"errors": "Ошибка получения списка пользователей",
		})
		logger.Log("Error", "GetInfoAccessByID", "Error get access:", err, accessID)
		return
	}

	accessInfo, err := h.services.IAccessHome.GetInfoAccessByID(accessID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"errors": "Ошибка получения списка пользователей",
		})
		logger.Log("Error", "GetInfoAccessByID", "Error get access:", err, accessID)
		return
	}


	c.JSON(http.StatusOK, map[string]interface{}{
		"AccessID": accessInfo.ID,
		"Login": accessInfo.Login,
		"AccessLevel": accessInfo.AccessLevel,
		"Email": accessInfo.Email,
	})

	logger.Log("Info", "", "Info of access", nil)
}

func (h *Handler) updateStatus(c *gin.Context) {
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

	var input pkg.AccessHome
	if err := c.BindJSON(&input); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	err := h.services.IAccessHome.UpdateStatus(userID, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"errors": "Ошибка обновления статуса",
		})
		logger.Log("Error", "UpdateStatus", "Error update access:", err, userID, input)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{})

	logger.Log("Info", "", "A status has been update", nil)
}

func (h *Handler) getListUserHome(c *gin.Context) {
	homeID := c.Param("homeID")

	listUser, err := h.services.IAccessHome.GetListUserHome(homeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"errors": "Ошибка получения списка пользователей",
		})
		logger.Log("Error", "GetListUserHome", "Error get access:", err, homeID)
		return
	}

	c.JSON(http.StatusOK,listUser)

	logger.Log("Info", "", "The list of users has been received", nil)
}

func (h *Handler) getInfoAccess(c *gin.Context) {
	accessID := c.Param("accessID")

	accessInfo, err := h.services.IAccessHome.GetInfoAccessByID(accessID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"errors": "Ошибка получения списка пользователей",
		})
		logger.Log("Error", "GetInfoAccessByID", "Error get access:", err, accessID)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"AccessID": accessInfo.ID,
		"Login": accessInfo.Login,
		"AccessLevel": accessInfo.AccessLevel,
		"Email": accessInfo.Email,
	})

	logger.Log("Info", "", "Info of access", nil)
}
