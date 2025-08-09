## Blog 项目

Blog Project 实现了以下 2 类功能：
- **用户管理：** 支持 用户注册、用户登录、获取用户列表、获取用户详情、更新用户信息、修改用户密码、注销用户 7 种用户操作；
- **博客管理：** 支持 创建博客、获取博客列表、获取博客详情、更新博客内容、删除博客、批量删除博客 6 种博客操作。

## Features

- 使用了简洁架构；
- 使用众多常用的 Go 包：gorm, casbin, govalidator, jwt-go, gin, cobra, viper, pflag, zap, pprof, grpc, protobuf 等；
- 规范的目录结构，使用 [project-layout](https://github.com/golang-standards/project-layout) 目录规范；
- 具备认证(JWT)和授权功能(casbin)；
- 独立设计的 log 包、error 包；
- 使用高质量的 Makefile 管理项目；
- 静态代码检查；
- 带有单元测试、性能测试、模糊测试、Mock 测试测试案例；
- 丰富的 Web 功能（调用链、优雅关停、中间件、跨域、异常恢复等）；
  - HTTP、HTTPS、gRPC 服务器实现；
  - JSON、Protobuf 数据交换格式实现；
- 项目遵循众多开发规范：代码规范、版本规范、接口规范、日志规范、错误规范、提交规范等；
- 访问 PostgreSQL 编程实现；
- 实现的业务功能：用户管理、博客管理；
- RESTful API 设计规范；
- OpenAPI 3.0/Swagger 2.0 API 文档；


## Installation

```bash
$ git clone https://github.com/marmotedu/miniblog.git
$ go work use miniblog # 如果 Go 版本 > 1.18
$ cd miniblog
$ make # 编译源码
```


## License

[MIT](https://choosealicense.com/licenses/mit/)
