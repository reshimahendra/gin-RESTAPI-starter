/*
   Package repository for 'User'
*/
package repository

import (
	m "github.com/reshimahendra/gin-starter/internal/account/model"
	"gorm.io/gorm"
)

// UserRepository is Interface for User Repository to communicate with the database
type UserRepository interface {
    Get(username string) (user *m.User, err error) 
    GetByEmail(username string) (user *m.User, err error)
    Gets() (users *[]m.User, err error)
    Create(input m.User) (user *m.User, err error)
    Update(username string, input m.User) (user *m.User, err error)

    /* Functional interface */
    CheckCredential(username string) (hashedPassword string, isActive bool)
    CheckCredentialByMail(email string) (hashedPassword string, isActive bool)
    UserNotFound(username string) (isUserNotFound bool)
    UserAvailable(username, email string) (isUserAvailable bool)
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
    // get user by username
    if err = r.db.Where("username = ?", username).First(&user).Error; err != nil {
        return nil, err
    }

    // get user role ID
    if user.RoleID >0 {
        var role m.Role
        if err = r.db.Where("id = ?", user.RoleID).First(&role).Error; err == nil {
            user.Role = &role
        }
    }

    return
}

// GetByEmail will fetch 'User' model by their 'email'
func (r *userRepository) GetByEmail(email string) (user *m.User, err error) {
    // get user by email
    if err = r.db.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }

    // get user role ID
    if user.RoleID >0 {
        var role m.Role
        if err = r.db.Where("id = ?", user.RoleID).First(&role).Error; err == nil {
            user.Role = &role
        }
    }

    return
}

// Gets will fetch ALL 'user' model 
func (r *userRepository) Gets() (users *[]m.User, err error) {
    if err = r.db.Find(&users).Error; err != nil {
        return nil, err
    }

    return
}

// Create will insert new 'User' record with it associates(if available) to the database
// It will returning 'User' data and 'error' value 
func (r *userRepository) Create(input m.User) (user *m.User, err error) {
    // err = r.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&input).Error
    if err = r.db.Create(&input).Error; err != nil {
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
func (r *userRepository) Update(username string, input m.User) (user *m.User, err error) {
    // Update user data while checking if some error occur
    if err = r.db.Where("username = ?", username).Updates(&input).Error; err != nil {
        return
    }

    
    // get user detail
    userTmp, err := r.Get(username)
    if err != nil {
        return
    }

    // pass new model struct (userTmp) to 'user' model output
    user = userTmp
   

    return
}

// CheckCredential will check 'User' credential by its username & password 
// return 
// * password (password of checked user for password comparation operation)
// * active (status whether the user is active or not)
func (r *userRepository) CheckCredential(username string) (hashedPassword string, isActive bool) {
    var u m.User
    if err := r.db.Where("username = ? ", username).First(&u).Error; err != nil {
        return "", false
    }

    return u.Password, u.Active
}

// CheckCredentialByMail will check 'User' credential by its email & password 
// return 
// * password (password of checked user for password comparation operation)
// * active (status whether the user is active or not)
func (r *userRepository) CheckCredentialByMail(email string) (hashedPassword string, isActive bool) {
    var u m.User
    if err := r.db.Where("email = ? ", email).First(&u).Error; err != nil {
        return "", false
    }

    hashedPassword = u.Password
    isActive = u.Active

    return
}

// UserNotFound will check user whether user founded or not.
func (r *userRepository) UserNotFound(username string) (isUserNotFound bool) {
    var tmpUser m.User
    isUserNotFound = r.db.Where("username = ?", username).First(&tmpUser).Error == gorm.ErrRecordNotFound

    return
}

// UserAvailable will check whether username or email is available
func (r *userRepository) UserAvailable(username, email string) (isUserAvailable bool) {
    var tmpUser m.User
    isUserAvailable = r.db.Where("username = ? OR email = ?", username, email).
        First(&tmpUser).Error == gorm.ErrRecordNotFound 

    return
}
