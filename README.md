## Installation instruction (required Docker)
```docker-compose up -d```
 
Default port is 8080, so ensure your 8080 port is available or change to your custom port on `docker-compose.yml` file

Backend APIs are relied on RESTful style
## 1. Register
Method: `POST`

URL: `/api/register`

Payload:
```json
{
  "email": "huuthuan.nguyen@hotmail.com",
  "password": "123@pass!word"
}
```
Response:
```json
{
    "status": 1,
    "messages": [
        "Successful."
    ],
    "data": {
        "id": 1,
        "email": "huuthuan.nguyen@hotmail.com",
        "is_active": true,
        "create_at": "2022-05-18T03:40:59.167725305Z",
        "updated_at": "2022-05-18T03:40:59.167825888Z"
    }
}
```

## 2. Login
Method: `POST`

URL: `/api/auth/login`

Payload:
```json
{
  "email": "huuthuan.nguyen@hotmail.com",
  "password": "123@pass!word"
}
```
Response:
```json
{
    "status": 1,
    "messages": [
        "Successful."
    ],
    "data": {
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Imh1dXRodWFuLm5ndXllbkBob3RtYWlsLmNvbSIsImV4cCI6MTY1MjkzMTcwM30.H5jjQzsbgXuEuK1XQT5t4wf6dB99TkygOxkFr-KY36o",
        "expired_at": "2022-05-19T03:41:43.197707881Z"
    }
}
```

## 3. Add product
Method: `POST`

URL: `/api/products`

Payload:
```json
{
  "sku": "MAC-001",
  "name": "Macbook Pro 14 inches",
  "quantity": 100,
  "price": 50000,
  "unit": "box",
  "status": 1
}
```
Response
```json
{
    "status": 1,
    "messages": [
        "Successful."
    ],
    "data": {
        "id": 1,
        "sku": "MAC-001",
        "name": "Macbook Pro 14 inches",
        "quantity": 100,
        "price": 50000,
        "unit": "box",
        "status": 1,
        "created_by": 2,
        "create_at": "2022-05-18T03:42:20.611294259Z",
        "updated_at": "2022-05-18T03:42:20.611327676Z"
    }
}
```

## 4. Update product
Method: `PUT`

URL: `/api/products/:id`

Payload:
```json
{
  "sku": "MAC-001",
  "name": "Macbook Pro 16 inches",
  "quantity": 200,
  "price": 20000,
  "unit": "box",
  "status": 1
}
```
Response:
```json
{
    "status": 1,
    "messages": [
        "Successful."
    ],
    "data": {
        "id": 1,
        "sku": "MAC-001",
        "name": "Macbook Pro 16 inches",
        "quantity": 200,
        "price": 20000,
        "unit": "box",
        "status": 1,
        "created_by": 2,
        "create_at": "2022-05-18T03:42:20.611294Z",
        "updated_at": "2022-05-18T03:43:25.03941Z"
    }
}
```

## 5. Delete product
Method: `DELETE`

URL: `/api/products/:id`

Payload:
```json

```
Response
```json
```

## 6. Product list
Method: `GET`

URL: `/api/products`

Payload:
```json

```
Response:
```json
{
    "status": 1,
    "messages": [
        "Successful."
    ],
    "data": {
        "items": [
            {
                "id": 1,
                "sku": "MAC-001",
                "name": "Macbook Pro 16 inches",
                "quantity": 200,
                "price": 20000,
                "unit": "box",
                "status": 1,
                "created_by": 2,
                "create_at": "2022-05-18T03:42:20.611294Z",
                "updated_at": "2022-05-18T03:43:25.03941Z"
            },
            {
            "id": 2,
            "sku": "IP-001",
            "name": "iPhone 13 Pro",
            "quantity": 100,
            "price": 10000,
            "unit": "box",
            "status": 1,
            "created_by": 2,
            "create_at": "2022-05-18T03:42:20.611294Z",
            "updated_at": "2022-05-18T03:43:25.03941Z"
            }
        ]
    }
}
```

## 7. Product list by SKU
Method: `GET`

URL: `/api/products?sku=MAC-001`

Payload:
```json

```
Response:
```json
{
    "status": 1,
    "messages": [
        "Successful."
    ],
    "data": {
        "items": [
            {
                "id": 1,
                "sku": "MAC-001",
                "name": "Macbook Pro 16 inches",
                "quantity": 200,
                "price": 20000,
                "unit": "box",
                "status": 1,
                "created_by": 2,
                "create_at": "2022-05-18T03:42:20.611294Z",
                "updated_at": "2022-05-18T03:43:25.03941Z"
            }
        ]
    }
}
```

## Postman Collection
`https://www.getpostman.com/collections/6e996b399b8afa638cbd`