# Alta Store REST API (ASRA)

## Overview

ASRA is REST-API specifically build to support online store system of ALTA. ASRA provides functionality that allows developers to behave either as customer or admin. ASRA is created using Golang, Gorm, and MYSQL as database.

The MVP of ASRA is :

* Login and register customers
* Customers can view list product based on product category
* Customers can add product to shopping cart
* Customers can see a list of products that have been added to the shopping cart
* Customers can delete the product list from the shopping cart
* Customers can checkout and make payment transactions

And this ERD of ASRA is like this :

![ERD of ASRA](./api_docs/erd.png)

## Tutorial

We provide another documentation using swagger to make it easy to understand the basic of API.

How to run this project :

1. git clone `https://github.com/jokosu10/project-alta-store.git`
2. cd `project-alta-store` and run `go install`
3. setup envirotment, rename file `.env.example` to `.env`. And setup this variabel envirotment.
4. after finished install this module, run command `go run swagger.go` for running this documentation to browser like this `localhost:8080/swaggerui`
5. run command `go run main.go` to running this server like this `localhost:8000`
6. run command `go test ./controller/ -cover` for running unit testing

## HTTP requests

There are 4 basic HTTP requests that you can use in this API:

* `POST` Create a resource
* `PUT` Update a resource
* `GET` Get a resource or list of resources
* `DELETE` Delete a resource

## HTTP Responses

Each response will include a code(repsonse code),message,status and data object that can be single object or array depending on the query.

## HTTP Response Codes

Each response will be returned with one of the following HTTP status codes:

* `200` `OK` The request was successful
* `400` `Bad Request` There was a problem with the request (security, malformed, data validation, etc.)
* `404` `Not found` An attempt was made to access a resource that does not exist in the API
* `500` `Server Error` An error on the server occurred

## API Documentation

### /register

#### POST
##### Summary

register customer

##### Description
passing customer data to register customer

### Parameter

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| username | body | customer usernmae | Yes | string
| email | body | customer email | Yes | string |
| password | body | customer password | Yes | string |

##### Responses
| Code | Description | 
| ---- | ----------- |
| 200 | success register customer |
| 500 | error registration customer | 

#### Request Body Parameter Example
```json
{
    "username":"lisa233",
    "email":"lisa@gmail5.com",
    "password":"312457"
}
```

#### Response Body Example
```json
{    
    "code": 200,
    "message": "success register customers",
    "status": "success"
}
```

### /login

#### POST
##### Summary

login as customer

##### Description
passing email and password to login as customer

####  Parameter
| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| email | body | customer email | Yes | string |
| password | body | customer password | Yes | string |

##### Responses
| Code | Description | 
| ---- | ----------- |
| 200 | success login customer |
| 400 | error login(invalid email/password) | 

#### Request Body Parameter Example
```json
{
   "email":"susi@gmail.com",
   "password":"john123"
}
```

#### Response Body Example
```json
{    
    "code": 200,
    "message": "success login customer",
    "status": "success",
    "data": {
        "id": 1,
        "email": "susi@gmail.com",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MTczNDg0ODgsInVzZXJJZCI6MX0.C4pvmawkwNP-SVJlIA0qbTct3qen4PPbhtHK9XRRpts"
    }
}
```
### /customers

### /customers/{id}

#### PUT
##### Summary

Update Customer Data

##### Description
passing several data to update  customer data 

####  Parameter
| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | customer id | Yes | string |
| email | body | customer email | No | string |
| adresss | body | customer address | No | string |
| bank_name | body | customer email | No | string |
| bank_account_number | body | customer address | No | string |

##### Responses
| Code | Description | 
| ---- | ----------- |
| 200 | success update customer data |
| 400 | error update data| 

#### Request Body Parameter Example
```json
{
    "email":"32ibu@gmail.com",
    "address": "jalan ahmad yani",
    "bank_name": "BCA",
    "bank_account_number": "23232323"
}
```

#### Response Body Example
```json
{    
    "code": 200,
    "message": "success update profil customer",
    "status": "success"
}
```

### /products

### /products?category={id}

##### GET
##### Summary

Get products

##### Description
Get all product based on category, if no id is passed then it will return all product in the store

#### Parameter
| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | query | category id | No | int |

