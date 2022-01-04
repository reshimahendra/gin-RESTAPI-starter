/*
   Package router for 'Account' module
*/
package router

import (
	"github.com/gin-gonic/gin"
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
    // api := a.db
}
