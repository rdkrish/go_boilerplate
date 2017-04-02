package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go_boilerplate/logger"
)

type ServiceController struct{}

func(s ServiceController) Get (c *gin.Context) {
	response := gin.H{
		"message":  "Get call to service",
		"status": "success",
	}
	logger.Log.Info("Get request received")
	c.JSON(http.StatusOK, response)
}

func(s ServiceController) Post (c *gin.Context) {
	response := gin.H{
		"message":  "Post call to service",
		"status": "success",
	}
	c.JSON(http.StatusOK, response)
}

func(s ServiceController) Delete (c *gin.Context) {
	response := gin.H{
		"message":  "Delete call to service",
		"status": "success",
	}
	c.JSON(http.StatusOK, response)
}
