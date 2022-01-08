/*
   Package service for 'User'
   It will implementing interfaces available on 'user repository'
   and create bridge for the 'repository' package and 'handler' package
*/
package service

import (
	"github.com/reshimahendra/gin-starter/internal/account/repository"
	"github.com/reshimahendra/gin-starter/internal/pkg/helper"
	E "github.com/reshimahendra/gin-starter/pkg/errors"
)

// UserService is Interface for User Repository with our Handler
type UserService interface {
    Get(username string) (user *UserResponse, err error) 
    GetByEmail(username string) (user *UserResponse, err error)
    Gets() (users *[]UserResponse, err error)
    Create(input UserRequest) (user *UserResponse, err error)
    Update(username string, input UserRequest) (user *UserResponse, err error)
    CheckCredential(username, password string) (isActive, isValid bool)
    CheckCredentialByMail(email, password string) (isActive, isValid bool)
    UserNotFound(username string) (isUserNotFound bool)
    UserAvailable(username, email string) (isUserAvailable bool)
}


// userService is a type wrapper for 'UserRepository' interface 
type userService struct {
    repo repository.UserRepository
}

// NewUser will return userService instance to connect the service and repository
func NewUser(repo repository.UserRepository) *userService {
    return &userService{repo: repo}
}

// Get will retreive User 'DTO' (by username) that prepared by user repository
// It will returning *UserResponse DTO and error status
func (s *userService) Get(username string) (user *UserResponse, err error) {
    userTmp, err := s.repo.Get(username)

    // Conver 'User' model to 'DTO' fromat
    user = UserToResponse(*userTmp)

    return 
}

// GetByEmail will fetch 'User' data based on given 'email' value
// It will returning *UserResponse DTO and error status
func (s *userService) GetByEmail(email string) (user *UserResponse, err error){
    userTmp, err := s.repo.GetByEmail(email)

    user = UserToResponse(*userTmp)
    return
}

// Gets will fetch all 'User' data and returning *[]UserResponse DTO
func (s *userService) Gets() (users *[]UserResponse, err error){
    userTmp, err := s.repo.Gets()

    var resTemp []UserResponse
    for _, user := range *userTmp {
        resTemp = append(resTemp, *UserToResponse(user)) 
    }
    
    users = &resTemp

    return
}

// Create will convert DTO to saveable format before passed to 'user repository'
// It will returning *UserResponse DTO and error status
func (s *userService) Create(input UserRequest) (user *UserResponse, err error){
    // check if password length is less than minimum required password length
    // exit process if the password too short
    if helper.PasswordTooShort(input.Password) { 
        err = E.NewSimpleError(E.ErrPasswordTooShort)
        return
    }

    // create hashed password for the user. if error, exit the process 
    hashPass, err := helper.HashPassword(input.Password)
    if err != nil {
        return
    }

    // convert from 'DTO' to 'User' model so we can process it to the database
    input.Password = hashPass
    inputUser := RequestToUser(input)

    // perform save operation
    savedUser, err := s.repo.Create(*inputUser)

    // convert back to 'response' DTO before sending back to the user
    if err == nil {
        user = UserToResponse(*savedUser)
    }

    return
}

// Update will send update data request to repository to update certain user
// It returning *UserResponse and error status
func (s *userService) Update(username string, input UserRequest) (user *UserResponse, err error) {
    // check whether username is valid/ registered as well as get hashedPassword
    // and active status for data comparison
    hashedPassword, isActive := s.repo.CheckCredential(username)
    if len(hashedPassword) == 0 && !isActive {
        err = E.NewSimpleError(E.ErrDataNotFound)

        return
    }

    // check if new password length is less than minimum required password length
    // if the password too short, exit the process 
    if helper.PasswordTooShort(input.Password) { 
        err = E.NewSimpleError(E.ErrPasswordTooShort)

        return 
    }

    // check whether on the new user data the password has been changed
    if !helper.CheckPasswordHash(input.Password, hashedPassword) {
        // if changed, create new hashed password
        hashedPassword, err = helper.HashPassword(input.Password)
        if err != nil {
            return
        }
    }
    input.Password = hashedPassword

    // convert request DTO to savable model format 
    inputUser := RequestToUser(input)

    // Send request to update user data to the repository 
    userTmp, err := s.repo.Update(username, *inputUser)
    if err != nil {
        return
    }

    // Convert 'User' data to response 'DTO' before forward it to the client
    user = UserToResponse(*userTmp)

    return
}

// CheckCredential will check wheter the given 'username' and 'password' is valid 
// and the 'checked' user status is active 
func (s *userService) CheckCredential(username, password string) (isActive, isValid bool){
    hashedPassword, isActive := s.repo.CheckCredential(username)
    isValid = helper.CheckPasswordHash(password, hashedPassword)

    return
}

// CheckCredentialByMail will check wheter the given 'email' and 'password' is valid 
// and the 'checked' user status is active 
func (s *userService) CheckCredentialByMail(email, password string) (isActive, isValid bool) {
    hashedPassword, isActive := s.repo.CheckCredentialByMail(email)
    isValid = helper.CheckPasswordHash(password, hashedPassword)

    return
}


// UserNotFound will send request to repository to check whether user data is found or not
func (s *userService) UserNotFound(username string) (isUserNotFound bool) {
    return s.repo.UserNotFound(username)
}

// UserAvailable will check whether username or email is available
func (s *userService) UserAvailable(username, email string) (isUserAvailable bool) {
    return s.repo.UserAvailable(username, email)
}
