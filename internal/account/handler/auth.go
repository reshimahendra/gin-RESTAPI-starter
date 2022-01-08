/*
   Package handler for Authentication
   Including:
   * Signup
   * Signin
   * RefreshToken
   * CheckToken
*/
package handler 

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/reshimahendra/gin-starter/internal/account/service"
	"github.com/reshimahendra/gin-starter/internal/pkg/auth"
	"github.com/reshimahendra/gin-starter/internal/pkg/helper"
	E "github.com/reshimahendra/gin-starter/pkg/errors"
	"github.com/reshimahendra/gin-starter/pkg/logger"
)

/**

  Token testing: using 'curl'
  *** Get token by 'SIGN IN'
  ❯ curl -s -X POST -H 'Accept: application/json' -H 'Content-Type: application/json' \
  --data '{"email":"lbw@example.com","password":"password"}' \
  http:/127.0.0.1:8000/auth/sign-in

  *** Refresh Token
  ❯ curl -X POST -H 'content-type:application/json/' http://127.0.0.1:8000/auth/refresh \
    -d '{"access_token":"<token>","refresh_token":"<refresh_token>","transmission_key":"<transmission_key>"}'

  *** Check Token
  ❯ curl -X POST -H 'content-type:application/json/' -H "Authorization: Bearer <token>" \
  http://127.0.0.1:8000/auth/check
*/

// Signup controller
func (h *userHandler) Signup(c *gin.Context) {    
    var userRequest service.UserRequest

    err := c.ShouldBindJSON(&userRequest)
    if err != nil {
        e := E.New(E.ErrRequestDataInvalid, err)
        logger.Errorf("%s. %v", E.ErrSignUpMsg, err)
        helper.APIErrorResponse(c, http.StatusBadRequest, e)

        return
    }

    // Signup controller logic here
    isAccountAvailable := h.service.UserAvailable(userRequest.Email, userRequest.Username)
    if !isAccountAvailable {
        err := E.NewSimpleError(E.ErrUserAlreadyRegistered)
        logger.Errorf("%s. %v", E.ErrSignUpMsg, err)
        helper.APIErrorResponse(c, http.StatusBadRequest, err)

        return
    }

    // create user account. exit if error
    userResponse, err := h.service.Create(userRequest)
    if err != nil {
        logger.Errorf("%s. %v", E.ErrSignUpMsg, err)
        helper.APIErrorResponse(c, http.StatusInternalServerError, err)

        return
    }

    // TODO: send mail to user based on their email address to activate the account
    // OR create 'Hook' on table 'User' AfterCreate to execute the operation

    helper.APIResponse(c, http.StatusOK, "successful signup", userResponse)
}

// Signin controller
func (h *userHandler) Signin(c *gin.Context) {
    // get login data from context
    var credential auth.AuthLoginDTO
    err := c.ShouldBindJSON(&credential)
    if err != nil {
        e := E.New(E.ErrRequestDataInvalid, err)
        logger.Errorf("%s: %v", E.ErrRequestDataInvalidMsg, err)
        helper.APIErrorResponse(c, http.StatusUnauthorized, e)

        return
    }

    // User is registered 
    isActive, isValid := h.service.CheckCredentialByMail(credential.Email, credential.Password)

    // user / password not valid, exit the process
    if !isValid {
        err := E.NewSimpleError(E.ErrPasswordNotMatch)
        logger.Errorf("login fail: %v", err)
        helper.APIErrorResponse(c, http.StatusUnauthorized, err)

        return
    }

    // User not active 
    if !isActive {
        err := E.NewSimpleError(E.ErrUserNotActive)
        logger.Errorf("login fail: %v", err)
        helper.APIErrorResponse(c, http.StatusUnauthorized, err)

        return
    }

    if isActive && isValid {
        token, err := auth.CreateToken(credential.Email)
        if err != nil {
            e := E.New(E.ErrTokenCreate, err)
            logger.Errorf("%s: %v", E.ErrTokenCreateMsg, e)
            helper.APIErrorResponse(c, http.StatusInternalServerError, e)

            return
        }

        // send token data response to the client
        authLoginResponse := auth.AuthLoginResponse{
            AccessToken     : token.AccessToken,
            RefreshToken    : token.RefreshToken,
            TransmissionKey : token.TransmissionKey,
        }

        helper.APIResponse(c, http.StatusOK, "successful signin", authLoginResponse)
    }
}

// Refresh Token
func (h *userHandler) RefreshToken(c *gin.Context) {
    mapToken := map[string]string{}

    decoder := json.NewDecoder(c.Request.Body)
    if err := decoder.Decode(&mapToken); err != nil {
        errs := []string{"REFRESH_TOKEN_ERROR"}
        c.JSON(http.StatusUnprocessableEntity, errs)
        return
    }

    defer c.Request.Body.Close()

    token, err := auth.TokenValid(mapToken["refresh_token"])
    if err != nil {
        e := E.New(E.ErrTokenRefresh, err)
        logger.Errorf("%s: %v", E.ErrTokenRefreshMsg, err)
        helper.APIErrorResponse(c, http.StatusUnauthorized, e)

        return
    }

    email := token.Claims.(jwt.MapClaims)["email"].(string)

    // Create new token 
    newToken, err := auth.CreateToken(email)
    if err != nil {
        e := E.New(E.ErrTokenCreate, err)
        logger.Errorf("%s: %v", E.ErrTokenCreateMsg, err)
        helper.APIErrorResponse(c, http.StatusInternalServerError, e)

        return
    }

    // send RefreshToken data response to the client
    authLoginResponse := auth.AuthLoginResponse{
        AccessToken : newToken.AccessToken,
        RefreshToken : newToken.RefreshToken,
        TransmissionKey : newToken.TransmissionKey,
    }

    helper.APIResponse(
        c,
        http.StatusOK,
        "successful refreshing token data",
        authLoginResponse,
    )
}

// Check token
func (h *userHandler) CheckToken(c *gin.Context) {
    var decToken string
    bearerToken := c.GetHeader("Authorization")
    authArray := strings.Split(bearerToken, " ")
    if len(authArray) == 2 {
        decToken = authArray[1]
    }

    if decToken == "" {
        err := E.NewSimpleError(E.ErrTokenNotFound)
        logger.Errorf("error check token: %v", err)
        helper.APIErrorResponse(c, http.StatusUnauthorized, err)

        return
    }

    token, err := auth.TokenValid(decToken)
    if err != nil {
        e := E.New(E.ErrTokenInvalid, err)
        logger.Errorf("%s: %v", E.ErrTokenInvalidMsg, e)
        helper.APIErrorResponse(c, http.StatusUnauthorized, e)

        return
    }

    // send response of the 'checkToken' result to client
    email := token.Claims.(jwt.MapClaims)["email"].(string)
    helper.APIResponse(
        c,
        http.StatusOK,
        "successful check token",
        email,
    )
}
