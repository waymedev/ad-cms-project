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
  - [3.材料相关](#3材料相关)
    - [3.1.获取所有材料列表](#31获取所有材料列表)
    - [3.2.获取单个材料](#32获取单个材料)
    - [3.3.修改材料](#33修改材料)
    - [3.4.删除材料](#34删除材料)
    - [3.5.新增材料](#35新增材料)
  - [4.绩效相关](#4绩效相关)
    - [4.1.查询当前用户订单](#41查询当前用户订单)
    - [4.2.订单审核](#42订单审核)
    - [4.3. 订单完成](#43-订单完成)
  - [5.财务相关](#5财务相关)
    - [5.1.财务列表](#51财务列表)
    - [5.2.单个财务列表](#52单个财务列表)
    - [5.3.更新财务](#53更新财务)
    - [5.4.添加财务](#54添加财务)
    - [5.5.删除财务](#55删除财务)

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
  "deadline_time": 1595142814,
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
  "deadline_time": 1595142814,
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


### 2.6.导出单个订单
- GET /api/download/:id
- :id 为订单ID

### 2.7.导出全部订单
- GET /api/download

## 3.材料相关

### 3.1.获取所有材料列表
- GET /api/m
- return:
```json
{
  "code": 0,
  "data": [
    {
      "system_id": 1,
      "name": "铁皮",
      "count": 20,
      "total": 0
    }
  ]
}
```

### 3.2.获取单个材料
- GET /api/m/:id
- return:
```json
{
  "code": 0,
  "data": {
    "system_id": 1,
    "name": "铁皮",
    "count": 20,
    "total": 0
  }
}
```

### 3.3.修改材料
- PATCH /api/m
- payload:
```json
{
  "system_id": 1,
  "name": "铁皮",
  "count": 100,
  "total": 0
}
```

- return:
```json
{
  "code": 0,
  "data": {
    "system_id": 1,
    "name": "铁皮",
    "count": 100,
    "total": 0
  }
}
```

### 3.4.删除材料
- DELETE /api/m/:id
- return:
```json
{
  "code": 0,
  "data": true
}
```

### 3.5.新增材料
- POST /api/m
- payload:
```json
{
  "name": "铁皮2",
  "count": 100,
  "total": 0
}
```

- return:
```json
{
  "code": 0,
  "data": true
}
```

## 4.绩效相关

### 4.1.查询当前用户订单
- GET /api/fund/:maker_id
- return:
```json
{
  "code": 0,
  "data": [
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
    },
    {
      "system_id": 4,
      "customer_name": "招商银行2",
      "file_name": "zs.jpg",
      "department": "铁皮部",
      "material_id": "[1,2,3,4]",
      "maker_id": 1,
      "process": "[\"铁皮\",\"木头\"]",
      "create_time": 1593853902,
      "deadline_time": 1593853902,
      "order_status": 0,
      "admin_status": 0,
      "origin_amount": 100,
      "discount": 0.8,
      "amount": 80
    },
    {
      "system_id": 6,
      "customer_name": "招商银行",
      "file_name": "zs.jpg",
      "department": "铁皮部",
      "material_id": "[1,2,3,4,5]",
      "maker_id": 1,
      "process": "[\"木头\"]",
      "create_time": 1593839927,
      "deadline_time": 1593839927,
      "order_status": 0,
      "admin_status": 0,
      "origin_amount": 100.01,
      "discount": 0.5,
      "amount": 50.005
    }
  ]
}
```

### 4.2.订单审核
- PATCH /api/order/admin
- payload:
  - admin_status : 0 未审核 1 已审核
```json
{
  "system_id" : 6,
  "admin_status" : 1
}
```

- return
```json
{
  "code": 0,
  "data": {
    "system_id": 6,
    "customer_name": "招商银行6",
    "file_name": "zs.jpg",
    "department": "铁皮部",
    "material_id": "[1,4]",
    "maker_id": 1,
    "process": "['木头']",
    "create_time": 1593839927,
    "deadline_time": 1593839927,
    "order_status": 0,
    "admin_status": 0,
    "origin_amount": 100.01,
    "discount": 0.5,
    "amount": 50.005
  }
}
```

### 4.3. 订单完成
- PATCH /api/order/status
- payload:
  - order_status : 0 未完成 1 已完成
```json
{
  "system_id" : 6,
  "order_status" : 1
}
```

- return:
```json
{
  "code": 0,
  "data": {
    "system_id": 6,
    "customer_name": "招商银行6",
    "file_name": "zs.jpg",
    "department": "铁皮部",
    "material_id": "[1,4]",
    "maker_id": 1,
    "process": "[\"木头\"]",
    "create_time": 1593839927,
    "deadline_time": 1593839927,
    "order_status": 1,
    "admin_status": 0,
    "origin_amount": 100.01,
    "discount": 0.5,
    "amount": 50.005
  }
}
```

## 5.财务相关

### 5.1.财务列表
- GET /api/fund
- return:
```json
{
  "code": 0,
  "data": [
    {
      "system_id": 2,
      "name": "材料进货",
      "amount": 123.123,
      "create_time": 1593854094,
      "order_id": 1
    },
    {
      "system_id": 3,
      "name": "",
      "amount": 123.123,
      "create_time": 1593922598,
      "order_id": 1
    }
  ]
}
```

### 5.2.单个财务列表
- GET /api/fund/:id
- return:
```json
{
  "code": 0,
  "data": {
    "system_id": 2,
    "name": "材料进货",
    "amount": 123.456,
    "create_time": 1593854094,
    "order_id": 1
  }
}
```

### 5.3.更新财务
- PATCH /api/fund
- payload:
```json
{
  "system_id": 2,
  "name": "材料进货",
  "amount": 123.456,
  "create_time": 1593854094,
  "order_id": 1
}
```

- return:
```json
{
  "code": 0,
  "data": {
    "system_id": 2,
    "name": "材料进货",
    "amount": 123.456,
    "create_time": 1593854094,
    "order_id": 1
  }
}
```

### 5.4.添加财务
- POST /api/fund
- payload:
```json
{
  "name": "材料进货",
  "amount": 123.456,
  "create_time": 1593854094
}
```
- return:
```json
{
  "code": 0,
  "data": true
}
```

### 5.5.删除财务
- DELETE /api/fund/:id
- return:
```json
{
  "code": 0,
  "data": true
}
```









