/*
   Package router for 'Account' module
*/
package account

import (
	"github.com/gin-gonic/gin"
	"github.com/reshimahendra/gin-starter/internal/account/router"
	"gorm.io/gorm"
)

type account struct {
    router *gin.Engine
    db *gorm.DB
}

// New will create new 'gin' Engine and 'gorm' database instance
func New(r *gin.Engine, db *gorm.DB) *account{
    return &account{
        router : r,
        db     : db,
    }
}

// Router() is the main router for 'Account' module
func (a *account) Router() {
    // Create user router instance and run it 
    user := router.NewUser(a.router, a.db)
    user.Run()
}
