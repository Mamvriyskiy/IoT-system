package handler

import (
	"net/http"
	"strings"
	"fmt"
	_ "github.com/santosh/gingo/docs"
	"github.com/Mamvriyskiy/database_course/main/logger"
	"github.com/Mamvriyskiy/database_course/main/pkg/service"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"time"
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

func TrafficLoggingMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()

        // Логируем запрос
        zap.L().Info("Incoming request",
            zap.String("method", c.Request.Method),
            zap.String("url", c.Request.URL.String()),
            zap.String("client_ip", c.ClientIP()),
            zap.Any("headers", c.Request.Header),
        )

        // Выполняем запрос
        c.Next()

        // Логируем ответ
        zap.L().Info("Outgoing response",
            zap.Int("status", c.Writer.Status()),
            zap.Duration("latency", time.Since(start)),
            zap.String("client_ip", c.ClientIP()),
        )
    }
}


// @title     Gingo Bookstore API
func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	router.Use(func(ctx *gin.Context) {
        fmt.Println("Requested URL:", ctx.Request.URL.String()) // Логируем URL запроса
        ctx.Next() // Продолжаем обработку запроса
    })
	
	router.Use(TrafficLoggingMiddleware())

	//router.Use(AuthMiddleware())

	router.Static("/css", "./templates/css")
	router.LoadHTMLGlob("templates/*.html")

	router.Static("/docs", "./docs")

	fmt.Println("+++")
    // Настройка Swagger UI
    router.GET("/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
        ginSwagger.URL("/docs/swagger.yaml")))

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

	auth.GET("/verification", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "checkcode.html", nil)
	})

	auth.GET("/password", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "changepswrd.html", nil)
	})

	auth.POST("/sign-up", h.SignUp)
	auth.POST("/sign-in", h.signIn)
	auth.PUT("/password", h.changePassword)
	auth.POST("/verification", h.checkCode)
	auth.POST("/code", h.code)

	api := router.Group("/api")

	home := api.Group("/homes")
	home.POST("/", h.createHome)
	home.GET("/", h.listHome)
	home.DELETE("/:homeID", h.deleteHome)
	home.PUT("/:homeID", h.updateHome)
	home.GET("/:homeID", h.infoHome)

	home.POST("/:homeID/accesses", h.addUser)
	home.DELETE("/:homeID/accesses/:accessID", h.deleteUser)
	home.GET("/:homeID/accesses", h.getListUserHome)
	home.PUT("/:homeID/accesses/:accessID", h.updateLevel)
	home.GET("/:homeID/accesses/:accessID", h.getInfoAccess)

	home.POST("/:homeID/devices", h.createDevice)
	home.GET("/:homeID/devices", h.getListDevice)
	home.DELETE("/:homeID/devices/:deviceID", h.deleteDevice)
	home.GET("/:homeID/devices/:deviceID", h.getInfoDevice)

	home.POST("/:homeID/devices/:deviceID/status", h.createDeviceHistory)
	home.GET("/:homeID/devices/:deviceID/history", h.getDeviceHistory)

	logger.Log("Info", "", "Create router", nil)

	return router
}
