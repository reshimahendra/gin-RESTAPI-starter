/*
   Package service for 'User Role'
   It will implementing interfaces available on 'user role repository'
   and create bridge for the 'repository' package and 'handler' package
*/
package service

import (
	"github.com/reshimahendra/gin-starter/internal/account/repository"
)

// UserRoleService is middleware interface for the repository and handler
type UserRoleService interface {
    Get(id uint) (roleResponse *RoleResponse, err error)
    Gets() (rolesResponse *[]RoleResponse, err error)
    Create(input RoleRequest) (roleResponse *RoleResponse, err error)
    Update(id uint, input RoleRequest) (roleResponse *RoleResponse, err error)
}

// userRoleService is wrapper for UserRoleRepository data type
type userRoleService struct {
    repo repository.UserRoleRepository  
}

// NewRole will create new instance to bridge 'user role repo' and 'user role handler'
func NewUserRole(r repository.UserRoleRepository) *userRoleService {
    return &userRoleService{repo: r}
}

// Get will get 'user role' based on given id.
// it will pass 'request dto' from 'handler' to 'repository'
func (s *userRoleService) Get(id uint) (roleResponse *RoleResponse, err error) {
    role, err := s.repo.Get(id)
    if err != nil {
        return
    }

    // convert 'user role' model to 'response dto'
    roleResponse = RoleToResponse(*role) 

    return
}

// Gets will get all 'user role' data 
// it will send request to 'repository' and pass the 'result dto' to handler
func (s *userRoleService) Gets() (rolesResponse *[]RoleResponse, err error) {
    roles, err := s.repo.Gets()
    if err != nil {
        return nil, err
    }

    var tmpUsersResponse []RoleResponse
    for _, role := range *roles {
        tmpUsersResponse = append(tmpUsersResponse, *RoleToResponse(role))
    }

    // conver to response dto
    rolesResponse = &tmpUsersResponse

    return
}

// Save will save data 'user role' to database based on given 'request dto' input
func (s *userRoleService) Create(input RoleRequest) (roleResponse *RoleResponse, err error) {
    // convert 'role' dto to savable model
    roleTmp := RequestToRole(input)

    role, err := s.repo.Create(*roleTmp)
    if err != nil {
        return nil, err
    }

    // convert 'user role' model to dto format
    roleResponse = RoleToResponse(*role)
    
    return
}

// Update will update 'user role' data based on given 'request dto' 
// the update will apply to data with matching param 'id'
func (s *userRoleService) Update(id uint, input RoleRequest) (roleResponse *RoleResponse, err error) {
    // convert 'user role' data dto to savable model 
    roleTmp := RequestToRole(input)

    role, err := s.repo.Update(id, *roleTmp)
    if err != nil {
        return nil, err
    }

    // convert 'user role' model to dto format
    roleResponse = RoleToResponse(*role)

    return
}
