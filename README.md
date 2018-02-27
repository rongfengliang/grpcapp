# grpc with xservice for better app build

## how ro run  local
- start server
```bash
cd grpcdemoserver
go run main.go
```
- client connect to server
```bash
cd grpcdemoclient // you must add hosts with 127.0.0.1 server
go run main.go
```

## run with docker-compose 
```bash
docker-compose build 
docker-compose up -d
docker-compose logs cli // you will see the call result
```