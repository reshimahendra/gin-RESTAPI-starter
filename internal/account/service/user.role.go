/*
   Package service for 'User Role'
   It will implementing interface available on 'user.role repository'
   to create bridge for the 'repository' package and 'handler' package
*/
    
package service

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
    Name        string  `json:"name" binding:"required"`
    Description string  `json:"description"`
}
