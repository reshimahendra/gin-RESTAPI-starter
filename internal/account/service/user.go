/*
    Package service for 'User'
*/
package service

import (
	"github.com/reshimahendra/gin-starter/internal/account/model"
	"github.com/reshimahendra/gin-starter/internal/pkg/helper"
	"github.com/reshimahendra/gin-starter/pkg/logger"
	"gorm.io/gorm"
)

func FindByCredential(email, password string) bool {
    // TODO: below just a sample, you need to create the actual logic here later on
    return email == "a@b.com" && password == "123"
}

/* USER */
// GET 'USER'by ID
func GetUser(db *gorm.DB, id string) (user *model.User, err error) {
    if err = db.Where("id = ?", id).First(&user).Error; err != nil {
        logger.Errorf("Error finding 'User' data: %v", err)
        return nil, err
    }

    // get user role ID
    if user.RoleID >0 {
        var role model.Role
        if err = db.Where("id = ?", user.RoleID).First(&role).Error; err == nil {
            user.Role = &role
        }
    }

    return
}

// GET ALL 'USERS'
func GetUsers(db *gorm.DB) (users *[]model.User, err error) {
    if err = db.Find(&users).Error; err != nil {
        logger.Errorf("Error getting 'user' data: %v", err)
        return nil, err
    }

    return
}

// GET `USER` by Username
func GetUserByUsername(db *gorm.DB, username string) (user *model.User, err error) {
    if err = db.Where("username = ?", username).Error; err != nil {
        logger.Errorf("Error getting 'user' data: %v", err)
        return nil, err
    }

    return
}

// GET `USER` by Email 
func GetUserByEmail(db *gorm.DB, email string) (user *model.User, err error) {
    if err = db.Where("email = ?", email).Error; err != nil {
        logger.Errorf("Error checking user data: %v", err)
        return nil, err
    }

    return
}

// Check `USER` exist by Username OR password
func IsUserExist(db *gorm.DB, email, username string) bool {
    if err := db.Where("email = ? OR username = ?", email, username).Error; err != nil {
        logger.Errorf("Error checking user data: %v", err)
        return false
    }

    return true
}

// GET 'USER' credential by 'USERNAME'
// return 
// * password string
// * active bool
// * registered bool 
func CredentialByUsername(db *gorm.DB, username, password string) (string, bool, bool) {
    var u model.User
    if err := db.Debug().Where("username = ? AND password = ?", username, password).First(&u).Error; err != nil {
        logger.Errorf("Error getting 'user' data: %v", err)
        return "", false, false
    }

    return u.Password, u.Active, true
}


// GET 'USER' credential by 'EMAIL'
// return 
// * password string
// * active bool
// * passwordMatch bool 
func CredentialByEmail(db *gorm.DB, email, password string) (isActive bool, isPasswordMatch bool) {
    
    var u model.User
    if err := db.Debug().Where("email = ? ", email).First(&u).Error; err != nil {
        logger.Errorf("Error getting credential data: %v", err)
        return false, false
    }
    isPasswordMatch = helper.CheckPasswordHash(password, u.Password) 
    isActive = u.Active

    return
}

// POST/ PUT 'USER'
func SaveUser(db *gorm.DB, input model.User) (user *model.User, err error) {
    err = db.Save(&input).Error;
    if err != nil {
        logger.Errorf("Error saving 'user' data: %v", err)
        return nil, err
    }

    if input.RoleID >0 {
        var role model.Role
        if err = db.Where("id = ?", input.RoleID).First(&role).Error; err == nil {
            input.Role = &role
        }
    }
    return &input, err
}

