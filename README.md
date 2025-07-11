# Monit-Online-Server

这是 **Monit-Online** 的服务端项目，用于支持前端 Monit-Online 的数据同步和联网功能。

## 项目介绍

Monit-Online-Server 是基于原版 [Monit](https://github.com/fzf404/Monit) 进行的二次开发，主要用于为 [Monit-Online](https://github.com/GitPeo/Monit-Online) 前端提供后端支持。目前仅实现了 **Todo** 数据的存储与同步。

## 前端仓库地址

👉 [Monit-Online 前端仓库](https://github.com/GitPeo/Monit-Online)

## 功能特性

- ☑️ 支持 Todo 数据的在线同步
- ☑️ 提供简单的 RESTful API
- ☑️ 轻量、易部署，基于 Go 实现
- 🚧 待完善：基于 Token 的身份识别（多用户管理）

## 快速启动

### 克隆代码
```bash
git clone https://github.com/GitPeo/Monit-Online-Server.git
cd Monit-Online-Server
```
### 创建数据库
```bash
# 执行以下 SQL 脚本创建数据库：
sql/monit.sql
```

### 修改配置文件
```bash
# 请根据你的环境修改配置文件：
config/config.yml
```

### 运行服务
```bash
# 使用以下命令启动服务：
go run .
```

