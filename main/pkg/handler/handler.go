package handler

import (
	"net/http"
	"strings"

	"github.com/Mamvriyskiy/DBCourse/main/logger"
	"github.com/Mamvriyskiy/DBCourse/main/pkg/service"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const signingKey = "jaskljfkdfndnznmckmdkaf3124kfdlsf"

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services: services}
}

// Middleware для извлечения данных из JWT и добавления их в контекст запроса.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Проверить URL запроса
		if !strings.HasPrefix(c.Request.URL.Path, "/api") {
			// Если URL неначинается с /api, пропустить проверку JWT
			c.Next()
			return
		}

		// Получить токен из заголовка запроса или из куки
		tokenString := c.GetHeader("Authorization")
		var err error
		if tokenString == "" {
			// Если токен не найден в заголовке, попробуйте из куки
			tokenString, err = c.Cookie("jwt")
			if err != nil {
				logger.Log("Error", "c.Cookie(jwt)", "Error", err, "jwt")
			}
		}

		// Проверить, что токен не пустой
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Empty token"})
			c.Abort()
			return
		}

		// Парсинг токена
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Здесь нужно вернуть ключ для проверки подписи токена.
			// В реальном приложении, возможно, это будет случайный секретный ключ.
			return []byte(signingKey), nil
		})
		// Проверить наличие ошибок при парсинге токена
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "detail:": err.Error()})
			c.Abort()
			return
		}

		// Добавить данные из токена в контекст запроса
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userId", claims["userId"])
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
	}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	router.Use(AuthMiddleware())

	router.Static("/css", "./templates/css")
	router.LoadHTMLGlob("templates/*.html")

	app := router.Group("/app")
	app.GET("/menu", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "menu.html", nil)
	})

	app.GET("/home", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", nil)
	})

	app.GET("/access", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "access.html", nil)
	})

	app.GET("/device", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "device.html", nil)
	})

	auth := router.Group("/auth")
	auth.GET("/sign-up", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "registr.html", nil)
	})

	auth.GET("/sign-in", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "auth.html", nil)
	})

	auth.GET("/reset-password", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "send.html", nil)
	})

	auth.GET("/checkcode", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "checkcode.html", nil)
	})

	auth.GET("/newpassword", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "changepswrd.html", nil)
	})

	auth.POST("/sign-up", h.signUp)
	auth.POST("/sign-in", h.signIn)
	auth.POST("/newpassword", h.changePassword)
	auth.POST("/checkcode", h.checkCode)
	auth.POST("/sendcode", h.sendCode)

	api := router.Group("/api")

	home := api.Group("/home")
	home.POST("/create", h.createHome)
	home.DELETE("/delete", h.deleteHome)
	home.PUT("/update", h.updateHome)
	home.GET("/list", h.listHome)

	access := api.Group("/access")
	access.POST("/add", h.addUser)
	access.DELETE("/delete", h.deleteUser)
	access.GET("/getList", h.getListUserHome)
	access.PUT("/level/", h.updateLevel)
	access.PUT("/status/", h.updateStatus)

	devices := api.Group("/device")
	devices.POST("/", h.createDevice)
	devices.PUT("/", h.getListDevice)
	devices.DELETE("/", h.deleteDevice)

	deviceHistory := api.Group("/history")
	deviceHistory.POST("/", h.createDeviceHistory)
	deviceHistory.PUT("/", h.getDeviceHistory)

	logger.Log("Info", "", "Create router", nil)

	return router
}
