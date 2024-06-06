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
	"sportshot/pkg/utils/global"
	"sportshot/pkg/utils/models/auth"
	"sportshot/pkg/utils/models/user"
)

func GetLoginPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.tmpl", gin.H{})
}

func GetRegisterPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "register.tmpl", gin.H{})
}

func Register(ctx *gin.Context) {
	// binding model
	var u user.User
	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// check if exists
	err := global.MySQLClient.Where("username = ?", u.Username).First(&user.User{}).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		zap.S().Debug("user name does not exist, allow to register user")
	} else {
		zap.S().Errorf("register user failed, username: %s, error: %v", u.Username, zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("register user failed, username: %s, error: %v", u.Username, zap.Error(err))})
		return
	}

	// start to register
	hashedPassword, err := HashPassword(u.Password)
	err = global.MySQLClient.Create(&user.User{
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
	var lc auth.LoginCredentials
	if err := ctx.ShouldBindJSON(&lc); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// check if exists
	var userInfo user.User
	err := global.MySQLClient.Where("username = ?", lc.Username).First(&userInfo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zap.S().Errorf("username %s does not exist", lc.Username)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("username or password incorrect, username: %s, error: %v", lc.Username, zap.Error(err))})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("login failed : %v", zap.Error(err))})
		}
	}

	// validate
	if !CheckPasswordHash(lc.Password, userInfo.Password) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("invalid password, username: %s, error: %v", lc.Username, zap.Error(errors.New("invalid password")))})
	}

	// success
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("login successfully, username: %s", lc.Username),
	})

}
