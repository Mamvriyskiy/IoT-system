#!/bin/bash

nginx -c /Users/ivanmamvriyskiy/Desktop/web/nginx/nginx.conf -p /Users/ivanmamvriyskiy/Desktop/web/nginx -s reload                             

# Запускаем POST и GET запросы одновременно
ab -n 3000 -c 100 -p data.json -T 'application/json' -v 1 http://127.0.0.1:8081/api/v1/auth/sign-in & 
ab -n 3000 -c 100 -T 'application/json' -v 1 http://127.0.0.1:8081/api/v1/auth/sign-in &
wait  # Дожидаемся завершения обоих процессов

# с
# curl -X POST http://localhost:8081/api/v1/auth/sign-in -d @data.json -H "Content-Type: application/json" -v                