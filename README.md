# Douyin-Simple-Demo
Douyin-Simple-Demo 是一个仿照抖音实现的简易版 APP，实现了视频流模块、用户模块、投稿模块和点赞模块。

## 项目介绍
1. 基于微服务架构，API 层使用 HTTP 框架 [Hertz](https://github.com/cloudwego/kitex)，微服务模块之间使用 RPC 框架 [Kitex](https://github.com/cloudwego/kitex) 通信
   
2. 使用 GORM 操作 MySQL 数据库

3. 使用 Redis 作为缓存，提升了接口性能

4. 使用 ETCD 进行服务发现和服务注册；

5. 使用 JWT 进行用户token的校验

7. 使用 Hertz 中间件 tracer 实现链路跟踪；

8. 使用 Kitex 中间件 klog 和 Hertz 中间件 hlog 进行日志记录

## 快速开始
1. 编辑 pkg/constants/constants.go 文件，修改相关配置

2. 启动相关服务，保证已经安装了 docker
```shell
make start
```

3. 启动 api 服务
```shell
make run_api
```

4. 启动 user 服务
```shell
make run_user
```

5. 启动 publish 服务
```shell
make run_publish
```

6. 启动 favorite 服务
```shell
make run_favorite
```

7. 启动 feed 服务
```shell
make run_feed
```

## 特别鸣谢
[字节跳动青训营](https://youthcamp.bytedance.com/)