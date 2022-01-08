/*
   Package repository, User Role Repository
*/
package repository

import (
	"github.com/reshimahendra/gin-starter/internal/account/model"
	E "github.com/reshimahendra/gin-starter/pkg/errors"
	"gorm.io/gorm"
)

// UserRoleRepository is CRUD interface for 'Role' model
type UserRoleRepository interface {
    Get(id uint) (role *model.Role, err error)
    Gets() (roles *[]model.Role, err error)
    Create(input model.Role) (role *model.Role, err error)
    Update(id uint, input model.Role) (role *model.Role, err error)
}

// Repository is struct with Gorm DB reference
type userRoleRepository struct {
    db *gorm.DB
}

// New will create new connection to the database for our main repository  
func NewUserRole(db *gorm.DB) *userRoleRepository {
    return &userRoleRepository{db: db}
}

// Get will get 'Role' data based on given param 'id'
func (r *userRoleRepository) Get(id uint) (role *model.Role, err error) {
    err = r.db.Where("id", id).First(&role).Error
    if err != nil {
        return nil, err
    }
    return
}

// Gets will get all 'Role' data 
func (r *userRoleRepository) Gets() (roles *[]model.Role, err error){
    err = r.db.Find(&roles).Error
    if err != nil {
        return nil, err
    }
    return
}

// Save will save 'Role' data based on user 'input request'
func (r *userRoleRepository) Create(input model.Role) (role *model.Role, err error){
    err = r.db.Create(&input).Error
    if err != nil {
        return nil, err
    }
    role = &input

    return
}

// Update will update 'Role' data based on param 'id' and 'input request'
func (r *userRoleRepository) Update(id uint, input model.Role) (role *model.Role, err error) {
    result := r.db.Where("id = ?", id).Updates(&input)
    err = result.Error

    if err != nil {
        return nil, err
    }

    // no data updated applied, send error response to notify client
    if result.RowsAffected == 0 {
        err = E.NewSimpleError(E.ErrUpdateDataFail)
        return nil, err
    }

    input.ID = id
    role = &input

    return
}
