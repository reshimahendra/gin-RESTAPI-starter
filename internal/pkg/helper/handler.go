/*
   Package helper for handling response to client
*/
package helper

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Response is response helper to send to client/ user
type Response struct{
    Status  int         `json:"status"`
    Method  string      `json:"method"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}

// ErrorResponse is a response that containing error 
// Details of the error shoul not be technical as it will be sent to the user
type ErrorResponse struct{
    Status  int         `json:"status"`
    Method  string      `json:"method"`
    Error   interface{} `json:"error"`
}

// APIResponse will send JSON response to the client and some additional detail
func APIResponse(c *gin.Context, statusCode int, message string, data interface{}) {
    // prepare the response before sending to the client 
    res := Response{
        Status  : statusCode,
        Method  : c.Request.Method,
        Message : message,
        Data    : data,
    }

    // Send wrapped data to client 
    c.JSON(
        statusCode,
        res,
    )

    // In case error occur (starting from code 400 and up)
    // it will send an abort opreation with error code as the header
    if statusCode >= 400 {
        defer c.AbortWithStatus(statusCode)
    }
}

// APIErrorResponse will send JSON response with error value to the client 
func APIErrorResponse(c *gin.Context, statusCode int, err interface{}) {
    // Prepare the data before sending to the client
    res := ErrorResponse{
        Status : statusCode,
        Method  : c.Request.Method,
        Error  : err,
    }

    // Send wrapped data to client
    c.JSON(
        statusCode,
        res,
    )

    defer c.AbortWithStatus(statusCode)
}

// ValidationError is a 'Request' error detail generator 
// it will break given error value into detailed error message 
// so we know which field is the cause of error upon handling request
func ValidationError(err error) (eMessage *[]string) {
    // Create list of error message
    eMsg := []string{}
    for _, e := range err.(validator.ValidationErrors) {
        msg := fmt.Sprintf("Field error :'%s', condition: '%s'", e.Field(), e.ActualTag())
        eMsg = append(eMsg, msg)
    }

    eMessage = &eMsg

    return
}

// APIValidationErrorResponse is wrapper of 2 function, 'APIErrorResponse' and 'ValidationError'
// It just to simplify it instead of calling both of them at same time (it mostly be called in tandem)
func APIValidationErrorResponse(c *gin.Context, statusCode int, err interface{}) {
    APIErrorResponse(c, statusCode, ValidationError(err.(error)))
}
