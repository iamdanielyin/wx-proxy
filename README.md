# wx-proxy

由于微信访问access_token有白名单机制，封装了常用的两个API：获取access_token和发送模板消息。

## 容器部署
```shell

```

## 启动应用

注入环境变量（必填）：
```shell
APP_ID=xxx
APP_SECRET=xxx
```

## 调用接口

### 查询`access_token`
> 内部封装为stable_token
```shell
curl --location --request GET 'http://localhost:8888/wx_proxy/access_token'
```
返回正常：
```json
{
  "code": 0,
  "data": {
    "token": "69_QPCGpw4YWgKVjDq3WV-FFLdFgTd5A4E4amKXw46hAWJrr8KhCR3di3wz1R4kc13ISYcj4LT1TzZPh_3bL0rKFupxVGq9F0OLjeISXp87z_p-bp_ypAV-CvRxpzgDBNhAGANLX"
  }
}
```
返回错误：
```json
{
  "code": -1,
  "msg": "查询失败"
}
```

### 发送模板消息
```shell
curl --location --request GET 'http://localhost:8888/wx_proxy/access_token'
```
请求体参数：
```json
{
  "touser": "OPENID",
  "template_id": "ngqIpbwh8bUfcSsECmogfXcV14J0tQlEpBO27izEYtY",
  "url": "http://weixin.qq.com/download",
  "miniprogram": {
    "appid": "xiaochengxuappid12345",
    "pagepath": "index?foo=bar"
  },
  "client_msg_id": "MSG_000001",
  "simple_data": {
    "keyword1": "巧克力",
    "keyword2": "39.8元",
    "keyword3": "2014年9月22日"
  }
}
```
返回正常：
```json
{
  "code": 0,
  "msg": "ok"
}
```
返回错误：
```json
{
  "code": -1,
  "msg": "调用失败"
}
```

## 高级功能（可选）

### 修改API前缀（默认为`/wx_proxy`）
启动时指定环境变量：
```shell
API_PREFIX=xxx
```

### 调用权限控制
启动时指定环境变量：
```shell
APP_SK=123
```
调用时通过请求头字段`SK`传入相同值可调用：
```shell
curl --location --request GET 'http://localhost:8888/wx_proxy/access_token' \
--header 'SK: 123'
```
正常返回：
```json
{
  "code": 0,
  "data": {
    "token": "69_QPCGpw4YWgKVjDq3WV-FFLdFgTd5A4E4amKXw46hAWJrr8KhCR3di3wz1R4kc13ISYcj4LT1TzZPh_3bL0rKFupxVGq9F0OLjeISXp87z_p-bp_ypAV-CvRxpzgDBNhAGANLX"
  }
}
```

错误返回：
```json
{
    "code": -1,
    "msg": "403 Forbidden"
}
```