package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"fmt"
	"net/http"
	"os"
)

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.UseRawPath = true
	router.UnescapePathValues = false

	router.Use(gin.RecoveryWithWriter(os.Stdout))
	router.Use(gin.LoggerWithWriter(os.Stdout))

	//router.GET("/metrics", prometheusHandler())

	router.GET("/bench", func(ctx *gin.Context) {
		var res [][]string
		for i := 0; i < 10; i++ {
			fmt.Println("ITERATION ", i)
			res2 := benchmark.ClientBench()
			res = append(res, res2)
		}

		ctx.JSON(http.StatusOK, res)
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

