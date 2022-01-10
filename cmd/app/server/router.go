/*
   Package server
   Setting up main router from various app
*/
package server

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/reshimahendra/gin-starter/internal/account"
	"github.com/reshimahendra/gin-starter/pkg/logger"
	"github.com/reshimahendra/gin-starter/pkg/middleware"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// Router will initialize and run registered app routers
func Router(db *gorm.DB) *gin.Engine {
    var router *gin.Engine

    // Make router instance
    switch viper.GetString("server.mode") {
    case "production":
        // Using release mode setting
        gin.SetMode(gin.ReleaseMode)
        router = gin.New()

        // SET Middleware
        // gin default Middleware
        router.Use(gin.Logger())
        router.Use(gin.Recovery())
    case "development": 
        router = gin.Default()
    default:
        router = gin.Default()
    }

    // Enable the server access log
    accessLog := viper.GetString("logger.access_log_name")
    if accessLog != "" {
        accessLog = "./log/.access.log"
    }

    // Preparing up our access log
    file, err := os.OpenFile(accessLog, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        logger.Errorf("Failed to create access log file: %v", err)
    } else {
        gin.DefaultWriter = io.MultiWriter(file)
    }

    // custom Middleware
    router.Use(middleware.CORS())
    router.Use(middleware.Security())

    // Run Account module router 
    account := account.New(router, db)
    account.Router()

    return router
}
