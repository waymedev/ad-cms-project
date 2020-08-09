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
    - [2.6.导出单个订单](#26导出单个订单)
    - [2.7.导出全部订单](#27导出全部订单)
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
  // 客户名称
  "customer_name": "招商银行",
  //
  "file": [
    {
      // 文件名
      "file_name":"1.jpg",
      // 材料名
      "material_name":"塑料"
    }
  ],
  // 加工部门
  "department": ["大喷","写真"],
  // 面积或尺寸，这里放两种值，要么是面积要么是尺寸 不提供可选。 前端字段可以表示未{面积/尺寸}
  "progress": "制作工艺",
  "area": 12,
  // 单价
  "price": 10,
  // 制作人ID
  "maker_id": 1,
  // 后期
  "after": "加框",
  // 截至时间
  "deadline_time": 1595142814,
  // 总价
  "amount": 100,
  "note": "备注"
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
      "system_id": 9,
      "customer_name": "招商银行",
      "file": [
        {
          // 文件名
          "file_name":"1.jpg",
          // 材料名
          "material_name":"塑料"
        }
      ],
      "department": ["大喷","写真"],
      "maker": "制作人",
      "create_time": 1596278558,
      "deadline_time": 1595142814,
      "order_status": 1,
      "area" : 12,
      "price": 10,
      "sum": 120,
      "after":"加后期",
      "progress": "制作工艺",
      "amount": 500,
      "note": "备注"
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
  "data":     {
      "system_id": 9,
      "customer_name": "招商银行",
      "file": [
        {
          // 文件名
          "file_name":"1.jpg",
          // 材料名
          "material_name":"塑料"
        }
      ],
      "department": ["大喷","写真"],
      "maker": "制作人",
      "create_time": 1596278558,
      "deadline_time": 1595142814,
      "order_status": 1,
      "area" : 12,
      "price": 10,
      "sum": 120,
      "department": ["大喷","写真"],
      "after":"加后期",
      "progress": "制作工艺",
      "amount": 500,
      "note": "备注"
    }
}
```

### 2.4.修改订单内容
- PATCH /api/order
- payload:
```json
    {
      "system_id": 9,
      "customer_name": "招商银行",
      "file": [
        {
          // 文件名
          "file_name":"1.jpg",
          // 材料名
          "material_name":"塑料"
        }
      ],
      "department": ["大喷","写真"],
      "deadline_time": 1595142814,
      "order_status": 1,
      "area" : 12,
      "price": 10,
      "after":"加后期",
      "progress": "制作工艺",
      "amount": 500,
      "note": "备注"
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

### 2.7.根据条件导出订单
- GET /api/download?start=1234&end=1234&name=招商银行

### 2.8.搜索订单
- POST /api/order/search
- payload
```json
{
  "start": 123,
  "edn":123,
  "name":"银行"
}
```

- return
```json
{
  "code": 0,
  "data": [
    {
      "system_id": 9,
      "customer_name": "招商银行",
      "file": [
        {
          // 文件名
          "file_name":"1.jpg",
          // 材料名
          "material_name":"塑料"
        }
      ],
      "department": ["大喷","写真"],
      "maker": "制作人",
      "create_time": 1596278558,
      "deadline_time": 1595142814,
      "order_status": 1,
      "area" : 12,
      "price": 10,
      "sum": 120,
      "department": ["大喷","写真"],
      "after":"加后期",
      "progress": "制作工艺",
      "amount": 500,
      "note": "备注"
    }
  ]
}
```

## 4.绩效相关

### 4.1.查询当前用户订单
- GET /api/eff/:maker_id
- return:
```json
{
  "code": 0,
  "data": [
    {
      "system_id": 9,
      "customer_name": "招商银行",
      "file_name": "zhaoshang.jpg",
      "department": "铁皮部",
      "material": [
        {
          "material_id": 1,
          "material_name": "铁皮",
          "material_num": 10
        },
        {
          "material_id": 4,
          "material_name": "钢板",
          "material_num": 11
        }
      ],
      "maker_id": 1,
      "process": [
        "打孔字"
      ],
      "create_time": 1596278558,
      "deadline_time": 1595142814,
      "order_status": 1,
      "admin_status": 1,
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
  "data": [
    {
      "system_id": 9,
      "customer_name": "招商银行",
      "file_name": "zhaoshang.jpg",
      "department": "铁皮部",
      "material": [
        {
          "material_id": 1,
          "material_name": "铁皮",
          "material_num": 10
        },
        {
          "material_id": 4,
          "material_name": "钢板",
          "material_num": 11
        }
      ],
      "maker_id": 1,
      "process": [
        "打孔字"
      ],
      "create_time": 1596278558,
      "deadline_time": 1595142814,
      "order_status": 1,
      "admin_status": 1,
      "origin_amount": 100.01,
      "discount": 0.5,
      "amount": 50.005
    }
  ]
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
  "data": [
    {
      "system_id": 9,
      "customer_name": "招商银行",
      "file_name": "zhaoshang.jpg",
      "department": "铁皮部",
      "material": [
        {
          "material_id": 1,
          "material_name": "铁皮",
          "material_num": 10
        },
        {
          "material_id": 4,
          "material_name": "钢板",
          "material_num": 11
        }
      ],
      "maker_id": 1,
      "process": [
        "打孔字"
      ],
      "create_time": 1596278558,
      "deadline_time": 1595142814,
      "order_status": 1,
      "admin_status": 1,
      "origin_amount": 100.01,
      "discount": 0.5,
      "amount": 50.005
    }
  ]
}
```

## 5.财务相关

### 5.1.财务列表
- GET /api/fund
- return:
```json
{
  "code": 0,
  "data": {
    "Funds": [
      {
        "system_id": 6,
        "name": "订单完成",
        "amount": 50.005,
        "create_time": 1596281271,
        "order_id": 9,
        "fund_type": -1
      }
    ],
    "amount": -50.005
  }
}
```

### 5.2.单个财务列表
- GET /api/fund/:id
- return:
```json
{
  "code": 0,
  "data": {
    "system_id": 6,
    "name": "订单完成",
    "amount": 50.005,
    "create_time": 1596281271,
    "order_id": 9,
    "fund_type": -1
  }
}
```

### 5.3.更新财务
- PATCH /api/fund
fund_type 1 收入
fund_type -1 支出

- payload:
```json
{
  "system_id": 2,
  "name": "材料进货",
  "amount": 123.456,
  "create_time": 1593854094,
  "order_id": 1,
  "fund_type" : -1
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
    "order_id": 1,
      "fund_type" : -1
  }
}
```

### 5.4.添加财务
- POST /api/fund
- payload:
```json
{
  "name": "房租",
  "fund_type" : -1,
  "amount": 100
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









