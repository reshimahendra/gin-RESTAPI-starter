/*
   Package 'service' helper for 'User' model
   * DTO to Model conversion
   * Model to DTO conversion
*/
package service

import (
	"github.com/reshimahendra/gin-starter/internal/account/model"
	"github.com/reshimahendra/gin-starter/pkg/logger"
)

// convert 'User' model to 'Response' struct before displaying to client
func UserToResponse(user model.User) *UserResponse{
    // check Role
    var role RoleResponse
    if &user.Role != nil {
        role.ID = user.Role.ID
        role.Description = user.Role.Description
        role.Name = user.Role.Name
    }

    logger.Infof("User ro response: %v", user)
    logger.Infof("User ro response: %v", user.Role)

    return &UserResponse{
        ID        : user.ID,
        Username  : user.Username,
        Firstname : user.Firstname,
        Lastname  : user.Lastname,
        Email     : user.Email,
        Password  : "[protected]",
        Active    : user.Active,
        RoleID    : user.RoleID,
        Role      : &role,
    }
}

// convert 'Request' data from client to 'User' model before pased to database
func RequestToUser(userRx UserRequest) *model.User{
    // check Role
    var role model.Role
    if userRx.Role != nil {
        role.ID = userRx.Role.ID
        role.Name = userRx.Role.Name
        role.Description = userRx.Role.Description
    }

    return &model.User{
        Username  : userRx.Username,
        Firstname : userRx.Firstname,
        Lastname  : userRx.Lastname,
        Email     : userRx.Email,
        Password  : userRx.Password,
        Active    : userRx.Active,
        RoleID    : userRx.RoleID,
        Role      : &role,
    }
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
        Role      : &role,
    }
}