##### Responses
| Code | Description | 
| ---- | ----------- |
| 200 | success get products |
| 400 | error bad request | 


#### Response Body Example
```json
{    
    "code": 200,
    "message": "success get product",
    "status": "success",
    "data":  [
        {
            "id": 1,
            "name": "Yonex Zerox",
            "categories_id": 4,
            "description": "Raket tercepat di Asia",
            "quantity": 30,
            "price": 45000,
            "unit": "pcs",
            "CreatedAt": "2021-03-28T20:01:40.21+07:00",
            "UpdatedAt": "2021-03-28T20:01:40.21+07:00"
        }
    ]
}
```

### /cartitems

#### POST
##### Summary

Adding cartitems to carts

##### Description
Adding product that customer wants to buy into their cart

####  Parameter
| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| carts_id | body | customer cart id | Yes | int |
| products_id | body | product id | Yes | int |
| quantity | body | quantity of product | Yes | int |

##### Responses
| Code | Description | 
| ---- | ----------- |
| 200 | success insert cartitems |
| 400 | invalid input| 
| 500 | product/cart not exist|

#### Request Body Parameter Example
```json
{
    "carts_id":4,
    "products_id":9,
    "quantity":3
}
```

#### Response Body Example
```json
{    
    "code": 200,
    "status": "success",
    "message": "success insert cartitems"
}
```

#### /cartitems?cart={cartid}

#### GET
##### Summary

get all cartitems inside specific cart

##### Description
get all cartitems that is inside usercart

####  Parameter
| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| carts_id | query | customer cart id | Yes | int |

##### Responses
| Code | Description | 
| ---- | ----------- |
| 200 | success get cartitems |
| 400 | bad request| 



#### Response Body Example
```json
{    
    "code": 200,
    "status": "success",
    "message": "Success Get Cartitems",
    "data": [
        {
            "id": 29,
            "carts_id": 4,
            "products_id": 9,
            "name": "Yonex e545",
            "price": 45000,
            "quantity": 3
        }
    ]
}
```

#### /cartitems/:cartitemid

#### DELETE
##### Summary

Delete specific cartitem

##### Description
Delete specific cartitem based on cartitemid

####  Parameter
| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| cartitemid | path |  id of cartitem to be deleted | Yes | int |

##### Responses
| Code | Description | 
| ---- | ----------- |
| 200 | success delete cartitems |
| 400 | error deleting cartitems(wrong cartitem id etc)| 



#### Response Body Example
```json
{    
     "code": 200,
     "message": "cartitems succesfully deleted",
     "status": "success"
}
```

### /orders

#### POST
##### Summary

Creating order from customer

##### Description
When customer checkout order and payment will be created, and customer cart item will be moved to checkout item

####  Parameter
| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| customers_id | body | customer id | Yes | int |
| courier_id | body | courier id | Yes | int |
| payment_method | body | method of payment | Yes | string |
| payment_start_date | body | date in format(yyyy-mm-dd hh:mm:ss)| Yes | string |
| payment_end_date | body | date in format(yyyy-mm-dd hh:mm:ss)| Yes | string |
| payment_status | body | status of payment | Yes | string |
| payment_amount | body | total payment amount | Yes | int |


##### Responses
| Code | Description | 
| ---- | ----------- |
| 200 | success create order |
| 400 | bad request(customer not exist,etc)| 

#### Request Body Parameter Example
```json
{
    "customers_id":4,
     "couriers_id":1,
     "payment_method":" BCA",
     "payment_start_date":"2021-03-29 10:42:44.710",
     "payment_end_date" : "2021-03-30 10:42:44.710",
     "payment_status": "waiting",
     "payment_amount" : 345000
}
```

#### Response Body Example
```json
{    
    "code": 200,
    "message": "success insert order",
    "status": "success"
}
```
### /payments/:id

#### PUT
##### Summary

Update payment status

##### Description
Update payment status when customer finishing their payment for the product

####  Parameter
| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | payment id that needs to be updated | Yes | int |

##### Responses
| Code | Description | 
| ---- | ----------- |
| 200 | success update payment |
| 400 | error update payment(payment id not exist etc)| 


#### Response Body Example
```json
{    
    "code": 200,
    "message": "success update payments",
    "status": "success"
}
```
