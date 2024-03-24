package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxpang93/credit-card-number-validator/utils"
)

func main() {
	router := gin.Default()

	router.GET("/ping", liveCheck)
	router.POST("/validate", validateHandler)

	router.Run(":8081")
	log.Print("haha")
}

func liveCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "pong",
	})
}

func validateHandler(c *gin.Context) {
	type creditCardInfo struct {
		CreditCardNumber string
	}

	var JsonBody creditCardInfo
	if err := c.BindJSON(&JsonBody); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "Unprocessable JSON Body",
		})
	}

	valid, msg := utils.ValidatePaymentCardNumber(JsonBody.CreditCardNumber)
	c.JSON(http.StatusOK, gin.H{
		"msg":   msg,
		"valid": valid,
	})
}
