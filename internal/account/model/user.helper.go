/*
    Package 'model' helper for 'User' model
    * DTO to Model conversion
    * Model to DTO conversion
*/
package model

// convert 'User' model to 'Response' struct before displaying to client
func UserToResponse(user User) *UserResponse{
    // check Role
    var role RoleResponse
    if user.Role != nil {
        role.ID = user.Role.ID
        role.Description = user.Role.Description
        role.Name = user.Role.Name
    }

    return &UserResponse{
        ID        : user.ID,
        Username  : user.Username,
        Firstname : user.Firstname,
        Lastname  : user.Lastname,
        Email     : user.Email,
        Password  : user.Password,
        Active    : user.Active,
        Role      : &role,
    }
}

// convert 'Request' data from client to 'User' model before pased to database
func RequestToUser(userRx UserRequest) *User{
    // check Role
    var role Role
    if userRx.Role != nil {
        role.ID = userRx.Role.ID
        role.Name = userRx.Role.Name
        role.Description = userRx.Role.Description
    }

    return &User{
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
