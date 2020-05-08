# shortlink
golang 短地址
### 接口文档

> 1.生成短地址接口
```
curl --location --request POST 'http://dwz.murphyyi.com/api/shorten' \
--header 'Content-Type: application/json' \
--data-raw '{
    "url": "http://www.baidu.com",
    "expiration_in_minutes": 10000
}'
```

> 2.短地址还原接口
```
curl --location --request GET 'http://dwz.murphyyi.com/api/info?shortUrl=B'
```
```

> 3.短地址访问-重定向（307）

```
$ curl http://127.0.0.0.1:8080/8dxu
```
http://dwz.murphyyi.com/B
