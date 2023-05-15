# user-balance-api

## Microservice for working with the user's balance.

## Running the application

```sh
docker-compose up --build
```

### Running the application without Docker

```sh
go run .
```

### Running tests (app must be running)

```sh
go test .
```

#

## HTTP Methods

```
"GET" / — Checking the server connection

    example: 
        "GET" :8080/
```
#
```
"GET" /ping — Checking the server connection

    example: 
        "GET" :8080/ping
```
#

```
"GET" /users/balance — Gets the user's balance

    example: 
        "GET" :8080/users/balance
```
```json
{
    "user_id" : "ec6761fa-4b02-4e93-a213-8fa96eb44d15"
}
```
#

```
"POST" /users/create — Create a new user

    example: 
        "POST" :8080/users/create
```

```json
{
    "user_id" : "ec6761fa-4b02-4e93-a213-8fa96eb44d15"
    "amount"   : "55.55"
}
```

#
```
"PATCH" /users/deposit — Depositing funds to the user's balance

    example: 
        "PACTCH" :8080/users/deposit
```
```json
{
    "id" : "ec6761fa-4b02-4e93-a213-8fa96eb44d15",
    "amount" : "1000"
}
```
#

```
"PATCH" /users/withdraw — Withdrawal of funds from the user's balance

    example: 
        "PATCH" :8080/users/withdraw
```
```json
{
    "id" : "ec6761fa-4b02-4e93-a213-8fa96eb44d15",
    "amount" : "1000"
}
```
#

```
"PATCH" /users/transfer — Transferring funds from one user's balance to another user's balance

    example: 
        "PATCH" :8080/users/transfer
```
```json
{
	"from_user_id" : "ec6761fa-4b02-4e93-a213-8fa96eb44d15",
	"to_user_id" : "ec6741fa-4b02-4e03-a303-0fa96eb15d15",
    "amount" : "1000"
}
```
#
