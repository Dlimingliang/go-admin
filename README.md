## go admin

## 技术选型

#### go相关

1. web框架 gin
2. 数据库orm gorm
3. 日志记录 zap
4. 配置文件读取 viper
5. API文档 swagger

#### 其他技术栈

1. 数据库 mysql
2. 缓存 redis

#### 接口返回

API接口的响应为JSON格式，并且会响应不同的HTTP状态码

1. HTTP状态码

- 200 成功
- 400 客户端参数错误 比如json格式错误
- 401 未认证 未提供认证数据或认证数据无效
- 403 未授权 无权限
- 409 冲突、业务异常
- 500 系统未知异常

2. API返回结果

```
{
    "code": 300,
    "msg": "message",
    "validationMsg : {"参数校验错误信息"}
    "data": {"数据"}
}
```

#### 总计划待做

1. 分布式id
2. 日志重复打印

## 2022-06-02 一期规划

- casbin
- api管理鉴权

问题:

1. singleflight是什么



