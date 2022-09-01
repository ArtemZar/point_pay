docker pull mongo


docker run \
-d \
-- name mongoDB \
-p 27017:27017 \
mongo:latest

docker ps


rest на :8080
grpс на :8081
mongodb :27017

You can also generate code without forward compatibility by setting an option on protoc-gen-grpc-go plugin (source):
protoc --go-grpc_out=require_unimplemented_servers=false:.


Create new account
curl --location --request POST 'localhost:8080/create_account?Email' \
--header 'Content-Type: application/json' \
--data-raw '{
"email": "newexampl@mail.com"
}'


Generate wallet id
curl --location --request POST 'localhost:8080/generate_address' \
--header 'Content-Type: application/json' \
--data-raw '{
"id": "631067e026554071bf42dbf7"
}'

Update balance after deposit
curl --location --request POST 'localhost:8080/deposit' \
--header 'Content-Type: application/json' \
--data-raw '{
"id": "630fc63efba343acfb7768b6",
"amount": "2000.5"
}'

Update balance after withdrawal
curl --location --request POST 'localhost:8080/withdrawl' \
--header 'Content-Type: application/json' \
--data-raw '{
"id": "630fc63efba343acfb7768b6",
"amount": "100.99"
}'