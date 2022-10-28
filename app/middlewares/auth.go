package middlewares

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/snykk/golib_backend/utils/token"
)

func AuthorizeJWT(jwtService token.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		token, err := token.GetToken(authHeader, jwtService)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
			return
		}

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "token is not valid"})
	}
}
