# Gin HTTP Server

## Run

```shell
go run cmd/apiserver.go -c configs/config.yaml
```

## Test

### create

```shell
curl -X POST -H "Content-Type: application/json" \
-d '{"metadata":{"name":"user99", "password":"admin"},"description":"admin user"}' \
http://127.0.0.1:8080/v1/users
```

### list

```shell
curl -X GET http://127.0.0.1:8080/v1/users
```

### get

```shell
curl -X GET http://127.0.0.1:8080/v1/users/user99
```

### update

```shell
curl -X PUT -H "Content-Type: application/json" \
-d '{"metadata":{"name":"user99"},"nickname":"xxx"}' \
http://127.0.0.1:8080/v1/users/user99
```

### delete

```shell
curl -X DELETE http://127.0.0.1:8080/v1/users/user99
```

### 