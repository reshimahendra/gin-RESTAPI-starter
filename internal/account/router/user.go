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
	"github.com/reshimahendra/gin-starter/pkg/middleware"
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
func (u *user) Run() {
    // Auto Migrate all model available on user module 
    u.db.AutoMigrate(
        &model.User{},
        &model.Role{},
    )
    // create repository, service, and handler instance for user model
    repoUser    := repository.NewUser(u.db)
    serviceUser := service.NewUser(repoUser)
    apiUser     := handler.NewUser(serviceUser)


    // create repository, service, and handler instance for user model
    repoRole        := repository.NewUserRole(u.db)
    serviceRole     := service.NewUserRole(repoRole)
    apiUserRole     := handler.NewUserRole(serviceRole)

    // Group for Auth
    // Api for AUTH using 'User' service package
    rAuth := u.router.Group("/auth")
    rAuth.POST("/sign-up", apiUser.Signup)
    rAuth.POST("/sign-in", apiUser.Signin)

    // Protected area for AUTH endpoint
    rAuth.Use(middleware.Authorize())
    {
        rAuth.POST("/refresh-token", apiUser.RefreshToken)
        rAuth.POST("/check-token", apiUser.CheckToken)
    }

    // Group for 'User' module
    r := u.router.Group("/user")

    // Non Protected area
    r.GET("/username/:user", apiUser.Get)
    r.GET("/email/:email", apiUser.GetByEmail)
    r.GET("/", apiUser.Gets)
    r.PUT("/:username", apiUser.Update)
    r.POST("/", apiUser.Create)
    
    // Protected area
    // Create group router for 'Role' under 'User' group router
    uRole := r.Group("/role")
    uRole.GET("/:id", apiUserRole.Get)
    uRole.GET("/", apiUserRole.Gets)
    uRole.PUT("/:id", apiUserRole.Update)
    uRole.POST("/", apiUserRole.Create)
}


