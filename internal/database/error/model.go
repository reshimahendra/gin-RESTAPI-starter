/*
   Package database to handle error
   containing custom error for our app
*/
package dberror

import "fmt"

const (
    // ErrDataEmpty is error code for empty data result
    ErrDataEmpty = iota

    // ErrDataNotFound is error code for not found data 
    ErrDataNotFound

    // ErrSaveDataFail is error code for 'failling' on saving data
    ErrSaveDataFail

    // ErrDataCouldNotUpdate is error code for 'failling' on update data
    ErrUpdateDataFail

    // ErrDeleteData is error code for failing to delete data
    ErrDeleteDataFail

    // ErrDataExist is error code when triying to save data on an already exist data
    // for example 'Primary Key' or 'Unique Constraint' already exist 
    ErrDataExist

    // ErrPasswordNotMatch is error code for password that are not match when compared
    ErrPasswordNotMatch

    // ErrPasswordTooShort is error code for too short password
    ErrPasswordTooShort
)


// Error is custom dbError struct
type Error struct {
    Code    uint        `json:"code"`
    Message string      `json:"message,omitempty"`
    Err     interface{} `json:"error,omitempty"`
}

// type DBError struct{}

func (e *Error) Error() string{
    return fmt.Sprintf("Code: %d, Message: %s, Error Detail: %v", e.Code, e.Message, e.Err)
}

// New will create new dbError instance
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
        case ErrDataEmpty           : message = "error data empty."
        case ErrDataNotFound        : message = "error data not found."
        case ErrSaveDataFail        : message = "error fail saving data."
        case ErrUpdateDataFail      : message = "error fail updating data."
        case ErrDeleteDataFail      : message = "error fail deleting data."
        case ErrDataExist           : message = "error data already exist."
        case ErrPasswordNotMatch    : message = "error password not match."
        case ErrPasswordTooShort    : message = "error password too short."
        default                     : message = "error while processing data."
    }

    return
}
