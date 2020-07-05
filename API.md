# API

<!-- TOC -->

- [API](#api)
  - [Pre](#pre)
  - [0.权限管理](#0权限管理)
  - [1.用户相关](#1用户相关)
    - [1.1. 登录](#11-登录)
    - [1.2. 用户列表](#12-用户列表)
    - [1.3. 单个用户](#13-单个用户)
    - [1.4. 修改用户](#14-修改用户)
    - [1.5. 删除用户](#15-删除用户)
    - [1.6.添加用户](#16添加用户)
  - [2.订单相关](#2订单相关)
    - [2.1.添加订单](#21添加订单)
    - [2.2.订单列表](#22订单列表)
    - [2.3.单个订单](#23单个订单)
    - [2.4.修改订单内容](#24修改订单内容)
    - [2.5.删除订单](#25删除订单)
  - [3. 材料相关](#3-材料相关)
  - [4. 绩效相关](#4-绩效相关)
  - [5. 财务相关](#5-财务相关)

<!-- /TOC -->

## Pre

鉴权采用 JWT TOKEN 除登录接口外其他接口 `Header` 需带 `Authorization` 头

## 0.权限管理

- 管理员：1
- 普通用户：0

## 1.用户相关

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
  "system_id": 1,
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

### 1.6.添加用户
- POST /api/user
- payload:
```json
{
  "username": "test1",
  "password": "test2",
  "type": 1
}
```

## 2.订单相关

### 2.1.添加订单
- POST /api/order
- payload:
  - 创建时间系统自动生成
  - 订单状态默认未完成、未审核
```json
{
  "customer_name": "招商银行",
  "file_name": "zhaoshang.jpg",
  "department": "铁皮部",
  "material_id": [1,2,3],
  "maker_id": 1,
  "process": ["铁皮字","打孔字"],
  "deadline_time":"2020-01-02",
  "origin_amount": 100.01,
  "discount": 0.5
}

```

- return:
```json
{
  "code": 0,
  "data": true
}
```

### 2.2.订单列表
- GET /api/order
- return:
```json
{
  "code": 0,
  "data": [
    {
      "system_id": 2,
      "customer_name": "招商银行",
      "file_name": "zs.jpg",
      "department": "铁皮部",
      "material_id": "[1,2,3,4,5]",
      "maker_id": 1,
      "process": "",
      "create_time": 1593839927,
      "deadline_time": 1593839927,
      "order_status": 0,
      "admin_status": 0,
      "origin_amount": 0,
      "discount": 0,
      "amount": 0
    },
    {
      "system_id": 3,
      "customer_name": "招商银行",
      "file_name": "zs.jpg",
      "department": "铁皮部",
      "material_id": "[1,2,3,4]",
      "maker_id": 1,
      "process": "[\"铁皮\",\"木头\"]",
      "create_time": 1593840644,
      "deadline_time": 1593840644,
      "order_status": 0,
      "admin_status": 0,
      "origin_amount": 0,
      "discount": 0,
      "amount": 0
    }
  ]
}
```

### 2.3.单个订单
- GET /api/order/:id
- return:
```json
{
  "code": 0,
  "data": {
    "system_id": 2,
    "customer_name": "招商银行",
    "file_name": "zs.jpg",
    "department": "铁皮部",
    "material_id": "[1,2,3,4,5]",
    "maker_id": 1,
    "process": "",
    "create_time": 1593839927,
    "deadline_time": 1593839927,
    "order_status": 0,
    "admin_status": 0,
    "origin_amount": 0,
    "discount": 0,
    "amount": 0
  }
}
```

### 2.4.修改订单内容
- PATCH /api/order
- payload:
```json
{
  "system_id": 1,
  "customer_name": "修改后",
  "file_name": "修改",
  "department": "modify",
  "material_id": [2,3,4],
  "maker_id": 1,
  "process": ["打孔字"],
  "deadline_time": "2020/10/10",
  "order_status": 1,
  "admin_status": 1
}
```

### 2.5.删除订单
- DELETE /api/order/:id
- return:
```json
{
  "code": 0,
  "data": true
}
```

## 3. 材料相关

## 4. 绩效相关

## 5. 财务相关












