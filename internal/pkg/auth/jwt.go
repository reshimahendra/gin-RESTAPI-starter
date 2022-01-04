/* 
    Authentication routine with jwt
*/
package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/reshimahendra/gin-starter/internal/pkg/helper"
	"github.com/reshimahendra/gin-starter/internal/config"
)

// AuthLoginDTO is 'DTO' (Data Transfer Object) to verify user on login 
type AuthLoginDTO struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

// AuthLoginResponse is 'DTO' (Data Transfer Object) to 'Response'
// or sending data to user upon 'login' or request 'refresh token'
type AuthLoginResponse struct {
    AccessToken     string  `json:"access_token"`
    RefreshToken    string  `json:"refresh_token"`
    TransmissionKey string  `json:"transmission_key"`
}

// TokenDetailsDTO is 'DTO' (data Transfer Object) containing
// details of token expiration time
type TokenDetailsDTO struct {
    AccessToken     string  `json:"access_token"`
    RefreshToken    string  `json:"refresh_token"`
    AtExpiresTime   time.Time
    RtExpiresTime   time.Time
    TransmissionKey string  `json:"transmission_key"`
}

// CreateToken will 'create' a jwt token
func CreateToken(email string) (*TokenDetailsDTO, error) {
    var err error

    config := config.GetConfig()
    tokenDetail := &TokenDetailsDTO{}
    tokenDetail.AtExpiresTime = time.Now().Add(
        time.Duration(config.Server.AccessTokenExpireDuration) * time.Hour)
    tokenDetail.RtExpiresTime = time.Now().Add(
        time.Duration(config.Server.RefreshTokenExpireDuration) * time.Hour)

    // Construct token
    accessTokenClaims := jwt.MapClaims{}
    accessTokenClaims["email"]     = email
    accessTokenClaims["user_uuid"] = "user_uuid"
    accessTokenClaims["exp"]       = time.Now().Add(time.Hour * 48).Unix()
    accessTokenClaims["uuid"]      = ""
    
    accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

    tokenDetail.AccessToken, err = accessToken.SignedString([]byte(config.Server.SecretKey))
    if err != nil {
        return nil, err
    }

    // Construct refresh token 
    refreshTokenClaims := jwt.MapClaims{}
    refreshTokenClaims["email"]     = email
    refreshTokenClaims["user_uuid"] = "user_uuid"
    refreshTokenClaims["exp"]       = time.Now().Add(time.Hour * 96).Unix()
    refreshTokenClaims["uuid"]      = ""

    refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

    tokenDetail.RefreshToken, err = refreshToken.SignedString([]byte(config.Server.SecretKey))
    if err != nil {
        return nil, err
    }

    // Generate secure key
    generateKey, err := helper.GenerateSecureKey(16)
    if err != nil {
        return nil, err
    }
    tokenDetail.TransmissionKey = generateKey

    return tokenDetail, nil 
}

// verifyToken will verify the given token 
func verifyToken(token string) (*jwt.Token, error) {
    config := config.GetConfig()
    verifiedToken, err := jwt.Parse(token, func (verifiedToken *jwt.Token) (interface{}, error) {
        if _, ok := verifiedToken.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", verifiedToken.Header["alg"])
        }
        return []byte(config.Server.SecretKey), nil
    })

    if err != nil {
        return verifiedToken, fmt.Errorf("Unauthorized access")
    }

    return verifiedToken, nil
}

// TokenValid will check whether the 'given' token was valid or not
func TokenValid(bearerToken string) (*jwt.Token, error) {
    token, err := verifyToken(bearerToken)
    if err != nil {
        if token != nil {
            return token, err
        }
        return nil, err
    }

    if !token.Valid {
        return nil, fmt.Errorf("Unauthorized access")
    }

    return token, nil
}
