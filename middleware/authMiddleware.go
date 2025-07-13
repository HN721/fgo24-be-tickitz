package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func VerifyToken(tokens string) (*jwt.Token, error) {
	godotenv.Load()
	return jwt.Parse(tokens, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
}
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.Next()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header required",
			})
			return
		}
		splitToken := strings.Split(authHeader, " ")
		if len(splitToken) != 2 || splitToken[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization format",
			})
			return
		}
		tokenStr := splitToken[1]
		token, err := VerifyToken(tokenStr)
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired token",
			})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if idFloat, ok := claims["id"].(float64); ok {
				userID := int(idFloat)
				c.Set("userID", userID)
			}
		}

		c.Next()
	}
}
