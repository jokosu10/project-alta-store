# Create PRODUCTS

    POST products
    
Creates a new products to be stored inside Alta Store




## Parameters
### URI Parameters
None
### JSON Body Parameters
Field | Data Type | Required | Description
--- | --- | --- | ---
categories_id | int| Y | categories_id is used to group item based on categories in categories_table
name | string | Y | Name of the product 
description | string | Y | Description of the product
quantity | int | Y | Quantity of the product 
price | int | Y | Price of the product
Unit | int | Y | Unit of the product, for example : "pcs"

## Example
### Request

    POST https://baseurl/products

#### Request Body 
```json
{
        "categories_id":2,
        "name":"Yonex Racket 500",
        "description":"Fastest racket",
        "quantity":50,
        "price":3500000,
        "unit":"pcs"
}
```

### Response
``` json
{
    "code": "200",
    "message":"Success creating product",
    "status":"success",
    "data": {
        "ID": 5,
        "categories_id":2,
        "name":"Yonex Racket 500",
        "description":"Fastest racket",
        "quantity":50,
        "price":3500000,
        "unit":"pcs"
    }
}
```
