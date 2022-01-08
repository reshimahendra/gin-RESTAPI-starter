/*
    Package errors
    * processing custom error
*/
package errors

import "fmt"

// simpleError is custom error struct that only showing simple error info 
type simpleError struct {
    Code    uint    `json:"code"`
    Message string  `json:"message"`
}

// Error method for displaying error string for 'simpleError' struct
func (e *simpleError) Error() string {
    return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

// New will create new 'simpleError' error instance 
func NewSimpleError(code uint) error {
    return &simpleError{
        Code    : code,
        Message : Message(code),
    }
}

// Error is custom error struct that can 'pass' the main Error detail via 'Err' field
type Error struct {
    Code    uint        `json:"code"`
    Message string      `json:"message,omitempty"`
    Err     interface{} `json:"error,omitempty"`
}

// Error method for displaying error string for 'Error' struct
func (e *Error) Error() string{
    return fmt.Sprintf("Code: %d, Message: %s, Error Detail: %v", e.Code, e.Message, e.Err)
}

// New will create new 'Error' error instance
func New(code int, e error) error { 
    return &Error{
        Code    : uint(code),
        Message : Message(uint(code)),
        Err     : e,
    }
}

// Message wiil return error message
func Message(code uint) (message string) {
    switch code {
        // database error
        case ErrDataIsEmpty         : message = ErrDataIsEmptyMsg 
        case ErrDataNotFound        : message = ErrDataNotFoundMsg
        case ErrGettingData         : message = ErrGettingDataMsg
        case ErrSaveDataFail        : message = ErrSaveDataFailMsg
        case ErrUpdateDataFail      : message = ErrUpdateDataFailMsg 
        case ErrDeleteDataFail      : message = ErrDeleteDataFailMsg
        case ErrDataAlreadyExist    : message = ErrDataAlreadyExistMsg
        
        // auth error
        case ErrSignUp                  : message = ErrSignUpMsg 
        case ErrSignIn                  : message = ErrSignInMsg 
        case ErrSignOut                 : message = ErrSignOutMsg
        case ErrUserNotRegistered       : message = ErrUserNotRegisteredMsg 
        case ErrUserAlreadyRegistered   : message = ErrUserNotRegisteredMsg 
        case ErrUserNotActive           : message = ErrUserNotActiveMsg 
        case ErrPasswordNotMatch        : message = ErrPasswordNotMatchMsg
        case ErrPasswordTooShort        : message = ErrPasswordTooShortMsg
        case ErrTokenCreate             : message = ErrTokenCreateMsg
        case ErrTokenRefresh            : message = ErrTokenRefreshMsg
        case ErrTokenInvalid            : message = ErrTokenInvalidMsg
        case ErrTokenNotFound           : message = ErrTokenNotFoundMsg

        // handler error
        case ErrParamIsEmpty        : message = ErrParamIsEmptyMsg
        case ErrParamIsInvalid      : message = ErrParamIsInvalidMsg
        case ErrUsernameIsInvalid   : message = ErrUsernameIsInvalidMsg
        case ErrEmailIsInvalid      : message = ErrEmailIsInvalidMsg
        case ErrRequestDataInvalid  : message = ErrRequestDataInvalidMsg

        // default/ unknown error
        default                     : message = "unknown error"
    }

    return
}
