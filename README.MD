## Description
This is a transaction-payment RESTful API project. The project is built with golang authentication, postgreSQL and JWT and implements a clean architecture. When a user authenticates logins, transfers between users can be made via a wallet with no maximum transfer fees. Then there is the account logout feature.

## Features
1. General Features
2. Authentication and Authorization
3. List of Transactions
4. Transfer
5. User Details
6. Log Transaksi
7. Database SQL
8. API documentation

## Database Design
<img src="./asset/Untitled.png" alt="e-wallet-ERD"/>

## API Request JSON File
### Source File: API/migration/req
```

## API Dokuemtation
```
[GIN-debug] GET    /v1/profile               --> GO-Payment/internal/controller.(*UserController).Profile-fm (3 handlers)
[GIN-debug] POST   /v1/login                 --> GO-Payment/internal/controller.(*AuthController).HandleLogin-fm (3 handlers)
[GIN-debug] POST   /v1/logout                --> GO-Payment/internal/controller.(*AuthController).Logout-fm (3 handlers)
[GIN-debug] GET    /v1/transactions          --> GO-Payment/internal/controller.(*TransactionController).GetTransactions-fm (3 handlers)
[GIN-debug] POST   /v1/auth-mid/transfer     --> GO-Payment/internal/controller.(*TransactionController).Transfer-fm (4 handlers)
[GIN-debug] POST   /v1/transfer              --> GO-Payment/internal/controller.(*TransactionController).Transfer-fm (3 handlers)
[GIN-debug] GET    /v1/log                   --> GO-Payment/internal/controller.(*LogController).GetAllLogs-fm (3 handlers)

```

#### Login

##### Request Method: GET
##### Request Path: 
```

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
```
#### Login

##### Request Method: POST
##### Request Path: 
```
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
```
#### Login

##### Request Method: POST
##### Request Path: 
```
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

```
#### GET Transactions

#####Request Method: GET
##### Respone :

Status Code 200 OK if successful
Status Code 500 Internal Server Error if an error occurs during the process
Status Code 404 Not Found if there are no transactions found

##### Request Path: 
```
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
```
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
```
#####Request Path: 
```
```json
http://127.0.0.1:8003/v1/transactions


```
```json

{
  "user_id": 3,
  "destination_id": "9876543210",
  "amount": 100,
  "history": "transfer 100 to user 1",
  "payment_method_type": "wallets"
}


{
  "data": {
    "id": 0,
    "user_id": 3,
    "amount": 100,
    "destination_id": "9876543210",
    "history": "transfer 100 to user 1",
    "created_at": "2023-04-19T10:36:49.279374928+07:00",
    "updated_at": "2023-04-19T10:36:49.279374988+07:00",
    "payment_method_type": "wallets"
  },
  "message": "Transfer successful",
  "status": "success"
}

```
#### GET LOG

```
#####Request Method: GET
```

```
#####Request Path: 
```

```
http://127.0.0.1:8003/v1/log

```
```
```json
{
  "status": "Success",
  "message": "logs retrieved successfully",
  "data": [
    {
      "id": 1,
      "user_id": 3,
      "event": "Transfer 100 from wallet 9876543210 to wallet 9876543210",
      "created_at": "2023-04-19T11:27:08.983195Z"
    },
    {
      "id": 2,
      "user_id": 3,
      "event": "Transfer 100 from wallet 9876543210 to wallet 9876543210",
      "created_at": "2023-04-19T11:40:45.939568Z"
    }
  ]
}