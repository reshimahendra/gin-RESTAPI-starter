/*
   Package router for 'User'
*/
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/reshimahendra/gin-starter/internal/account/handler"
	"github.com/reshimahendra/gin-starter/internal/account/model"
	"github.com/reshimahendra/gin-starter/internal/account/repository"
	"github.com/reshimahendra/gin-starter/internal/account/service"
	"gorm.io/gorm"
)

// user is type wrapper for gin engine and gorm database
type user struct {
    router *gin.Engine
    db *gorm.DB
}

// New will create new 'gin' Engine and 'gorm' database instance
func NewUser(r *gin.Engine, db *gorm.DB) *user{
    return &user{
        router : r,
        db     : db,
    }
}

// Run will execute router for 'User' model 
func (a *user) Run() {
    // Auto Migrate all model available on user module 
    a.db.AutoMigrate(
        &model.User{},
        &model.Role{},
    )
    // Group for 'User' module
    r := a.router.Group("/user")
    
    // create repository, service, and handler instance for user model
    repo := repository.NewUser(a.db)
    service := service.NewUser(repo)
    api := handler.NewUser(service)

    // Protected area

    // Non Protected area
    r.GET("/:user", api.Get)
    r.POST("/", api.Save)
}


