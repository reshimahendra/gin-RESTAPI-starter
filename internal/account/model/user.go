/*
   Package model for 'User'
*/
package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/reshimahendra/gin-starter/internal/database/model"
	"gorm.io/gorm"
)

// User is a struct for 'User' model
type User struct {
    model.BaseUUID
    Username string `gorm:"type:varchar(30);not null;unique" json:"username"`
    Firstname string `gorm:"type:varchar(30);not null" json:"first_name"`
    Lastname string `gorm:"type:varchar(30)" json:"last_name"`
    Email string `gorm:"type:varchar(100);not null;unique;" json:"email"`
    Password string `gorm:"not null" json:"password"`
    Active bool `gorm:"default:false" json:"active"`
    RoleID uint `gorm:"type:int" json:"role_id"`
    Role *Role `gorm:"foreignKey:RoleID" json:"role"` 
}

//  Table name for User 
func (u *User) TableName() string {
    return "User"
}

// BeforeCreate is hook for 'User' model 'Before Create' operation
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
    u.ID = uuid.New()
    u.CreatedAt = time.Now()
    u.UpdatedAt = time.Now()
    return
}

// BeforeUpdate is hook for 'User' model 'Before Update' operation
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
    u.UpdatedAt = time.Now()
    return
} 
