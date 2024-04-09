package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"simple-queue-writer/internal"
	"simple-queue-writer/internal/config"
	"simple-queue-writer/internal/util"

	"github.com/gin-gonic/gin"

	"golang.org/x/exp/maps"
)

var (
	KO_RES = gin.H{"status": "KO"}
	OK_RES = gin.H{"status": "OK"}
	ErrSvc = errors.New("service_error")
)

type EmailRequestBody struct {
	Email string
}

func main() {
	l, cfg, svc := new()

	route(l, cfg, svc)
}

func new() (*util.Logger, *config.Config, *internal.Service) {
	cfg, err := config.ConfigLoad()
	if err != nil {
		log.Fatal(err)
	}

	l, err := util.NewLogger()
	if err != nil {
		log.Fatal(err)
	}

	svc, err := internal.NewService(l, cfg)
	if err != nil {
		log.Fatal(fmt.Errorf("%v: %w", err, ErrSvc))
	}

	return l, cfg, svc
}

func route(l *util.Logger, cfg *config.Config, svc *internal.Service) {
	router := gin.Default()

	router.POST("/email", func(c *gin.Context) {
		var requestBody EmailRequestBody

		err := c.BindJSON(&requestBody)
		if err != nil {
			l.ErrorLogger.Println(err)
			c.JSON(http.StatusBadRequest, KO_RES)
			return
		}

		err = svc.SendToQueue(requestBody.Email)
		if err != nil {
			if errors.Is(err, internal.ErrBadParam) {
				errMap := map[string]any{
					"error": err.Error(),
				}
				maps.Copy(errMap, KO_RES)
				c.JSON(http.StatusBadRequest, errMap)
			} else {
				c.JSON(http.StatusInternalServerError, KO_RES)
			}
		} else {
			c.JSON(http.StatusOK, OK_RES)
		}
	})

	router.Run(":" + cfg.Server.Port)
}
