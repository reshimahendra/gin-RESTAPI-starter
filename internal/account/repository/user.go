/*
   Package repository for 'User'
*/
package repository

import (
	"github.com/reshimahendra/gin-starter/internal/account/model"
	m "github.com/reshimahendra/gin-starter/internal/account/model"
	"github.com/reshimahendra/gin-starter/pkg/logger"
	"gorm.io/gorm"
)

// UserRepository is Interface for User Repository to communicate with the database
type UserRepository interface {
    Get(username string) (user *m.User, err error) 
    GetByEmail(username string) (user *m.User, err error)
    Gets() (users *[]m.User, err error)
    CheckCredential(username string) (hashedPassword string, isActive bool)
    Save(input m.User) (user *m.User, err error)
    Update(username string) (user *m.User, err error)
}

// userRepository is type wrapper for our database instance
type userRepository struct {
    db *gorm.DB
}

// NewUser will create new repository connection to our database
func NewUser(db *gorm.DB) *userRepository {
    return &userRepository{db: db}
}

// Get will fetch 'User' model by 'username'
func (r *userRepository) Get(username string) (user *m.User, err error) {
    var userTmp m.User

    // get user by username
    if err = r.db.Where("username = ?", username).First(&userTmp).Error; err != nil {
        logger.Errorf("Error finding 'User' data: %v", err)
        return nil, err
    }

    // get user role ID
    if userTmp.RoleID >0 {
        var role m.Role
        if err = r.db.Where("id = ?", userTmp.RoleID).First(&role).Error; err == nil {
            userTmp.Role = &role
        }
    }

    return
}

// GetByEmail will fetch 'User' model by their 'email'
func (r *userRepository) GetByEmail(email string) (user *m.User, err error) {
    // get user by email
    if err = r.db.Where("email = ?", email).First(&user).Error; err != nil {
        logger.Errorf("Error finding 'User' data: %v", err)
        return nil, err
    }

    // get user role ID
    if user.RoleID >0 {
        var role model.Role
        if err = r.db.Where("id = ?", user.RoleID).First(&role).Error; err == nil {
            user.Role = &role
        }
    }

    return
}

// Gets will fetch ALL 'user' model 
func (r *userRepository) Gets() (users *[]m.User, err error) {
    if err = r.db.Find(&users).Error; err != nil {
        logger.Errorf("Error getting 'user' data: %v", err)
        return nil, err
    }

    return
}

// CheckCredential will check 'User' credential by its username & password 
// return 
// * password (password of checked user for password comparation operation)
// * active (status whether the user is active or not)
func (r *userRepository) CheckCredential(username string) (hashedPassword string, isActive bool) {
    var u model.User
    if err := r.db.Where("username = ? ", username).First(&u).Error; err != nil {
        logger.Errorf("Error getting 'user' data: %v", err)
        return "", false
    }

    return u.Password, u.Active
}


// Save will save 'User' data to the database
// It will returning 'User' data and 'error' value 
func (r *userRepository) Save(input m.User) (user *m.User, err error) {
    // err = r.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&input).Error
    if err = r.db.Save(&input).Error; err != nil {
        logger.Errorf("Error saving user: %v", err)

        return nil, err 
    }

    if input.RoleID > 0 {
        var role m.Role
        if err = r.db.Where("id = ?", input.RoleID).First(&role).Error; err == nil {
            input.Role = &role
        }
    }

    user = &input

    return 
}

// Update will update user data to the database 
// It will return the 'User' data and 'error' status after the operation
func (r *userRepository) Update(username string) (user *m.User, err error) {
    err = r.db.Where("username = ?", username).First(&user).Error
    if err != nil {
        logger.Errorf("Error updating user: %v", err)
        return nil, err
    }

    if user.RoleID >0 {
        var role m.Role
        if err = r.db.Where("id = ?", user.RoleID).First(&role).Error; err == nil {
            user.Role = &role
        }
    }

    return
}
