package handler

import (
	"net/http"
	// "fmt"

	"github.com/Mamvriyskiy/DBCourse/main/logger"
	"github.com/Mamvriyskiy/DBCourse/main/pkg"
	"github.com/gin-gonic/gin"
)

func (h *Handler) addUser(c *gin.Context) {
	userID, ok := c.Get("userId")
	if !ok {
		logger.Log("Warning", "Get", "Error get userID from context", nil, "userID")
		return
	}

	var access pkg.Access
	if err := c.BindJSON(&access); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	var intVal float64
	if val, ok := userID.(float64); ok {
		intVal = val
	} else {
		logger.Log("Error", "userID.(float64)", "Error:", ErrNoFloat64Interface, "")
	}

	idAccess, err := h.services.IAccessHome.AddUser(int(intVal), access)
	if err != nil {
		logger.Log("Error", "AddUser", "Error create access:",
			err, int(intVal), access.AccessLevel, access.Email)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"accessID": idAccess,
	})

	logger.Log("Info", "", "The user has been granted access", nil)
}

func (h *Handler) deleteUser(c *gin.Context) {
	userID, ok := c.Get("userId")
	if !ok {
		logger.Log("Warning", "Get", "Error get userID from context", nil, "userID")
		return
	}

	var input pkg.Access
	if err := c.BindJSON(&input); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	var intVal float64
	if val, ok := userID.(float64); ok {
		intVal = val
	} else {
		logger.Log("Error", "userID.(float64)", "Error:", ErrNoFloat64Interface, "")
	}

	err := h.services.IAccessHome.DeleteUser(int(intVal), input)
	if err != nil {
		logger.Log("Error", "DeleteUser", "Error delete access:", err, int(intVal), input)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{})

	logger.Log("Info", "", "The user's access was deleted", nil)
}

func (h *Handler) updateLevel(c *gin.Context) {
	userID, ok := c.Get("userId")
	if !ok {
		logger.Log("Warning", "Get", "Error get userID from context", nil, "userID")
		return
	}

	var update pkg.Access
	if err := c.BindJSON(&update); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	var intVal float64
	if val, ok := userID.(float64); ok {
		intVal = val
	} else {
		logger.Log("Error", "userID.(float64)", "Error:", ErrNoFloat64Interface, "")
	}

	err := h.services.IAccessHome.UpdateLevel(int(intVal), update)
	if err != nil {
		logger.Log("Error", "UpdateLevel", "Error update access:", err, intVal, update)
		return
	}

	c.JSON(http.StatusOK, getAllListUserResponse{})

	logger.Log("Info", "", "A level has been update", nil)
}

func (h *Handler) updateStatus(c *gin.Context) {
	userID, ok := c.Get("userId")
	if !ok {
		logger.Log("Warning", "Get", "Error get userID from context", nil, "userID")
		return
	}

	var input pkg.AccessHome
	if err := c.BindJSON(&input); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	var intVal float64
	if val, ok := userID.(float64); ok {
		intVal = val
	} else {
		logger.Log("Error", "userID.(float64)", "Error:", ErrNoFloat64Interface, "")
	}

	err := h.services.IAccessHome.UpdateStatus(int(intVal), input)
	if err != nil {
		logger.Log("Error", "UpdateStatus", "Error update access:", err, userID, input)
		return
	}

	logger.Log("Info", "", "A status has been update", nil)
}

type getAllListUserResponse struct {
	Data []pkg.ClientHome `json:"data"`
}

func (h *Handler) getListUserHome(c *gin.Context) {
	userID, ok := c.Get("userId")
	if !ok {
		logger.Log("Warning", "Get", "Error get userID from context", nil, "userID")
		return
	}

	var intVal float64
	if val, ok := userID.(float64); ok {
		intVal = val
	} else {
		logger.Log("Error", "userID.(float64)", "Error:", ErrNoFloat64Interface, "")
	}

	listUser, err := h.services.IAccessHome.GetListUserHome(int(intVal))
	if err != nil {
		logger.Log("Error", "GetListUserHome", "Error get access:", err, int(intVal))
		return
	}

	// fmt.Println(listUser)

	c.JSON(http.StatusOK, getAllListUserResponse{
		Data: listUser,
	})

	logger.Log("Info", "", "The list of users has been received", nil)
}
