# Alta Store REST API (ASRA)

## Overview

ASRA is REST-API specifically build to support online store system of ALTA. ASRA provides functionality that allows developers to behave either as customer or admin. ASRA is created using Golang, Gorm, and MYSQL as database.

## Tutorial
We provide another documentation using swagger(PUT_LINK_HERE) to make it easy to understand the basic of API 

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
* `401` `Unauthorized` The supplied API credentials are invalid
* `403` `Forbidden` The credentials provided do not have permission to access the requested resource
* `404` `Not found` An attempt was made to access a resource that does not exist in the API
* `405` `Method not allowed` The resource being accessed doesn't support the method specified (GET, POST, etc.).
* `500` `Server Error` An error on the server occurred


## Resources

### Admin
- **[<code>POST</code> CREATE Products](/api_docs/POST_PRODUCTS.md)**
