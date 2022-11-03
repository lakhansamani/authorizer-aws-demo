package middleware

import (
	"log"
	"os"
	"strings"

	"github.com/authorizerdev/authorizer-go"
	"github.com/gin-gonic/gin"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

// AuthMiddleware is used to authorize user for given api calls
func AuthorizeMiddleware() gin.HandlerFunc {
	authorizerURL := os.Getenv("AUTHORIZER_ENDPOINT")
	authorizerClientID := os.Getenv("AUTHORIZER_CLIENT_ID")

	defaultHeaders := map[string]string{}
	authorizerClient, err := authorizer.NewAuthorizerClient(authorizerClientID, authorizerURL, "", defaultHeaders)
	if err != nil {
		log.Fatal("Please set Authorizer environment variable correctly", err, authorizerClientID, authorizerURL)
	}

	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			respondWithError(c, 401, "Authorization Header Required")
			return
		}

		authSplit := strings.Split(authHeader, " ")
		if len(authSplit) != 2 {
			respondWithError(c, 401, "Invalid Authorization Header")
			return
		}

		if strings.ToLower(authSplit[0]) != "bearer" {
			respondWithError(c, 401, "Invalid Bearer Token")
			return
		}

		jwtToken := strings.TrimPrefix(authHeader, "Bearer ")

		res, err := authorizerClient.ValidateJWTToken(&authorizer.ValidateJWTTokenInput{
			TokenType: authorizer.TokenTypeIDToken,
			Token:     jwtToken,
		})
		if err != nil || !res.IsValid {
			log.Println("error from authorizer", err, res)
			respondWithError(c, 401, "Invalid JWT Token")
			return
		}

		c.Set("user_id", res.Claims["id"])
		c.Next()
	}
}
