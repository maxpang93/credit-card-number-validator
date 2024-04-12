package utils

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
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

/*
Luhn Algorithm:
1. Given the full card number, remove the last number as the check digit
2. Now, with the remaining numbers (payload) as an array of digits, reverse the array
3. Double the digit at even position and sum up the resulting digit (e.g. 18 = 1+8 = 9)
4. Reduce the array by summing up all the elements (totalSum)
5. Calculate the check digit by this formula: `checkDigit = (10 - (totalSum mod 10)) mod 10`
6. Match with the last digit removed at step 1.

Example:
card number: 17893729974
last digit = 4
payload = 1789372997
reversedArray = [7,9,9,2,7,3,9,8,7,1]
doubleAtEvenPosition = [14,9,18,2,14,3,18,8,14,1]
finalArray = [5,9,9,2,5,3,9,8,5,1]
totalSum = 5 + 9 + 9 + 2 + 5 + 3 + 9 + 8 + 5 + 1 = 56
checkDigit = (10 - (56 mod 10)) mod 10 = 4
*/
func validateWithLuhnAlgorithm(paymentCardNumber string) bool {
	log.Printf("Validating '%s' with Luhn Algorithm.", paymentCardNumber)

	// convert the string to array of int
	var digitArray []int
	for _, char := range paymentCardNumber {
		num, _ := strconv.Atoi(string(char))
		digitArray = append(digitArray, num)
	}

	// extract last digit
	lastDigit := digitArray[len(digitArray)-1]

	// remove last digit from array
	digitArray = digitArray[:len(digitArray)-1]

	// reverse array
	for i, j := 0, len(digitArray)-1; i < j; i, j = i+1, j-1 {
		digitArray[i], digitArray[j] = digitArray[j], digitArray[i]
	}

	// double value at even position and sum the resulting digits
	for idx, digit := range digitArray {
		if idx%2 == 0 {
			digitArray[idx] = sumDigit(digit * 2)
		}
	}

	// calculate totalSum
	var totalSum int
	for _, value := range digitArray {
		totalSum += value
	}

	checkDigit := (10 - (totalSum % 10)) % 10

	return checkDigit == lastDigit
}

func sumDigit(digit int) int {
	digitString := strconv.Itoa(digit)
	var sum int
	for _, char := range digitString {
		integer, _ := strconv.Atoi(string(char))
		sum += integer
	}
	return sum
}
