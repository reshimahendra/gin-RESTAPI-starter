/*
   Package handler for 'User'
*/
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reshimahendra/gin-starter/internal/account/service"
	"github.com/reshimahendra/gin-starter/internal/pkg/helper"
	"github.com/reshimahendra/gin-starter/pkg/logger"
	E "github.com/reshimahendra/gin-starter/pkg/errors"
)

// userHandler is Handler type wrapper for UserService
type userHandler struct {
    service service.UserService
}

// NewUser will return userHandler instance to communicate with UserServicee 
func NewUser(srv service.UserService) *userHandler{
    return &userHandler{service: srv}
}

// Get will retreive User 'DTO' (by username) that prepared by passed by 'service' 
func (h *userHandler) Get(c *gin.Context) {
    // get the param value that shipped with context
    uname := c.Params.ByName("user")
    if uname == "" {
        logger.Errorf("invalid username.")
        helper.APIErrorResponse(c, http.StatusBadRequest, "invalid username.")
        return
    }

    // Tell the 'service' to order 'repository' to get user data
    dtoResponse, err := h.service.Get(uname)
    if err != nil {
        logger.Errorf("error retreiving user data: %v", err)
        helper.APIErrorResponse(c, http.StatusBadRequest, err)
        return
    }

    helper.APIResponse(c, http.StatusOK, "success retreiving "+uname+" data." ,dtoResponse)

}

// GetByEmail will retreive User 'DTO' (by email) that prepared by passed by 'service' 
func (h *userHandler) GetByEmail(c *gin.Context) {
    // get the param value that shipped with context
    email := c.Params.ByName("email")
    // logger.Infof("v", helper.MailIsValid(email))
    if !helper.EmailIsValid(email) {
        logger.Errorf("invalid email.")
        helper.APIErrorResponse(c, http.StatusBadRequest, "invalid email.")
        return
    }

    // Tell the 'service' to order 'repository' to get user data
    dtoResponse, err := h.service.GetByEmail(email)
    if err != nil {
        logger.Errorf("error retreiving user data: %v", err)
        helper.APIErrorResponse(c, http.StatusBadRequest, err)
        return
    }

    helper.APIResponse(c, http.StatusOK, "success retreiving "+email+" data." ,dtoResponse)

}

func (h *userHandler) Gets(c *gin.Context) {
    usersDto, err := h.service.Gets()
    if err != nil {
        logger.Errorf("error retreiving user data.")
        helper.APIErrorResponse(c, http.StatusBadRequest, err)
        return
    }

    helper.APIResponse(c, http.StatusOK, "success retreiving users data.", usersDto)
}

// Create will process request data from client, forward it to 'service'
// before passed to 'repository' and saved to the database
func (h *userHandler) Create(c *gin.Context) {
    var res service.UserRequest

    // check if the data is OK
    err := c.ShouldBindJSON(&res)
    if err != nil {
        // save error to logfile
        logger.Errorf("error processing request data: %v", err)

        helper.APIErrorResponse(c, http.StatusBadRequest, err)
        return
    }

    dtoResponse, err := h.service.Create(res)
    if err != nil {
        // save error 'save user data' to logfile
        logger.Errorf("error saving data: %v", err)
        helper.APIErrorResponse(c, http.StatusInternalServerError, 
            "error saving user data. process aborted",
            // err,
        )
        return
    }

    helper.APIResponse(
        c, 
        http.StatusOK, 
        "success saving "+dtoResponse.Username+" data.", 
        dtoResponse)
}

// Update will sending request update data from client, forward it to 'service'
// before passed to 'repository' and finally update the data to the database
func (h *userHandler) Update(c *gin.Context) {
    // Get user param (username value)
    uname := c.Params.ByName("username")
    if uname == "" {
        err := E.NewSimpleError(E.ErrParamIsInvalid)
        logger.Errorf("update user. %v", err)
        helper.APIErrorResponse(c, http.StatusBadRequest, err)
        return
    }

    // get data from the context and bind it as json 
    var userReq service.UserRequest
    err := c.ShouldBindJSON(&userReq)
    if err != nil {
        e := E.New(E.ErrRequestDataInvalid, err)
        logger.Errorf("update user. %s: %v", E.ErrRequestDataInvalidMsg, e)

        helper.APIErrorResponse(c, http.StatusBadRequest, e)
        return
    }

    // request 'service' to process updating data. If error, abort update
    dtoResponse, err := h.service.Update(uname, userReq)
    if err != nil {
        // saving error 'updating user data' to logfile
        e := E.New(E.ErrUpdateDataFail, err)
        logger.Errorf("update user. %s: %v", E.ErrUpdateDataFailMsg, e)
        helper.APIErrorResponse(c, http.StatusBadRequest, e)
        return
    }

    helper.APIResponse(c, http.StatusOK, "success updating "+uname+" data.", dtoResponse)
}
