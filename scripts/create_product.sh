#!/bin/bash

curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Product 1",
    "description": "Test product",
    "images": ["https://example.com/img1.jpg"]
  }'
