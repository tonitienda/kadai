package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var mySigningKey = []byte(os.Getenv("AUTH0_SIGNING_SECRET")) // The secret you get from Auth0

type Auth interface {
	GetUserIdBySub(sub string) (string, bool)
}

func TokenAuthMiddleware(auth Auth) gin.HandlerFunc {

	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")[7:] // Skip "Bearer "

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return mySigningKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			sub := claims["sub"]
			userId, ok := auth.GetUserIdBySub(sub.(string))

			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "User could not be authenticated"})
				c.Abort()
				return
			}

			c.Set("userId", userId)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
	}
}
