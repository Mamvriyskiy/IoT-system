package handler

import (
	"net/http"
	"strconv"

	"github.com/Mamvriyskiy/DBCourse/main/logger"
	"github.com/Mamvriyskiy/DBCourse/main/pkg"
	"github.com/gin-gonic/gin"
)

type tmp struct {
	Email       string `json:"email"`
	AccessLevel string `json:"accessLevel"`
}

func (h *Handler) addUser(c *gin.Context) {
	userID, ok := c.Get("userId")
	if !ok {
		logger.Log("Warning", "Get", "Error get userID from context", nil, "userID")
		return
	}

	var res tmp
	if err := c.BindJSON(&res); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	num, err := strconv.Atoi(res.AccessLevel)
	if err != nil {
		logger.Log("Error", "strconv.Atoi(res.AccessLevel)", "Error str to int:", err, "")
	}

	input := pkg.AddUserHome{
		Email:       res.Email,
		AccessLevel: num,
	}

	var intVal float64
	if val, ok := userID.(float64); ok {
		intVal = val
	} else {
		logger.Log("Error", "userID.(float64)", "Error:", ErrNoFloat64Interface, "")
	}

	idAccess, err := h.services.IAccessHome.AddUser(int(intVal), input.AccessLevel, input.Email)
	if err != nil {
		logger.Log("Error", "AddUser", "Error create access:",
			err, int(intVal), input.AccessLevel, input.Email)
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

	var input pkg.AddUserHome
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

	err := h.services.IAccessHome.DeleteUser(int(intVal), input.Email)
	if err != nil {
		logger.Log("Error", "DeleteUser", "Error delete access:", err, int(intVal), input.Email)
		return
	}

	logger.Log("Info", "", "The user's access was deleted", nil)
}

func (h *Handler) updateLevel(c *gin.Context) {
	userID, ok := c.Get("userId")
	if !ok {
		logger.Log("Warning", "Get", "Error get userID from context", nil, "userID")
		return
	}

	var res tmp
	if err := c.BindJSON(&res); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	num, err := strconv.Atoi(res.AccessLevel)
	if err != nil {
		logger.Log("Error", "strconv.Atoi(res.AccessLevel)", "Error str to int:", err, "")
	}

	input := pkg.AddUserHome{
		Email:       res.Email,
		AccessLevel: num,
	}

	var intVal float64
	if val, ok := userID.(float64); ok {
		intVal = val
	} else {
		logger.Log("Error", "userID.(float64)", "Error:", ErrNoFloat64Interface, "")
	}

	err = h.services.IAccessHome.UpdateLevel(int(intVal), input)
	if err != nil {
		logger.Log("Error", "UpdateLevel", "Error update access:", err, intVal, input)
		return
	}

	logger.Log("Info", "", "A level has been update", nil)
}

func (h *Handler) updateStatus(c *gin.Context) {
	idUser := 2
	var input pkg.AccessHome
	if err := c.BindJSON(&input); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	err := h.services.IAccessHome.UpdateStatus(idUser, input)
	if err != nil {
		logger.Log("Error", "UpdateStatus", "Error update access:", err, idUser, input)
		return
	}

	logger.Log("Info", "", "A status has been update", nil)
}

type getAllListUserResponse struct {
	Data []pkg.ClientHome `json:"data"`
}

func (h *Handler) getListUserHome(c *gin.Context) {
	homeID := 1
	listUser, err := h.services.IAccessHome.GetListUserHome(homeID)
	if err != nil {
		logger.Log("Error", "GetListUserHome", "Error get access:", err, homeID)
		return
	}

	c.JSON(http.StatusOK, getAllListUserResponse{
		Data: listUser,
	})

	logger.Log("Info", "", "The list of users has been received", nil)
}
