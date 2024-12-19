# calc
HTTP Calculator in Go

# How to start
You can simply run it by
```
go run cmd/main.go
```
or build
```
go build -o OUTPUT_BINARY cmd/main.go
```
and after that - run as standalone application

# Examples
## Correct request
```http
   ### Wrong HTTP Method. Expect code 500 and {"error": "Internal server error"}
   GET http://localhost/api/v1/calculate

   ### OK Expression. Expect code 200 and {"result": 6.000000}
   POST http://localhost/api/v1/calculate
   Content-Type: application/json

   {
     "expression": "2+2*2"
   }

   ### Empty or Incorrect expression. Expect code 422 and {"error": "Expression is not valid"}
   POST http://localhost/api/v1/calculate
   Content-Type: application/json

   {
     "expression": ""
   }

   ### Invalid JSON format. Expect code 500 and {"error": "Internal server error"}
   POST http://localhost/api/v1/calculate
   Content-Type: application/json

   {
     "expression": "
   }
```

You can do a simple test with curl like
```
curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
```