/*
   Authentication Handler
   Including:
   * Signup
   * Signin
   * RefreshToken
   * CheckToken
*/
package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/reshimahendra/gin-starter/internal/account/model"
	"github.com/reshimahendra/gin-starter/internal/account/service"
	"github.com/reshimahendra/gin-starter/internal/pkg/auth"
	"github.com/reshimahendra/gin-starter/internal/pkg/helper"
	"github.com/reshimahendra/gin-starter/pkg/logger"
	"gorm.io/gorm"
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

const (
    userSigninErr  = "User email and password does not match or not exist."
    tokenInvalid   = "Token is already expired or not valid."
    tokenNotFound  = "Token could not found!"
    tokenCreateErr = "Could not create token!"
)

type Controller struct {
    db *gorm.DB
}

// Signup controller
func (ctl *Controller) Signup(c *gin.Context) {    
    var u model.UserRequest

    err := c.ShouldBindJSON(&u)
    if err != nil {
        logger.Errorf("Error signup: %v", err)
        c.AbortWithStatusJSON(400, gin.H{"error":"Bad request"})
    }

    fmt.Println(u)


    // Signup controller logic here
    isUserExist := service.IsUserExist(ctl.db, u.Email, u.Username)
    fmt.Println(isUserExist)
    if isUserExist {
        logger.Errorf("User already exist. Signup aborted.")
        c.AbortWithStatus(400)
    }

    // Generate hash password for the new user
    u.Password, err = helper.HashPassword(u.Password)
    if err != nil {
        logger.Errorf("Cannot create hash password on signup: %v", err)
        c.AbortWithStatusJSON(400, gin.H{"error":"Cannot create hash password"})
        return
    }

    // convert 'dto User request' to 'model User'
    responseUser := service.RequestToUser(u)

    user, err := service.SaveUser(ctl.db, *responseUser)
    if err != nil {
        logger.Errorf("Error signup: %v", err)
        c.AbortWithStatusJSON(400, gin.H{"error":"Signup error"})
        return
    }

    // TODO: send mail to user based on their email address to activate the account
    // OR create 'Hook' on table 'User' AfterCreate to execute the operation

    c.JSON(200, user)
}

// Signin controller
func (ctl *Controller) Signin(c *gin.Context) {
    var credential auth.AuthLoginDTO
    err := c.ShouldBindJSON(&credential)
    if err != nil {
        c.JSON(
            http.StatusUnauthorized,
            gin.H{
                "code"    : http.StatusUnauthorized,
                "message" : tokenNotFound,
            },
        )
        return
    }

    // User is registered 
    isActive, isPasswordMatch := service.CredentialByEmail(ctl.db, credential.Email, credential.Password)

    // User not found
    if !isPasswordMatch {
        c.AbortWithStatus(400)
    }

    // User not active 
    if !isActive {
        c.JSON(403, gin.H{
            "error": "User not active.",
        })
        return
    }

    if isActive && isPasswordMatch {
        token, err := auth.CreateToken(credential.Email)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusInternalServerError, tokenCreateErr)
            return
        }

        authLoginResponse := auth.AuthLoginResponse{
            AccessToken     : token.AccessToken,
            RefreshToken    : token.RefreshToken,
            TransmissionKey : token.TransmissionKey,
        }

        c.JSON(http.StatusOK, authLoginResponse)
    } else {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
            "code"    : http.StatusUnauthorized,
            "message" : userSigninErr,
            "token"   : nil,
        })
    }
}

// Refresh Token
func (ctl *Controller) RefreshToken(c *gin.Context) {
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
        c.AbortWithStatusJSON(http.StatusUnauthorized, tokenInvalid)
        return
    }

    email := token.Claims.(jwt.MapClaims)["email"].(string)

    // Create new token 
    newToken, err := auth.CreateToken(email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, tokenCreateErr)
        return
    }

    authLoginResponse := auth.AuthLoginResponse{
        AccessToken : newToken.AccessToken,
        RefreshToken : newToken.RefreshToken,
        TransmissionKey : newToken.TransmissionKey,
    }

    c.JSON(http.StatusOK, authLoginResponse)
}

// Check token
func (ctl *Controller) CheckToken(c *gin.Context) {
    var decToken string
    bearerToken := c.GetHeader("Authorization")
    authArray := strings.Split(bearerToken, " ")
    if len(authArray) == 2 {
        decToken = authArray[1]
    }

    if decToken == "" {
        c.AbortWithStatusJSON(http.StatusUnauthorized, tokenNotFound)
        return
    }

    token, err := auth.TokenValid(decToken)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusUnauthorized, tokenInvalid)
        return
    }

    email := token.Claims.(jwt.MapClaims)["email"].(string)
    // claims := token.Claims.(*jwt.MapClaims)
    // email := claims["email"]

    c.JSON(http.StatusOK, email)
}
