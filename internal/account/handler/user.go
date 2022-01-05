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
        logger.Errorf("Error request data. 'user' param required.")
        helper.APIErrorResponse(c, http.StatusBadRequest, "user param is missing.")
        return
    }

    // Tell the 'service' to order 'repository' to get user data
    dtoResponse, err := h.service.Get(uname)
    if err != nil {
        logger.Errorf("Error while fetching user data: %v", err)
        // helper.APIErrorResponse(c, http.StatusBadRequest, err)
        helper.APIValidationErrorResponse(c, http.StatusBadRequest, err)
        return
    }

    helper.APIResponse(c, http.StatusOK, "Fetch data success." ,dtoResponse)

}

// Save will process request data from client, forward it to 'service'
// before passed to 'repository' and saved to the database
func (h *userHandler) Save(c *gin.Context) {
    var res service.UserRequest

    // check if the data is OK
    err := c.ShouldBindJSON(&res)
    if err != nil {
        // save error to logfile
        logger.Errorf("Error processing request data: %v", err)

        helper.APIErrorResponse(c, http.StatusBadRequest, err)
        return
    }

    dtoResponse, err := h.service.Save(res)
    if err != nil {
        // save error 'save user data' to logfile
        logger.Errorf("Error saving data: %v", err)
        helper.APIErrorResponse(c, http.StatusInternalServerError, 
            "error while saving data. process aborted",
            // err,
        )
        return
    }

    helper.APIResponse(c, http.StatusOK, "Save data success.", dtoResponse)
}
