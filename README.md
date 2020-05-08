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

```
> POST /api/shorten
Content-Type: application/json;charset=utf-8
Params：
  {
    "url": "https://www.baidu.com",
    "expiration_in_minutes": 100
  }
Response：
  {
      "code": "0",
      "errMsg": "OK",
      "data": {
          "shortUrl": "8dxu",
          "longUrl": "https://www.baidu.com"
      }
  }
```

> 2.短地址还原接口
```
curl --location --request GET 'http://dwz.murphyyi.com/api/info?shortUrl=B'
```
```
> GET /api/info?shortUrl=8dxu
Response：
  {
      "code": "0",
      "errMsg": "OK",
      "data": {
          "url": "https://www.baidu.com",
          "created_at": "2020-02-24 12:21:02.151994 +0800 CST m=+94.081029585",
          "expiration_in_minutes": 100
      }
  }
```

> 3.短地址访问-重定向（307）

```
$ curl http://127.0.0.0.1:8080/8dxu
```
http://dwz.murphyyi.com/B
