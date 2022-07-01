package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Users")
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("GetUser")
	}
}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Signup")
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Login")
	}
}

func HashPassword(password string) string {

}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {

}
