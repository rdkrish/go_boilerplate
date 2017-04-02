package server

import (
	"go_boilerplate/config"
	"github.com/gin-gonic/gin"
)

// Init function to initialize the service
func Init() {
	config := config.GetConfig()
	gin.SetMode(config.GetString("GIN_MODE"))
	r := NewRouter()
	r.Run(config.GetString("SERVER"))
}
