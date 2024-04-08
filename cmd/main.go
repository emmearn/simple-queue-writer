package main

import (
	"log"
	"net/http"
	"simple-queue-writer/internal/config"
	"simple-queue-writer/internal/util"

	"github.com/gin-gonic/gin"
)

var (
	KO_RES = gin.H{"status": "KO"}
	OK_RES = gin.H{"status": "OK"}
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

	router := gin.Default()

	router.POST("/phone", func(c *gin.Context) {
		var desc string

		err := c.BindJSON(&desc)
		if err != nil {
			l.ErrorLogger.Println(err.Error())
			c.JSON(http.StatusBadRequest, KO_RES)
		}

		err = svc.CreateCommission(desc)
		if err != nil {
			l.ErrorLogger.Println(err.Error())
			c.JSON(http.StatusInternalServerError, KO_RES)
		} else {
			c.JSON(http.StatusAccepted, OK_RES)
		}
	})

	router.Run("localhost:" + cfg.Server.Port)
}
