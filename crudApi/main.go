package main

import (
	"crudApi/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Info().Msg("Starting server")
	routes := gin.Default()

	routes.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome Home")
	})

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}