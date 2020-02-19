goredis-core goredis 实现

注意：

1. 错误信息统一由 core 包描述，其他过程直接往调用者返回
2. 参数解析由 server 包中的各个 rpc 服务函数来做，各取所需的参数，调用 core 包
3. 如果返回值中有 error 且不为 nil ,则原本的返回值为无意义的字符串（""）

# TODO: 

## 待实现命令

keys:

- [ ] del
- [ ] exist
- [ ] rename

strings:

- [ ] append
- [ ] incr incrby
- [ ] decr decrby
- [ ] setex
- [ ] setnx

lists:

- [ ] llen
- [ ] lpush,rpush
- [ ] lpop,rpop
- [ ] lrange
- [ ] rpoplpush
- [ ] lindex
- [ ] lrem

set:集合运算

- [ ] sadd
- [ ] spop
- [ ] srem

## 待实现功能

- [ ] 增加 key 过期功能
- [ ] 排序
- [ ] 集合运算