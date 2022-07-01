package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("GetFoods")
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("GetFood")
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("CreateFood")
	}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("UpdateFood")
	}
}
