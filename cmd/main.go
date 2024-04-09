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
)

var (
	KO_RES      = gin.H{"status": "KO"}
	OK_RES      = gin.H{"status": "OK"}
	ErrBadParam = errors.New("bad_parameter")
	ErrSvc      = errors.New("service_error")
)

type EmailRequestBody struct {
	Email string
}

func main() {
	l, svc := new()

	route(l, svc)
}

func new() (*util.Logger, *internal.Service) {
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

	return l, svc
}

func route(l *util.Logger, svc *internal.Service) {
	router := gin.Default()

	router.POST("/email", func(c *gin.Context) {
		var requestBody EmailRequestBody

		err := c.BindJSON(&requestBody)
		if err != nil {
			l.ErrorLogger.Println(err)
			c.JSON(http.StatusBadRequest, KO_RES)
			return
		}

		isValid := util.IsEmailValid(requestBody.Email)
		if !isValid {
			l.ErrorLogger.Printf("Invalid eMail: %s", requestBody.Email)
			c.JSON(http.StatusBadRequest, fmt.Errorf("%v: %w", err, ErrBadParam))
			return
		}

		err = svc.SendToQueue(requestBody.Email)
		if err != nil {
			l.ErrorLogger.Println(err.Error())
			c.JSON(http.StatusInternalServerError, KO_RES)
		} else {
			l.InfoLogger.Println("Successfully sent!")
			c.JSON(http.StatusAccepted, OK_RES)
		}
	})

	router.Run()
}
