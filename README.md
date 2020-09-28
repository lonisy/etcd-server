# etcd-server
简易配置中心

# 配置中心
对外提供一个接口，返回对应 application 的配置数据。

## 接口 
接口地址：config.ai-shanghai.com/api/etc
接口参数：app_key
返回参数：body 为 base64编码的 AES 加密数据。

## 项目逻辑
1、每个 application 可以包含多个 app_key, 每个 app_key 对应一个 16 位的秘钥。
2、所有数据存在 redis 中。

## Redis 缓存配置

```shell 
# 生成 app_key 或者 secret_key
date +%s |sha256sum |base64 |head -c 16;echo

# 创建新的 appkey 和对应的秘钥
redis-cli -h 127.0.0.1 -p 6379 -a qwe..123 hmset etc:app:app_key key value

redis-cli -h 127.0.0.1 -p 6379 -a qwe..123 hmset etc:app:OTI0Y2YyYWU0MzI4 app_name "kar_engine"
redis-cli -h 127.0.0.1 -p 6379 -a qwe..123 hmset etc:app:OTI0Y2YyYWU0MzI4 app_key "OTI0Y2YyYWU0MzI4"
redis-cli -h 127.0.0.1 -p 6379 -a qwe..123 hmset etc:app:OTI0Y2YyYWU0MzI4 secret_key "MDFjOWEwMzU5NDAx"


# 获取appkey list，hash 数据结构。KEY etc:category
redis-cli -h 127.0.0.1 -p 6379 -a qwe..123 hgetall etc:category
redis-cli -h 127.0.0.1 -p 6379 -a qwe..123 hgetall etc:kar_engine

# 根据appkey设置配置
redis-cli -h 127.0.0.1 -p 6379 -a qwe..123 set etc:$app_key 'config_json_string'
redis-cli -h 127.0.0.1 -p 6379 -a qwe..123 set etc:OTI0Y2YyYWU0MzI4 'config_json_string'

# 根据appkey获取配置
redis-cli -h 127.0.0.1 -p 6379 -a qwe..123 get etc:$app_key


```
