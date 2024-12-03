package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
)

// AllowedHostsMiddleware checks if the request host is allowed
func AllowedHostsMiddleware(allowedHosts []string) gin.HandlerFunc {
    return func(c *gin.Context) {
        requestHost := c.Request.Host
        for _, allowedHost := range allowedHosts {
            if strings.EqualFold(requestHost, allowedHost) {
                c.Next() // Proceed to the next handler
                return
            }
        }
        // If no match, reject the request
        c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
            "error": "Forbidden: Host not allowed",
        })
    }
}
