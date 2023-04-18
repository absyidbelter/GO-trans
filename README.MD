## Description
This is a transaction-payment RESTful API project. The project is built with golang authentication, postgreSQL and JWT and implements a clean architecture. When a user authenticates logins, transfers between users can be made via a wallet with no maximum transfer fees. Then there is the account logout feature.

## Features
1. General Features
1. Authentication and Authorization
1. List of Transactions
1. Transfer
1. User Details
1. Database SQL
1. API documentation

## Database Design
<img src="./asset/Untitled.png" alt="e-wallet-ERD"/>
## API Documentation
Transaction-payment-API/migration/req

## API Dokuemtation
POST /login
Handles user login and returns a JWT token if the user is authenticated.


#### GET PROFILE

#####Request Method: GET
#####Request Path: 

```json
http://127.0.0.1:8003/v1/profile
```
```json
{
  "status": "Success",
  "data": [
    {
      "id": 1,
      "name": "bima",
      "email": "bima@example.com",
      "password": "password123"
    },
    {
      "id": 2,
      "name": "budiatun",
      "email": "budi@example.com",
      "password": "mypassword"
    },
    {
      "id": 3,
      "name": "markzus",
      "email": "markzus@example.com",
      "password": "securepassword"
    }
  ]
}

#### Login

##### Request Method: POST
##### Request Path: 

```json
http://127.0.0.1:8003/v1/login
```
```json
{
  "status": "Success",
  "data": [
    {
      "id": 1,
      "name": "bima",
      "email": "bima@example.com",
      "password": "password123"
    },
    {
      "id": 2,
      "name": "budiatun",
      "email": "budi@example.com",
      "password": "mypassword"
    },
    {
      "id": 3,
      "name": "markzus",
      "email": "markzus@example.com",
      "password": "securepassword"
    }
  ]
}

#### Login

##### Request Method: POST
##### Request Path: 

```json
http://127.0.0.1:8003/v1/login
```
```json
{
    "name": "bima",
    "password": "password123"
}

```
```json
{
  "status": "Success",
  "message": "Login success",
  "data": {
    "expired": "0s",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODE4NTU2NDksInVzZXJfaWQiOjF9.-Y-BAFUNVygKFnLuSGHAfTXOypYVn0A3iyw7pTFCjks"
  }
}


#### GET Transactions

#####Request Method: GET
##### Respone :

Status Code 200 OK if successful
Status Code 500 Internal Server Error if an error occurs during the process
Status Code 404 Not Found if there are no transactions found

#####Request Path: 

```json
http://127.0.0.1:8003/v1/transactions
```
```json
{
  "status": "success",
  "data": [
    {
      "id": 1,
      "user_id": 1,
      "destination_id": 2,
      "amount": 1000000,
      "payment_method_type": "GOPAY",
      "status": "SUCCESS",
      "created_at": "2023-04-18T13:50:20Z",
      "updated_at": "2022-04-18T13:50:20Z"
    },
    {
      "id": 2,
      "user_id": 1,
      "destination_id": 3,
      "amount": 2000000,
      "payment_method_type": "wallet",
      "status": "SUCCESS",
      "created_at": "2023-04-18T13:51:30Z",
      "updated_at": "2023-04-18T13:51:30Z"
    }
  ]
}

###### In case of failure

{
  "status": "error",
  "message": "Failed to get transactions"
}

#### Transfer
##### Request Method: POST
##### Respone :

Status Code 200 OK if successful
Status Code 500 Internal Server Error if an error occurs during the process
Status Code 404 Not Found if there are no transactions found

#####Request Path: 

```json
http://127.0.0.1:8003/v1/transactions
```
```json
{
    "user_id": 2,
    "destination_id": 1234,
    "amount": 3000,
    "status": "succes",
    "created_at": "2023-04-18T13:51:30Z"
}
