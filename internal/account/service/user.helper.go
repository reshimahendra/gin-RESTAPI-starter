/*
   Package 'service' helper for 'User' model
   * DTO to Model conversion
   * Model to DTO conversion
*/
package service

import (
	"github.com/google/uuid"
	"github.com/reshimahendra/gin-starter/internal/account/model"
)

// UserRequest is 'DTO' (Data Transfer Object) for 'User' request
// It will receive 'User' data and processed (save/update) to database
type UserRequest struct {
    Username  string        `json:"username" binding:"required"`
    Firstname string        `json:"first_name" binding:"required"`
    Lastname  string        `json:"last_name"`
    Email     string        `json:"email" binding:"required"`
    Password  string        `json:"password" binding:"required"`
    Active    bool          `json:"active,default=false"`
    RoleID    uint          `json:"role_id" binding:"required"`
    Role      *RoleRequest  `json:"role"`
}

// UserResponse is 'DTO' (Data Transfer Object) for 'User' model 
// It will send 'User' data to client 
type UserResponse struct {
    ID        uuid.UUID     `json:"id"`
    Username  string        `json:"username"`
    Firstname string        `json:"first_name"`
    Lastname  string        `json:"last_name"`
    Email     string        `json:"email"`
    Password  string        `json:"password"`
    Active    bool          `json:"active"`
    RoleID    uint          `json:"role_id"`
    Role      *RoleResponse `json:"role"`
}

// Credential is 'DTO' (Data Transfer Object) for user response 
// This is used for login/ credential checking or other similar operation
type Credential struct {
    Username  string `json:"username"`
    Email     string `json:"email"`
    Password  string `json:"password"`
    Active    bool   `json:"active"`
}

// convert 'User' model to 'Response' struct before displaying to client
func UserToResponse(user model.User) (userResponse *UserResponse) {
    var dtoResponse UserResponse
    // check Role
    if user.Role != nil {
        dtoResponse.Role = &RoleResponse{
            ID : user.Role.ID,
            Description : user.Role.Description,
            Name : user.Role.Name,
        }
    } else {
        dtoResponse.Role = nil
    }

    dtoResponse.ID = user.ID
    dtoResponse.Username  = user.Username
    dtoResponse.Firstname = user.Firstname
    dtoResponse.Lastname  = user.Lastname
    dtoResponse.Email     = user.Email
    dtoResponse.Password  = "[protected]"
    dtoResponse.Active    = user.Active
    dtoResponse.RoleID    = user.RoleID

    userResponse = &dtoResponse
    return
}

// convert 'Request' data from client to 'User' model before pased to database
func RequestToUser(userRx UserRequest) (user *model.User) {
    var userTmp model.User
    // check Role
    if userRx.Role != nil {
        userTmp.Role = &model.Role{
            ID          : userRx.Role.ID,
            Name        : userRx.Role.Name,
            Description : userRx.Role.Description,
        }
    }

    userTmp.Username  = userRx.Username
    userTmp.Firstname = userRx.Firstname
    userTmp.Lastname  = userRx.Lastname
    userTmp.Email     = userRx.Email
    userTmp.Password  = userRx.Password
    userTmp.Active    = userRx.Active
    userTmp.RoleID    = userRx.RoleID
    
    user = &userTmp

    return
}

// UserResponseToRequest will convert 'Response' data from client 
// into 'Request' before pased to database
func ResponseToRequestUser(userTx UserResponse) *UserRequest {
    // check Role
    var role RoleRequest
    if userTx.Role != nil {
        role.ID = userTx.Role.ID
        role.Name = userTx.Role.Name
        role.Description = userTx.Role.Description
    }

    return &UserRequest{
        Username  : userTx.Username,
        Firstname : userTx.Firstname,
        Lastname  : userTx.Lastname,
        Email     : userTx.Email,
        Password  : userTx.Password,
        Active    : userTx.Active,
        RoleID    : userTx.RoleID,
    }
}
