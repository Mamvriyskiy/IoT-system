package handler

import (
	"errors"
	"net/http"

	"github.com/Mamvriyskiy/DBCourse/main/logger"
	"github.com/Mamvriyskiy/DBCourse/main/pkg"
	"github.com/gin-gonic/gin"
)

var ErrNoFloat64Interface = errors.New("отсутствует интерфейс {} для float64")

type getAllListDevices struct {
	Data []pkg.Devices `json:"data"`
}

func (h *Handler) getListDevice(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		logger.Log("Warning", "Get", "Error get userID from context", nil, "userID")
		return
	}

	var intVal float64
	if val, ok := id.(float64); ok {
		intVal = val
	} else {
		logger.Log("Error", "userID.(float64)", "Error:", ErrNoFloat64Interface, "")
	}

	listDevices, err := h.services.IDevice.GetListDevices(int(intVal))
	if err != nil {
		logger.Log("Error", "GetListDevices", "Error get list devices:", err, id)
		return
	}

	c.JSON(http.StatusOK, getAllListDevices{
		Data: listDevices,
	})

	logger.Log("Info", "", "The history of the device was obtained", nil)
}

func (h *Handler) createDevice(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		logger.Log("Warning", "Get", "Error get userID from context", nil, "userID")
		return
	}

	var input pkg.Devices
	if err := c.BindJSON(&input); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	var intVal float64
	if val, ok := id.(float64); ok {
		intVal = val
	} else {
		logger.Log("Error", "userID.(float64)", "Error:", ErrNoFloat64Interface, "")
	}

	idDevice, err := h.services.IDevice.CreateDevice(int(intVal), &input)
	if err != nil {
		logger.Log("Error", "CreateDevice", "Error create device:", err, id, &input)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"idDevice": idDevice,
	})

	logger.Log("Info", "", "A device has been created", nil)
}

func (h *Handler) deleteDevice(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		logger.Log("Warning", "Get", "Error get userID from context", nil, "userID")
		return
	}

	var input pkg.Devices
	if err := c.BindJSON(&input); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	var intVal float64
	if val, ok := id.(float64); ok {
		intVal = val
	} else {
		logger.Log("Error", "userID.(float64)", "Error:", ErrNoFloat64Interface, "")
	}

	err := h.services.IDevice.DeleteDevice(int(intVal), input.Name)
	if err != nil {
		logger.Log("Error", "DeleteDevice", "Error delete device:", err, id, input.Name)
		return
	}

	logger.Log("Info", "", "A device has been deleted", nil)
}
