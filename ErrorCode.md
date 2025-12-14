# 错误码

# 1. 命名规范

```
{DOMAIN}-{TYPE}-{NUMBER}

DOMAIN: 领域缩写
  - USER: 用户域
  - KB: 知识库域 (Knowledge Base)
  - CONV: 对话域 (Conversation)
  - NOTIF: 通知域

TYPE: 错误类型
  - E: 实体/值对象验证错误
  - B: 业务逻辑错误
  - R: 资源错误
  - A: 认证/授权错误

NUMBER: 三位数字编号，也可以用 {TYPE}-{NUMBER}
```

# 2. 错误码示意

## 2.1 公共领域
PUBLIC-B-START: 服务启动失败

## 2.2 用户领域
USER-E-NAME-001: 用户名不能为空
USER-E-NAME-002: 用户名超长。用户名的字符长度应在 255 的长度范围内
USER-E-PASSWORD-001: 用户密码不能为空
USER-E-PASSWORD-002: 用户密码格式不合法。应包含大小写字母、数字，且长度不小于 6 位
