/*
   Package handler for 'User role'
*/
package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/reshimahendra/gin-starter/internal/account/service"
	"github.com/reshimahendra/gin-starter/internal/pkg/helper"
	E "github.com/reshimahendra/gin-starter/pkg/errors"
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
        e := E.New(E.ErrParamIsInvalid, err)
        logger.Errorf("get user role. %s: %v", E.ErrParamIsEmptyMsg, e)
        helper.APIErrorResponse(c, http.StatusBadRequest, e)
        return

    }

    // send request to get data 'user role' from the service module
    roleResponse, err := h.service.Get(uint(roleID))
    if err != nil {
        e := E.New(E.ErrGettingData, err)
        logger.Errorf("get user role. %s: %v", E.ErrGettingDataMsg, e)
        helper.APIErrorResponse(c, http.StatusBadRequest, e)
        return
    }

    // if no error found, send response data to user
    helper.APIResponse(
        c,
        http.StatusOK,
        "successful retreiving "+roleResponse.Name+" role data.",
        roleResponse,
    )
}

// Gets will send request to 'service' module to get ALL 'user role' data
func (h *userRoleHandler) Gets(c *gin.Context) {
    // get all user role data response
    rolesResponse, err := h.service.Gets()
    if err != nil {
        e := E.New(E.ErrGettingData, err)
        logger.Errorf("gets user role. %s: %v", E.ErrGettingDataMsg, e)
        helper.APIErrorResponse(c, http.StatusInternalServerError, e)
        return
    }

    // send user role dto data to client
    helper.APIResponse(
        c,
        http.StatusOK,
        "successful retreiving user roles data.",
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
        e := E.New(E.ErrRequestDataInvalid, err)
        logger.Errorf("create user role. %s: %v", E.ErrRequestDataInvalidMsg, e)
        helper.APIErrorResponse(c, http.StatusBadRequest, e)
        return
    }

    // send request to the 'service' module to save the binded data (roleRequestDto)
    roleResponseDto, err := h.service.Create(roleRequestDto)
    if err != nil {
        e := E.New(E.ErrSaveDataFail, err)
        logger.Errorf("create user role. %s: %v", E.ErrSaveDataFailMsg, e)
        helper.APIErrorResponse(c, http.StatusBadRequest, e)
        return
    }

    // send saved 'user role' data as a response to the client
    helper.APIResponse(
        c,
        http.StatusOK,
        "seccessful creating "+ roleResponseDto.Name+" role.",
        roleResponseDto,
    )
}

func (h *userRoleHandler) Update(c *gin.Context) {
    // get param id from request url
    id := c.Params.ByName("id")

    // if param id not found, response with error data and exit process
    roleID, err := strconv.Atoi(id)
    if err != nil  {
        e := E.New(E.ErrParamIsInvalid, err)
        logger.Errorf("update user role. %s: %v", E.ErrParamIsInvalidMsg, e)
        helper.APIErrorResponse(c, http.StatusBadRequest, e)
        return
    }

    // get input data from the request 'context'
    var roleRequestDto service.RoleRequest
    err = c.ShouldBindJSON(&roleRequestDto)

    // check for data binding error. if found, exit process
    if err != nil {
        e := E.New(E.ErrRequestDataInvalid, err)
        logger.Errorf("update user role. %s: %v", E.ErrRequestDataInvalidMsg, e)
        helper.APIErrorResponse(c, http.StatusBadRequest, e)
        return
    }

    // send request to the 'service' module to update the binded data (roleRequestDto)
    roleResponseDto, err := h.service.Update(uint(roleID), roleRequestDto)
    if err != nil {
        e := E.New(E.ErrUpdateDataFail, err)
        logger.Errorf("update user role. %s: %v", E.ErrUpdateDataFailMsg, e)
        helper.APIErrorResponse(c, http.StatusBadRequest, e)
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
