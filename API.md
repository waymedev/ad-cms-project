# API

<!-- TOC -->

- [API](#api)
  - [Pre](#pre)
  - [0. 权限管理](#0-权限管理)
  - [1. 用户相关](#1-用户相关)
    - [1.1. 登录](#11-登录)
    - [1.2. 用户列表](#12-用户列表)
    - [1.3. 单个用户](#13-单个用户)
    - [1.4. 修改用户](#14-修改用户)
    - [1.5. 删除用户](#15-删除用户)
  - [2. 订单相关](#2-订单相关)
  - [3. 材料相关](#3-材料相关)
  - [4. 绩效相关](#4-绩效相关)
  - [5. 财务相关](#5-财务相关)

<!-- /TOC -->

## Pre

鉴权采用 JWT TOKEN 除登录接口外其他接口 `Header` 需带 `Authorization` 头

## 0. 权限管理

- 管理员：1
- 普通用户：0

## 1. 用户相关

### 1.1. 登录

- POST /api/login
- payload:

```json
{
  "username": "admin",
  "password": "admin"
}
```

- return:

```json
{
  "code" : 0,
  "data" : {
    "user": {
      "username" : "admin",
      "type" : 1
  },
    "token" : "ABCDEFG"
  }

}
```

### 1.2. 用户列表

- GET /api/user
- payload:

header:  Authorization

- return:

```json
{
  "code" : 0,
  "data" : [
    {
      "username" : "admin1",
      "password" : "admin1",
      "type" : "1",
      "creatTime" : "2010/10/10 12:12:12",
      "updateTime": "2010/10/10 12:12:12"
    },
    {
      "username" : "admin2",
      "password" : "admin2",
      "type" : "0",
      "creatTime" : "2010/10/10 12:12:12",
      "updateTime": "2010/10/10 12:12:12"
    }
  ]
}
```

### 1.3. 单个用户

- GET /api/user/:id
- return:

```json
{
  "code" : 0,
  "data" : {
    "username" : "admin2",
    "password" : "admin2",
    "type" : "0",
    "creatTime" : "2010/10/10 12:12:12",
    "updateTime": "2010/10/10 12:12:12"
  }
}
```

### 1.4. 修改用户
- PATCH /api/user/:id
- payload:
```json
{
  "username" : "update",
  "password" : "update",
  "type" : 1
}
```

### 1.5. 删除用户
- DELETE /api/user/:id
- return:

```json
{
  "code" : 0,
  "data" : true
}
```

## 2. 订单相关

## 3. 材料相关

## 4. 绩效相关

## 5. 财务相关












