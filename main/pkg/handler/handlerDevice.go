package handler

import (
	"errors"
	"net/http"

	"github.com/Mamvriyskiy/database_course/main/logger"
	"github.com/Mamvriyskiy/database_course/main/pkg"
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
		c.JSON(http.StatusOK, map[string]interface{}{
			"errors": "Ошибка получения списка устройств",
		})
		logger.Log("Error", "userID.(float64)", "Error:", ErrNoFloat64Interface, "")
		return 
	}

	listDevices, err := h.services.IDevice.GetListDevices(int(intVal))
	if err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"errors": "Ошибка получения списка пользователей",
		})
		logger.Log("Error", "GetListDevices", "Error get list devices:", err, id)
		return
	}

	if listDevices == nil || len(listDevices) == 0 {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"errors": "список устройств пуст",
		})
		return 
	}

	logger.Log("Info", "ListDev:", "Error get list devices:", nil, listDevices, len(listDevices))

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
		c.JSON(http.StatusOK, map[string]interface{}{
			"errors": "Ошибка создания устройства",
		})
		logger.Log("Error", "userID.(float64)", "Error:", ErrNoFloat64Interface, "")
		return
	}

	idDevice, err := h.services.IDevice.CreateDevice(int(intVal), &input)
	if err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"errors": "Ошибка создания устройства",
		})
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
		c.JSON(http.StatusOK, map[string]interface{}{
			"errors": "Ошибка удаления устройства",
		})
		logger.Log("Error", "userID.(float64)", "Error:", ErrNoFloat64Interface, "")
		return 
	}

	err := h.services.IDevice.DeleteDevice(int(intVal), input.Name, input.Home)
	if err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"errors": "Ошибка удаления устройства",
		})
		logger.Log("Error", "DeleteDevice", "Error delete device:", err, id, input.Name)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{})

	logger.Log("Info", "", "A device has been deleted", nil)
}
