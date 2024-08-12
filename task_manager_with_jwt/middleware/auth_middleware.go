package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Abzaek/GO-APIs/task-manager-with-mongo/controllers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var JwtKey = []byte("abzaeko")

type CustomClaim struct {
	ID   string `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

func AuthMiddleware(role string, ctrl *controllers.Control) gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.String(http.StatusBadRequest, "Header is required")
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")

		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.String(http.StatusBadRequest, "invalid header")
			fmt.Println(authParts[1])
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(
			authParts[1],
			&CustomClaim{},
			func(t *jwt.Token) (interface{}, error) {
				return JwtKey, nil
			},
		)

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token", "err": err.Error()})
			c.Abort()
			return
		}

		if role == "admin" {

			if claims, ok := token.Claims.(*CustomClaim); ok {

				if claims.Role != "admin" && ctrl.GetSingleUser(claims.ID).Role != "admin" {

					c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
					c.Abort()
					return

				}

			}

		}
		c.Next()
	}
}
