package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"os"
	"strings"
	"time"
)

type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.StandardClaims
}

// generateJWTKey creates a new 256-bit secret key for JWT.
func generateJWTKey() (string, error) {
	newKey := make([]byte, 32) // 256 bits
	if _, err := rand.Read(newKey); err != nil {
		return "", err
	}
	encodedKey := base64.StdEncoding.EncodeToString(newKey)
	return encodedKey, nil
}

func GetJwtKey() (string, error) {
	// try to get JWT Key by env
	jwtKeyString := os.Getenv("JWT_SECRET")

	// if not defined in env, try to get from file
	if jwtKeyString == "" {
		data, err := os.ReadFile("jwtkey")
		if err != nil || len(data) == 0 {
			zap.S().Infof("failed when read jwtKey: %v", err)
		} else {
			jwtKeyString = string(data)
			return jwtKeyString, nil
		}
	}

	// create if no JWT Key
	if viper.GetBool("JWT_SECRET_AUTO") && jwtKeyString == "" {
		zap.S().Errorf("start to create jwtKey")
		// generate
		encodedKey, err := generateJWTKey()
		if err != nil {
			zap.S().Errorf("unable to generate jwtKey: %v", err)
			return "", err
		}

		// write the secret
		err = os.WriteFile("jwtKey", []byte(encodedKey), 0600)
		if err != nil {
			zap.S().Errorf("failed to write key: %v", err)
			return "", err
		}
		zap.S().Infof("generated secret key: %s", encodedKey)
		jwtKeyString = encodedKey
	}

	return jwtKeyString, nil
}

func CreateToken(userID uuid.UUID) (string, error) {
	// set expire time
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// fetch JWT Key
	jwtKeyString, err := GetJwtKey()
	if err != nil {
		return "", err
	}
	jwtKey := []byte(jwtKeyString)

	// create the token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(jwtToken string, jwtKey []byte) (*jwt.Token, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtKeyString, err := GetJwtKey()
		if err != nil {
			zap.S().Errorf("failed to get JWT key: %v", err)
			return
		}

		jwtKey := []byte(jwtKeyString)

		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token not found"})
			ctx.Abort()
			return
		}

		token, err := ValidateToken(tokenString, jwtKey)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized: " + err.Error()})
			ctx.Abort()
			return
		}

		// Token is valid, set user ID to the context
		claims, ok := token.Claims.(*Claims)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token claims are not valid"})
			ctx.Abort()
			return
		}

		ctx.Set("userID", claims.UserID)
		ctx.Next()
	}
}
