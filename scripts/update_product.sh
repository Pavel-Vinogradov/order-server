#!/bin/bash

curl -X PUT http://localhost:8080/products/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Product 123",
    "description": "Test product 123",
    "images": ["https://example.com/img1.jpg"]
  }'
