# YOYOJ-GO-RAG

这是一个准备使用 golang 来实现主要服务逻辑的 RAG 项目。用来思考 AI 的用法。基于 RAG 来探索 AI 应用的广泛性，思考 AI 真正落地的方案，同时也是给自己思考 AI 项目构成的一个底层思考。以便后续扩展到任意的其他 AI 框架。

## 技术选型

主要开发语言：Go

AI 辅助开发语言：Python

前端主要协助交互方案：React + TS、React Native + TS

后端开发框架预选：Fiber

主要探索的功能方向：

1、文档资料识别检索

2、网络资料加载检索

3、OCR 文档资料检索

4、自定义交互模型

5、本地模型支持

6、数据库、记忆、索引分离和协同策略。

> 同时，这个项目也是我初步学习 GO web后端开发的项目。



# 需要准备

1. 数据库 postgres
2. 向量数据库
3. 缓存 Redis



# 开始

因为 service 目录下的 main 包有所拆分，所以需要更改启动方式

```bash
go run ./cmd/service
```



# 构建

```bash
go build -o build/server ./cmd/service
```

