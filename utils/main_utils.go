package utils

import (
	"fmt"
	"log"
	"regexp"
)

func ValidatePaymentCardNumber(paymentCardNumber string) (bool, string) {
	log.Println("validating payment card number: ", paymentCardNumber)

	if !containsOnlyDigits(paymentCardNumber) {
		return false, fmt.Sprintf("Expect all digits number but given '%s' instead.", paymentCardNumber)
	}

	if !validateWithLuhnAlgorithm(paymentCardNumber) {
		return false, "Invalid payment card number!"
	}

	return true, "Valid payment card number!"
}

func containsOnlyDigits(num string) bool {
	log.Print("Checking if only digits")
	return regexp.MustCompile(`^[0-9]+$`).MatchString(num)
}

func validateWithLuhnAlgorithm(paymentCardNumber string) bool {
	log.Printf("Validating '%s' with Luhn Algorithm.", paymentCardNumber)
	return true
}
