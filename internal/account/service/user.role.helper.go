/*
   Package service for 'User Role'
   It will implementing interface available on 'user.role repository'
   to create bridge for the 'repository' package and 'handler' package
*/

package service

import "github.com/reshimahendra/gin-starter/internal/account/model"

// RoleRequest is 'DTO' (Data Transfer Object) for 'Role' model
// It will receive 'Role' data and processed (save/update) to database
type RoleRequest struct {
    ID          uint    `json:"id"`
    Name        string  `json:"name" binding:"required"`
    Description string  `json:"description"`
}

// RoleResponse is 'DTO' (Data Transfer Object) for 'Role' model
// It will send 'role' data to client 
type RoleResponse struct {
    ID          uint    `json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
}

// RoleToResponse will convert 'Role' model to 'Response DTO' format
func RoleToResponse(role model.Role) (roleResponse *RoleResponse) {
    return &RoleResponse{
        ID          : role.ID,
        Name        : role.Name,
        Description : role.Description,
    }
}

// RequestToRole will convert user 'input request' role dto to savable 'role model' format
func RequestToRole(roleRx RoleRequest) (role *model.Role) {
    return &model.Role{
        ID          : roleRx.ID,
        Name        : roleRx.Name,
        Description : roleRx.Description,
    }
}

// RequestToResponseRole will convert 'request DTO' to 'Response DTO' format
func RequestToResponseRole(roleRx RoleRequest) (roleResponse *RoleResponse) {
    return &RoleResponse{
        ID          : roleRx.ID,
        Name        : roleRx.Name,
        Description : roleRx.Description,
    }
}
