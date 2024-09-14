# 分布式Go服务项目

## 项目概述

本项目是一个分布式Go服务系统，包含多个微服务模块，用于处理Web请求、API调用和后台数据处理。项目架构设计旨在支持高并发处理和服务间的高效通信。

## 服务架构

本项目包括以下主要服务：

- **服务注册**：管理服务实例的注册与发现。
- **日志服务**：提供日志记录功能。
- **评分服务**（Grading Service）：处理特定的数据评分逻辑。
- **Portal**：用户界面和前端服务。

### 服务交互

- 用户通过 **Portal** 发送请求。
- **Portal** 服务将请求转发到后端服务如 **评分服务**。
- **服务注册** 维护所有服务实例的注册信息，并向其他服务提供服务查找功能。
- **日志服务** 负责记录整个系统的日志信息。

![服务架构图](path/to/your/architecture_image.png)  <!-- 如果你有架构图的图片，替换这里的路径 -->

## 快速启动

### 环境要求

- Go 1.15+
- Docker (可选)

### 运行服务

1. 克隆项目到本地：
   ```bash
   git clone https://your-repository-url.git
   cd your-project-directory
2. 启动服务注册：
   ```bash
   go run cmd/registryservice/main.go
   ```
3. 启动日志服务：
   ```bash
   go run cmd/logservice/main.go
   ```
4. 启动评分服务：
   ```bash
   go run cmd/gradingservice/main.go
   ```
5. 启动Portal：
   ```bash
   go run cmd/portal/main.go
   ```
## 服务监控
```
2024/09/14 10:33:12 Heartbeat check passed for Portal
2024/09/14 10:33:12 Heartbeat check passed for GradingService
...
```
