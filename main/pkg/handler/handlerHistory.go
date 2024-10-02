package handler

import (
	"crypto/rand"
	"math/big"
	"fmt"
	"net/http"

	"github.com/Mamvriyskiy/database_course/main/logger"
	"github.com/Mamvriyskiy/database_course/main/pkg"
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
	id, ok := c.Get("userId")
	if !ok {
		logger.Log("Warning", "Get", "Error get userID from context", nil, "userID")
		return
	}

	var input pkg.AddHistory
	if err := c.BindJSON(&input); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	history := pkg.AddHistory{
		Name:             input.Name,
		Home:             input.Home,
	}

	var intVal float64
	if val, ok := id.(float64); ok {
		intVal = val
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"errors": "Ошибка создания истории",
		})
		logger.Log("Error", "userID.(float64)", "Error:", ErrNoFloat64Interface, "")
		return
	}

	idHistory, err := h.services.IHistoryDevice.CreateDeviceHistory(int(intVal), history)
	fmt.Println(idHistory, err)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"errors": "Ошибка создания истории",
		})
		logger.Log("Error", "CreateDeviceHistory", "Error create history:", err, id, history)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"idHistory": idHistory,
	})

	logger.Log("Info", "", "The device's history has been created", nil)
}

type getAllListResponse struct {
	Data []pkg.DevicesHistory `json:"data"`
}

func (h *Handler) getDeviceHistory(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		logger.Log("Warning", "Get", "Error get userID from context", nil, "userID")
		return
	}

	var info pkg.AddHistory
	if err := c.BindJSON(&info); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	var intVal float64
	if val, ok := id.(float64); ok {
		intVal = val
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"errors": "Ошибка получения истории",
		})
		logger.Log("Error", "userID.(float64)", "Error:", ErrNoFloat64Interface, "")
		return
	}

	input, err := h.services.IHistoryDevice.GetDeviceHistory(int(intVal), info.Name, info.Home)
	if err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"errors": "Ошибка получения истории",
		})
		logger.Log("Error", "GetDeviceHistory", "Error get history:", err, id, info.Name)
		return
	}

	c.JSON(http.StatusOK, getAllListResponse{
		Data: input,
	})

	logger.Log("Info", "", "The history of the device was obtained", nil)
}
