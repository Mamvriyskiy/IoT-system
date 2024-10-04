package handler

import (
	"crypto/rand"
	"math/big"
	"net/http"

	"github.com/Mamvriyskiy/database_course/main/logger"

	"github.com/gin-gonic/gin"
)

func generateRandomInt(max int) int {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		logger.Log("Error", "userID.(float64)", "Error:", ErrNoFloat64Interface, "")
		return 0
	}
	return int(n.Int64())
}

func generateRandomFloat(max float64) float64 {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max*100)))
	if err != nil {
		logger.Log("Error", "userID.(float64)", "Error:", ErrNoFloat64Interface, "")
		return 0.0
	}
	return float64(n.Int64()) / 100.0
}

func (h *Handler) createDeviceHistory(c *gin.Context) {
	deviceID := c.Param("deviceID")

	historyID, err := h.services.IHistoryDevice.CreateDeviceHistory(deviceID)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"errors": "Ошибка создания истории",
		})
		logger.Log("Error", "CreateDeviceHistory", "Error create history:", err, deviceID)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"historyID": historyID,
	})

	logger.Log("Info", "", "The device's history has been created", nil)
}

func (h *Handler) getDeviceHistory(c *gin.Context) {
	deviceID := c.Param("deviceID")

	historyList, err := h.services.IHistoryDevice.GetDeviceHistory(deviceID)
	if err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"errors": "Ошибка получения истории",
		})
		logger.Log("Error", "GetDeviceHistory", "Error get history:", err, deviceID)
		return
	}

	c.JSON(http.StatusOK, historyList)

	logger.Log("Info", "", "The history of the device was obtained", nil)
}
