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

#### 总计划待做

1. 优雅地终止
2. 分布式id


## 2022-06-02 一期规划

1. 完成用户、角色、菜单基本功能的开发
2. 用户登录、并且动态菜单及权限校验
3. 完成api管理，不同的角色可访问不同的api

详细计划:

- 实现用户的增、删、改、查、分页,注意事务相关,暂时不做角色相关的东西
- 入参校验、返回结果统一、系统异常、业务异常情况处理、返回结果转换、
- swagger文档
- 跨域处理
- json统一处理
```
用户表设计
1. id 主键
2. username 用户名
3. password 密码
4. nick_name 昵称
5. mobile 电话
6. email 邮箱
7. header_img 头像
# 通用字段
1. create_time 创建时间
2. create_user 创建人
3. update_time 更新时间
4. update_user 更新人
5. enable 是否删除
```
- 角色和菜单
- 实现用户的登录、登录
- 字典管理
- 操作历史
- api管理

问题:





