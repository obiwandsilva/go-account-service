package controllers

import "github.com/gin-gonic/gin"

// GetAccountInfo is a handler for the route
func GetAccountInfo(ctx *gin.Context) {
	accountID := ctx.Param("accountId")

	ctx.Header("content-type", "application/json")
	ctx.JSON(200, gin.H{
		"id":   accountID,
		"name": "Wand",
		"age":  24,
	})
}
