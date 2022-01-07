/*
   Package 'service' helper for 'User' model
   * DTO to Model conversion
   * Model to DTO conversion
*/
package service

import (
	"github.com/reshimahendra/gin-starter/internal/account/model"
)

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
func UserResponseToRequest(userTx UserResponse) *UserRequest {
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
