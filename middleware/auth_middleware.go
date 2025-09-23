package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/IlhamLamp/cmty-core-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware menerima jwtSecret dari caller (main/routes)
func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if err := validateAuthHeader(authHeader); err != nil {
			utils.Error(ctx, http.StatusUnauthorized, err, "Unauthorized")
			ctx.Abort()
			return
		}

		tokenStr := strings.SplitN(authHeader, " ", 2)[1]
		token, err := parseJWTToken(tokenStr, jwtSecret)
		if err != nil {
			utils.Error(ctx, http.StatusUnauthorized, err, "Unauthorized")
			ctx.Abort()
			return
		}

		claims, err := extractClaims(token)
		if err != nil {
			utils.Error(ctx, http.StatusUnauthorized, err, "Unauthorized")
			ctx.Abort()
			return
		}

		setClaimsToContext(ctx, claims)
		ctx.Next()
	}
}

// validateAuthHeader checks the Authorization header format.
func validateAuthHeader(authHeader string) error {
	if authHeader == "" {
		return errorString("missing Authorization header")
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return errorString("invalid Authorization header")
	}
	return nil
}

// parseJWTToken parses and validates the JWT token.
func parseJWTToken(tokenStr, jwtSecret string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(jwtSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, errorString("invalid or expired token")
	}
	return token, nil
}

// extractClaims extracts and validates claims from the token.
func extractClaims(token *jwt.Token) (jwt.MapClaims, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errorString("invalid token claims")
	}
	if exp, ok := claims["exp"].(float64); ok {
		if time.Unix(int64(exp), 0).Before(time.Now()) {
			return nil, errorString("token expired")
		}
	}
	return claims, nil
}

// setClaimsToContext sets userID and role from claims to gin.Context.
func setClaimsToContext(c *gin.Context, claims jwt.MapClaims) {
	if sub, ok := claims["sub"].(string); ok {
		c.Set("userID", sub)
	}
	if role, ok := claims["role"].(string); ok {
		c.Set("role", role)
	}
}

// errorString is a simple error implementation for string messages.
type errorString string

func (e errorString) Error() string { return string(e) }
