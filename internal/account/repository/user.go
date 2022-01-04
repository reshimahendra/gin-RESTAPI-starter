/*
   Package repository for 'User'
*/
package repository

import m "github.com/reshimahendra/gin-starter/internal/account/model"

type userRepository interface {
    GetUser(username string) (user *m.UserResponse, err error) 
    GetEmail(username string) (user *m.UserResponse, err error)
    Gets() (users *[]m.UserResponse, err error)
    CheckCredential(username, password string) (hassPassword string, isActive, isValid bool)
}
