# goredis 

使用 go 做 redis 的简易实现

## 功能

可客户端连接，可 Web 端直接使用

客户端与服务端：三种连接方式：tcp,http,rpc

涉及知识点：

1. tcp，http ,rpc net 包
2. 命令行操作 flag 包
3. 内存管理
4. 数据结构，string,list,set,zset,hash(先从string开始)
5. 持久化 os 包
6. 集群

先只支持字符串存储，把框架搭起来，后期再增加支持的数据结构。

## TODOs

### server

- [ ] 命令拆解