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

func main() {
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

	router := gin.Default()

	router.POST("/email", func(c *gin.Context) {
		var email string

		err := c.BindJSON(&email)
		if err != nil {
			l.ErrorLogger.Println(err.Error())
			c.JSON(http.StatusBadRequest, KO_RES)
		}

		isValid := util.IsEmailValid(email)
		if !isValid {
			l.ErrorLogger.Println("Invalid eMail!")
			c.JSON(http.StatusBadRequest, fmt.Errorf("%v: %w", err, ErrBadParam))
		}

		err = svc.SendToQueue(email)
		if err != nil {
			l.ErrorLogger.Println(err.Error())
			c.JSON(http.StatusInternalServerError, KO_RES)
		} else {
			l.InfoLogger.Println("Successfully sent!")
			c.JSON(http.StatusAccepted, OK_RES)
		}
	})

	router.Run()
	// router.Run("localhost:" + cfg.Server.Port)
}
