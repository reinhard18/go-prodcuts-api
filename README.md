# go-prodcuts-api

# Register a new user

````bash
curl -X POST http://localhost:8080/register \
 -H "Content-Type: application/json" \
 -d '{"username": "testuser", "password": "testpass123"}'```

# Login and get JWT token
```bash
curl -X POST http://localhost:8080/login \
 -H "Content-Type: application/json" \
 -d '{"username": "testuser", "password": "testpass123"}'```

# Get all products
```bash
curl -X GET http://localhost:8080/products \
 -H "Authorization: Bearer <token>"```

# Get single product
```bash
curl -X GET http://localhost:8080/products/1 \
 -H "Authorization: Bearer <token>"```

# Create new product
```bash
curl -X POST http://localhost:8080/products \
 -H "Authorization: Bearer <token>" \
 -H "Content-Type: application/json" \
 -d '{"name": "New Product", "price": 29.99}'```

# Update product
```bash
curl -X PUT http://localhost:8080/products/1 \
 -H "Authorization: Bearer <token>" \
 -H "Content-Type: application/json" \
 -d '{"name": "Updated Product", "price": 39.99}'```

# Delete product
```bash
curl -X DELETE http://localhost:8080/products/1 \
 -H "Authorization: Bearer <token>"```
````
