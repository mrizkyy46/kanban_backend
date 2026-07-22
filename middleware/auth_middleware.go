package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, ok := getBearerToken(c)
		if !ok {
			return
		}

		token, ok := parseJWT(tokenString)
		if !ok {
			abortWithUnauthorized(c, "Token invalid or expired")
			return
		}

		userID, ok := getUserIDFromClaims(token)
		if !ok {
			abortWithUnauthorized(c, "User not found")
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}

func getBearerToken(c *gin.Context) (string, bool) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		abortWithUnauthorized(c, "No token access (Authorization header needed)")
		return "", false
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		abortWithUnauthorized(c, "Invalid format token: Bearer <token>")
		return "", false
	}
	return tokenString, true
}

func parseJWT(tokenString string) (*jwt.Token, bool) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signin method: %v", t.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil || token == nil || !token.Valid {
		return nil, false
	}
	return token, true
}

func getUserIDFromClaims(token *jwt.Token) (uuid.UUID, bool) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.Nil, false
	}

	userIDStr, ok := claims["user_id"].(string)
	if !ok {
		return uuid.Nil, false
	}

	parsedUserID, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, false
	}
	return parsedUserID, true
}

func abortWithUnauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": message})
	c.Abort()
}
