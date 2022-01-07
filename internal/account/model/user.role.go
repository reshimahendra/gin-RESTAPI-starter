/*
    Package model for 'Role'
*/
package model

// User Role
type Role struct {
    ID          uint    `gorm:"primaryKey;" json:"id"`
    Name        string  `gorm:"type:varchar(25);" json:"name"`
    Description string  `json:"description"`
}

func (r *Role) TableName() string {
    return "Role"
}
