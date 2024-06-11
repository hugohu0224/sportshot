package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"sportshot/internal/webserver/auth"
	"sportshot/pkg/utils/global"
	"sportshot/pkg/utils/models"
)

func GetLoginPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.tmpl", gin.H{})
}

func GetRegisterPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "register.tmpl", gin.H{})
}

func Register(ctx *gin.Context) {
	// binding model
	var u models.User
	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// check if exists
	err := global.MySQLClient.Where("username = ?", u.Username).First(&models.User{}).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		zap.S().Debug("user name does not exist, allow to register user")
	} else {
		zap.S().Errorf("register user failed, username: %s, error: %v", u.Username, zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("register user failed, username: %s, error: %v", u.Username, zap.Error(err))})
		return
	}

	// start to register
	hashedPassword, err := HashPassword(u.Password)
	err = global.MySQLClient.Create(&models.User{
		Model:        gorm.Model{},
		ID:           uuid.UUID{},
		Username:     u.Username,
		Password:     hashedPassword,
		RefreshToken: "",
		Active:       true,
	}).Error
	if err != nil {
		zap.S().Errorf("register user failed, username: %s, error: %v", u.Username, zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("register user failed, username: %s, error: %v", u.Username, zap.Error(err))})
		return
	}

	// success
	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("register user %s successfully", u.Username)})

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func AuthenticateLogin(ctx *gin.Context) {
	// binding model
	var lc models.LoginCredentials
	if err := ctx.ShouldBindJSON(&lc); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// check if exists
	var userInfo models.User
	err := global.MySQLClient.Where("username = ?", lc.Username).First(&userInfo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zap.S().Errorf("username %s does not exist", lc.Username)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("username or password incorrect, username: %s, error: %v", lc.Username, zap.Error(err))})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("login failed : %v", zap.Error(err))})
			return
		}
	}

	// validate
	if !CheckPasswordHash(lc.Password, userInfo.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("invalid password, username: %s, error: %v", lc.Username, zap.Error(errors.New("invalid password")))})
		return
	}

	// success
	jwtToken, err := auth.CreateToken(userInfo.ID)
	ctx.JSON(http.StatusOK, gin.H{
		"message":  fmt.Sprintf("login successfully, username: %s", lc.Username),
		"jwtToken": jwtToken,
	})
}

func Validate(ctx *gin.Context) {
	jwtToken := ctx.Query("token")
	if jwtToken == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Token is required"})
		return
	}

	jwtKeyString, err := auth.GetJwtKey()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	jwtKey := []byte(jwtKeyString)

	_, err = auth.ValidateToken(jwtToken, jwtKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "valid": false})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"valid": true})
}
