# credit-card-number-validator
Just a simple Golang project to gain proficiency.

## The web server only has two endpoints:
1. `/ping` to check if the server is working.
2. `/validate` to validate a credit card number using Luhn Algorithm

To start up the web server, either 
-  At root, run:
   ```bash
    go run main.go
   ```
- Or use Docker:
   - Build image:
   ```bash
   docker build -t credit_card_number_validator .
   ```
   - Run container:
   ```bash
   docker run -p 8081:8080 credit_card_number_validator
   ```

## Example:
1. `/ping`
    ```bash
    curl -X GET "localhost:8081/ping"
    ```
2. `/validate`
    ```bash
     curl -X POST "localhost:8081/validate" -H "Content-Type: application/json" -d '{"creditCardNumber": "17893729974"}'
    ```



