/*
    Package model for 'Role'
*/
package model

// User Role
type Role struct {
    ID          uint    `gorm:"primaryKey;autoincrement;" json:"id"`
    Name        string  `gorm:"unique;type:varchar(25);" json:"name"`
    Description string  `json:"description"`
}

func (r *Role) TableName() string {
    return "Role"
}
