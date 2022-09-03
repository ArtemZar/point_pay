# PointPay

## ТЗ Go developer
1. Написать два микро-сервиса:
- Название: Accounts
- Название: Банк
2. Микро-сервис Accounts, должен:
- Иметь серверную архитектуру и реализовывать в protobuf v3 следующее
  Структуру Account с полями: ID (генерить уникальный UUID при создании), walletID типа uint64, balance тип string
  Сервисы:
1. Create Account (создает новый аккаунт, walletID должен быть пустым, balance 0)
2. Get Accounts (возвращает accounts) stream
3. Generate Address (генерирует random UUID и записывает в walletID)
4. Deposit (добавляет сумму к balance, которая передаётся в request)  возвращает account
5. Withdrawal (вычитает сумму из balance, которая передаётся в request) возвращает account
   Repository/Store должен быть реализован на Mongo
   Logs на Zap
3. Микро-сервис Банк:
- Использует тот же protobuf
- Каждый сервис из Accounts преобразует в REST endpoints написанные на Gin. Т.е. является proxy сервисом для Accounts.
- Вызывает соответствующие сервисы rpc на микро-сервисе Accounts

## Use MongoDB via Docker

+ get image  
docker pull mongo

+ run container  
docker run -d --name mongoDB -p 27017:27017 mongo:latest

+ check container  
docker ps

## Default ports
+ rest localhost:8080
+ grpс localhost8081
+ mongodb localhost:27017

## Notes for generate proto
You can also generate code without forward compatibility by setting an option on protoc-gen-grpc-go plugin (source):
protoc --go-grpc_out=require_unimplemented_servers=false:.

## Tests endpoints by curl
+ Create new account  
curl --location --request POST 'localhost:8080/create_account?Email' \
--header 'Content-Type: application/json' \
--data-raw '{
"email": "newexampl@mail.com"
}'

+ Generate wallet id  
curl --location --request POST 'localhost:8080/generate_address' \
--header 'Content-Type: application/json' \
--data-raw '{
"id": "631067e026554071bf42dbf7" //example id
}'

+ Update balance after deposit  
curl --location --request POST 'localhost:8080/deposit' \
--header 'Content-Type: application/json' \
--data-raw '{
"id": "630fc63efba343acfb7768b6", //example id
"amount": "2000.5"
}'

+ Update balance after withdrawal  
curl --location --request POST 'localhost:8080/withdrawl' \
--header 'Content-Type: application/json' \
--data-raw '{
"id": "630fc63efba343acfb7768b6", //example id
"amount": "100.99"
}'

+ Get accounts  
  curl --location --request GET 'localhost:8080/get_accounts'

## TODO
+ tests
+ rest: returns of other statuses besides 200
+ init configs
+ refactoring