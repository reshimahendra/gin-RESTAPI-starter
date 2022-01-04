/*
   Middleware to prevent unauthorized access
*/
package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/reshimahendra/gin-starter/internal/pkg/auth"
)

const (
    tokenInvalid   = "Token is already expired or not valid."
    tokenNotFound  = "Token could not found!"
)

// Authorize is middleware to prevent unauthorized access
func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenStr string
		bearerToken := c.GetHeader("Authorization")
		strArr := strings.Split(bearerToken, " ")
		if len(strArr) == 2 {
			tokenStr = strArr[1]
		}

		if tokenStr == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, tokenNotFound)
			return
		}

		token, err := auth.TokenValid(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, tokenInvalid)
			return
		}

		if err != nil && !token.Valid {
		//     claims := token.Claims.(jwt.MapClaims)
		//     fmt.Println(claims)
		// } else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code"    : http.StatusUnauthorized,
				"message" : tokenInvalid,
				"token"   : nil,
			})
		}

	}
}
