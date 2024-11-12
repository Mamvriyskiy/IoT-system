package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"fmt"
	bench "github.com/Mamvriyskiy/database_course/main/benchmark/testsBench"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"os"
)

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.UseRawPath = true
	router.UnescapePathValues = false

	router.Use(gin.RecoveryWithWriter(os.Stdout))
	router.Use(gin.LoggerWithWriter(os.Stdout))

	//router.GET("/metrics", prometheusHandler())

	// os.Setenv("TESTCONTAINERS_RYUK_DISABLED", "true")
	// os.Setenv("TESTCONTAINERS_DEBUG", "true")

	router.GET("/metrics", prometheusHandler())
	router.GET("/bench", func(ctx *gin.Context) {
		var res [][]string
		for i := 0; i < 101; i++ {
			fmt.Println("ITERATION ", i)
			res2 := bench.ClientBench()
			res = append(res, res2)
		}

		ctx.IndentedJSON(http.StatusOK, res)
	})

	s := http.Server{
		Addr:    fmt.Sprintf(":8081"),
		Handler: router,
	}

	err := s.ListenAndServe()
	if err != nil && !errors.Is(http.ErrServerClosed, err) {
		panic(err)
	}
}

