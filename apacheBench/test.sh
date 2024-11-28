#!/bin/bash

nginx -c /Users/ivanmamvriyskiy/Desktop/web/nginx/nginx.conf -p /Users/ivanmamvriyskiy/Desktop/web/nginx -s reload                             

# Запускаем POST и GET запросы одновременно
# ab -n 12 -c 2 -p data.json -T 'application/json' -g out.data -v 1 http://127.0.0.1:8081/api/v1/auth/sign-in & 
ab -n 12 -c 3 -T 'application/json' -g out2.data -v 1 http://127.0.0.1:8081/api/v1/api/devices &
wait  # Дожидаемся завершения обоих процессов


# curl -X POST http://localhost:8081/api/v1/auth/sign-in -d @data.json -H "Content-Type: application/json" -v                
