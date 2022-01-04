/*
   Package model for 'User'
*/
package model

import (
	"time"

	"github.com/reshimahendra/gin-starter/internal/database/model"
	"gorm.io/gorm"
)

// User is a struct for 'User' model
type User struct {
    model.Base
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
    u.CreatedAt = time.Now()
    u.UpdatedAt = time.Now()
    return
}

// BeforeUpdate is hook for 'User' model 'Before Update' operation
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
    u.UpdatedAt = time.Now()
    return
} 


// UserRequest is 'DTO' (Data Transfer Object) for 'User' request 
// It will receive 'User' data and processed (save/update) to database
type UserRequest struct {
    Username  string        `json:"username" binding:"required"`
    Firstname string        `json:"first_name" binding:"required"`
    Lastname  string        `json:"last_name"`
    Email     string        `json:"email" binding:"required"`
    Password  string        `json:"password" binding:"required"`
    Active    bool          `json:"active"`
    RoleID    uint          `json:"role_id"`
    Role      *Role         `json:"role"`
}

// UserResponse is 'DTO' (Data Transfer Object) for 'User' model 
// It will send 'User' data to client 
type UserResponse struct {
    ID        uint64        `json:"id"`
    Username  string        `json:"username"`
    Firstname string        `json:"first_name"`
    Lastname  string        `json:"last_name"`
    Email     string        `json:"email"`
    Password  string        `json:"password"`
    Active    bool          `json:"active"`
    RoleID    uint          `json:"role_id"`
    Role      *Role         `json:"role"`
}

// Credential is 'DTO' (Data Transfer Object) for user response 
// This is used for login/ credential checking or other similar operation
type Credential struct {
    Username  string `json:"username"`
    Email     string `json:"email"`
    Password  string `json:"password"`
    Active    bool   `json:"active"`
}
