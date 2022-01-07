/*
   Package handler for 'User role'
*/
package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/reshimahendra/gin-starter/internal/account/service"
	dbErr "github.com/reshimahendra/gin-starter/internal/database/error"
	"github.com/reshimahendra/gin-starter/internal/pkg/helper"
	"github.com/reshimahendra/gin-starter/pkg/logger"
)

// userRoleHandler is type wrapper for UserRole service to communicate with
// user role service
type userRoleHandler struct {
    service service.UserRoleService
}

// NewUserRole will create new instance to communicate with service 
// via NewUserRole
func NewUserRole(srv service.UserRoleService) *userRoleHandler {
    return &userRoleHandler{service: srv}
}

// Get will send request to 'service' module to get 'user role' data
func (h *userRoleHandler) Get(c *gin.Context) {
    // get param id from request url
    id := c.Params.ByName("id")
    
    // if param id not found, response with error data and exit process
    roleID, err := strconv.Atoi(id)
    if err != nil {
        err := dbErr.New(dbErr.ErrParamIDEmpty, nil)
        logger.Errorf("param id not valid: %v", err)
        helper.APIErrorResponse(c, http.StatusBadRequest, err)
        return

    }

    // send request to get data 'user role' from the service module
    roleResponse, err := h.service.Get(uint(roleID))
    if err != nil {
        logger.Errorf("error retreiving user role data: %v", err)
        helper.APIErrorResponse(c, http.StatusBadRequest, err)
        return
    }

    // if no error found, send response data to user
    helper.APIResponse(
        c,
        http.StatusOK,
        "succesfull retreiving "+roleResponse.Name+" role data.",
        roleResponse,
    )
}

// Gets will send request to 'service' module to get ALL 'user role' data
func (h *userRoleHandler) Gets(c *gin.Context) {
    // get all user role data response
    rolesResponse, err := h.service.Gets()
    if err != nil {
        logger.Errorf("error retreiving user role data: %v", err)
        helper.APIErrorResponse(c, http.StatusInternalServerError, err)
        return
    }

    // send user role dto data to client
    helper.APIResponse(
        c,
        http.StatusOK,
        "succesfull retreiving user roles data.",
        rolesResponse,
    )
}

// Create will send request to 'service' module to create new 'user role' data 
func (h *userRoleHandler) Create(c *gin.Context) {
    // get input data from the request 'context'
    var roleRequestDto service.RoleRequest
    err := c.ShouldBindJSON(&roleRequestDto)

    // check for data binding error. if found, exit process
    if err != nil {
        logger.Errorf("error on binding user role data: %v", err)
        helper.APIErrorResponse(c, http.StatusBadRequest, err)
        return
    }

    // send request to the 'service' module to save the binded data (roleRequestDto)
    roleResponseDto, err := h.service.Create(roleRequestDto)
    if err != nil {
        logger.Errorf("error creating user role data: %v", err)
        helper.APIErrorResponse(c, http.StatusBadRequest, err)
        return
    }

    // send saved 'user role' data as a response to the client
    helper.APIResponse(
        c,
        http.StatusOK,
        "seccesfull creating "+ roleResponseDto.Name+" role.",
        roleResponseDto,
    )
}

func (h *userRoleHandler) Update(c *gin.Context) {
    // get param id from request url
    id := c.Params.ByName("id")

    // if param id not found, response with error data and exit process
    roleID, err := strconv.Atoi(id)
    if err != nil  {
        err := dbErr.New(dbErr.ErrParamIDEmpty, nil)
        logger.Errorf("param id not valid: %v", err)
        helper.APIErrorResponse(c, http.StatusBadRequest, err)
        return
    }

    // get input data from the request 'context'
    var roleRequestDto service.RoleRequest
    err = c.ShouldBindJSON(&roleRequestDto)

    // check for data binding error. if found, exit process
    if err != nil {
        logger.Errorf("error on binding user role data: %v", err)
        helper.APIErrorResponse(c, http.StatusBadRequest, err)
        return
    }

    // send request to the 'service' module to update the binded data (roleRequestDto)
    roleResponseDto, err := h.service.Update(uint(roleID), roleRequestDto)
    if err != nil {
        logger.Errorf("error updating user role data: %v", err)
        helper.APIErrorResponse(c, http.StatusBadRequest, err)
        return
    }

    // send updated 'user role' data as a response to the client
    helper.APIResponse(
        c,
        http.StatusOK,
        "seccesfull updating "+ roleResponseDto.Name+" role.",
        roleResponseDto,
    )
}
