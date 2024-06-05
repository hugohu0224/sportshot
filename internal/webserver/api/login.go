package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sportshot/pkg/utils/models/auth"
)

func GetLoginPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.tmpl", gin.H{})
}

func Register(ctx *gin.Context) {

}

func ValidateLogin(ctx *gin.Context) {
	var lc auth.LoginCredentials
	if err := ctx.ShouldBindJSON(&lc); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("username: %v, password: %v", lc.Username, lc.Password),
	})
	//if validateCredentials(username, password) {
	//	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	//} else {
	//	// This will cause response.ok to be false on the client side
	//	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	//}

}
