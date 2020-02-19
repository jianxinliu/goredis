# goredis 

使用 go 做 redis 的简易实现（仅实现部分命令），用于学习 Go 语言。

## 功能

可客户端连接，可 Web 端直接使用

先只支持字符串存储，把框架搭起来，后期再增加支持的数据结构。

## todo

web 端访问

## How to

### 1. 运行 goredis 服务器

```sh
go run ./goredis/app.go
```

### 2. 运行客户端

```sh
go run ./goredis/client/client.go
```

## 支持的命令

- [x] SET
- [x] GET
- [x] KEYS