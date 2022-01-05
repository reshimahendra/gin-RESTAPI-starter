/*
   Package repository, Main repository
*/
package repository

import "gorm.io/gorm"

// Repository is struct with Gorm DB reference
type repository struct {
    db *gorm.DB
}

// New will create new connection to the database for our main repository  
func New(db *gorm.DB) *repository {
    return &repository{db: db}
}

