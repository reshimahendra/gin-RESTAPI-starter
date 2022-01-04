/*
    Package model for 'Role'
*/
package model

// User Role
type Role struct {
    ID          uint    `gorm:"primaryKey;autoIncrement" json:"id"`
    Name        string  `gorm:"type:varchar(25);unique;" json:"name"`
    Description string  `json:"description"`
}

func (r *Role) TableName() string {
    return "Role"
}

// RoleRequest is 'DTO' (Data Transfer Object) for 'Role' model
// It will receive 'Role' data and processed (save/update) to database
type RoleRequest struct {
    ID          uint    `json:"id"`
    Name        string  `json:"name" binding:"required"`
    Description string  `json:"description"`
}

// RoleResponse is 'DTO' (Data Transfer Object) for 'Role' model
// It will send 'role' data to client 
type RoleResponse struct {
    ID          uint    `json:"id"`
    Name        string  `json:"name" binding:"required"`
    Description string  `json:"description"`
}
