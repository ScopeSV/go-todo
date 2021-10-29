package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func isBearer(authToken string) bool {
	if authToken != "" {
		return strings.Contains(authToken, "Bearer ")
	}
	return false
}

func addCurrentUserToReq(c *gin.Context, token string) {
	c.Set("currentUser", token)
}

func extractJWTFromToken(t string) string {
	return strings.Split(t, "Bearer ")[1]
}

func IsAuthenticated(c *gin.Context) {
	authToken := c.GetHeader("Authorization")
	isBearer := isBearer(authToken)
	if isBearer {
		jwt := extractJWTFromToken(authToken)
		// Sjekke jwt her
		addCurrentUserToReq(c, jwt)
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "INVALID TOKEN"})
		return
	}
}
