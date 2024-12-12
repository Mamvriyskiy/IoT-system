package handler

import (
	"errors"
	"net/http"
	"fmt"

	"github.com/Mamvriyskiy/database_course/main/logger"
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/gin-gonic/gin"
)

var ErrNoFloat64Interface = errors.New("отсутствует интерфейс {} для float64")

func (h *Handler) getListDevice(c *gin.Context) {
	homeID := c.Param("homeID")

	listDevices, err := h.services.IDevice.GetListDevices(homeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"errors": "Ошибка получения списка пользователей",
		})
		logger.Log("Error", "GetListDevices", "Error get list devices:", err, homeID)
		return
	}

	if listDevices == nil || len(listDevices) == 0 {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"errors": "список устройств пуст",
		})
		return 
	}

	logger.Log("Info", "ListDev:", "Error get list devices:", nil, listDevices, len(listDevices))

	c.JSON(http.StatusOK, listDevices)

	logger.Log("Info", "", "The history of the device was obtained", nil)
}

func (h *Handler) createDevice(c *gin.Context) {
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
	fmt.Println("===", homeID, "===")
	accessLevel, err := h.services.IUser.GetAccessLevel(userID, homeID)
	if accessLevel != 4 || err != nil {
		c.JSON(http.StatusForbidden, map[string]string{
			"errors": "Недостаточно прав для удаления",
		})
		logger.Log("Error", "GetAccessLevel", "Error GetAccessLevel home:", err, accessLevel)
		return
	}

	var input pkg.DevicesHandler
	if err := c.BindJSON(&input); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	device, err := h.services.IDevice.CreateDevice(homeID, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"errors": "Ошибка создания устройства",
		})
		logger.Log("Error", "CreateDevice", "Error create device:", err, homeID, &input)
		return
	}

	c.JSON(http.StatusOK, device)

	logger.Log("Info", "", "A device has been created", nil)
}

func (h *Handler) deleteDevice(c *gin.Context) {
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

	deviceID := c.Param("deviceID")

	err = h.services.IDevice.DeleteDevice(deviceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"errors": "Ошибка удаления устройства",
		})
		logger.Log("Error", "DeleteDevice", "Error delete device:", err, deviceID)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{})

	logger.Log("Info", "", "A device has been deleted", nil)
}

func (h *Handler) getInfoDevice(c *gin.Context) {
	deviceID := c.Param("deviceID")

	device, err := h.services.IDevice.GetInfoDevice(deviceID)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"errors": "Устройство не существует",
		})
		logger.Log("Error", "DeleteDevice", "Error delete device:", err, deviceID)
		return
	}

	c.JSON(http.StatusOK, device)

	logger.Log("Info", "", "A device has been deleted", nil)
}
