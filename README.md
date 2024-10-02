# KWIC Analyzer

## 项目概述

KWIC Analyzer 是一个用于分析文本的工具，采用四种不同的软件体系结构实现，包括主程序-子程序结构、面向对象结构、事件系统结构和管道-过滤结构。用户可以通过前端页面选择不同的体系结构，输入待分析的内容，系统将返回相应的结果。

通过本工具，用户可以更好的理解四种不同的软件体系结构

## 目录结构

```
.
├── README.md                # 项目的说明文档
├── cmd
│   └── main.go              # 程序入口
├── go.mod                   # Go module 文件
├── go.sum                   # Go module 依赖的 checksum 文件
├── internal
│   ├── api
│   │   └── kwic.go          # API 相关的定义
│   ├── handlers
│   │   └── kwic.go          # 请求处理函数
│   └── logic
│       ├── kwic_0.go        # 主程序-子程序结构实现
│       ├── kwic_1.go        # 面向对象结构实现
│       ├── kwic_2.go        # 事件系统结构实现
│       └── kwic_3.go        # 管道-过滤结构实现
├── pkg
│   └── consts
│       └── consts.go        # 常量定义
└── static
    └── kwic.html            # 前端页面
```

## 使用说明

1. **安装依赖**: 确保 Go 环境已安装，并在项目目录下运行以下命令以安装依赖：
   ```bash
   go mod tidy
   ```

2. **运行程序**: 通过以下命令启动服务器：
   ```bash
   go run cmd/main.go
   ```

3. **访问页面**: 打开浏览器，访问 `http://localhost:8080/kwic`，即可进入 KWIC 分析页面。

4. **使用功能**: 选择体系结构类型，输入待分析的内容，点击“分析”按钮以获取解析结果。