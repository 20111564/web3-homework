---

## 项目简介

golang练手博客demo
---

## 技术栈

### 后端

- **Go 1.21+**
- **Gin**
- **GORM**

---

## 安装与运行

### 环境要求

- Go 1.21 或更高版本
- MySQL 8.0+

### 后端启动

```bash
# 安装依赖
go mod tidy

# 运行服务
go run main.go
```

## 项目结构

```
├── README.md
├── api
│   ├── comment.go
│   ├── post.go
│   └── user.go
├── common
│   ├── e
│   │   └── globalError.go
│   └── r
│       └── r.go
├── db
│   └── mysql.go
├── go.mod
├── go.sum
├── initialize
├── main.go
├── middleware
│   └── jwt.go
├── models
│   ├── comment.go
│   ├── page.go
│   ├── post.go
│   ├── request
│   │   ├── comment.go
│   │   └── post.go
│   ├── response
│   │   ├── loginToken.go
│   │   └── page.go
│   └── user.go
├── router
│   └── router.go
├── service
│   ├── commentService.go
│   └── postService.go
└── utils
    ├── jwt.go
    └── user_holder.go

```

## API 接口文档

### POST 注册

POST /api/user/register

> Body 请求参数

```json
{
  "username": "lujie",
  "password": "123456"
}
```

请求参数

| 名称   | 位置   | 类型     | 必选 | 说明   |
|------|------|--------|----|------|
| body | body | object | 否  | none |

> 返回示例

```json
{
  "code": 0,
  "data": "string",
  "msg": "string"
}
```

### POST 登录

POST /api/user/login

> Body 请求参数

```json
{
  "username": "lujie",
  "password": "123456"
}
```

请求参数

| 名称   | 位置   | 类型     | 必选 | 说明   |
|------|------|--------|----|------|
| body | body | object | 否  | none |

> 返回示例

```json
{
  "code": 0,
  "data": {
    "token": "string"
  },
  "msg": "string"
}
```

## 文章

### POST 文章-创建

POST /api/post/create

> Body 请求参数

```json
{
  "title": "《钢铁是怎么练成的1》",
  "content": "《钢铁是怎么练成的》《钢铁是怎么练成的》《钢铁是怎么练成的》《钢铁是怎么练成的》《钢铁是怎么练成的》"
}
```

请求参数

| 名称      | 位置     | 类型     | 必选 | 说明   |
|---------|--------|--------|----|------|
| x-token | header | string | 是  | none |
| body    | body   | object | 否  | none |

> 返回示例

```json
{}
```

### POST 文章-分页列表

POST /api/post/page

> Body 请求参数

```json
{
  "page": 2,
  "pageSize": 10
}
```

请求参数

| 名称      | 位置     | 类型     | 必选 | 说明   |
|---------|--------|--------|---|------|
| body    | body   | object | 否 | none |

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "list": [
      {
        "ID": 0,
        "CreatedAt": "string",
        "UpdatedAt": "string",
        "DeletedAt": null,
        "title": "string",
        "content": "string",
        "userId": 0,
        "User": {
          。
          。
          。
        }
      }
    ],
    "total": 0,
    "page": 0,
    "pageSize": 0
  },
  "msg": "string"
}
```

### GET 文章-详情

GET /api/post/detail/2

> 返回示例

```json
{
  "code": 0,
  "data": {
    "ID": 0,
    "CreatedAt": "string",
    "UpdatedAt": "string",
    "DeletedAt": null,
    "title": "string",
    "content": "string",
    "userId": 0,
    "User": {
      "ID": 0,
      "CreatedAt": "string",
      "UpdatedAt": "string",
      "DeletedAt": null,
      "username": "string",
      "password": "string",
      "email": "string"
    }
  },
  "msg": "string"
}
```

### POST 文章-编辑

POST /api/post/edit

> Body 请求参数

```json
{
  "id": 1,
  "title": "《钢铁是怎么练成的F》",
  "content": "《钢铁是怎么练成的》《钢铁是怎么练成的》《钢铁是怎么练成的》《钢铁是怎么练成的》《钢铁是怎么练成的》"
}
```

请求参数

| 名称      | 位置     | 类型     | 必选 | 说明   |
|---------|--------|--------|----|------|
| x-token | header | string | 是  | none |
| body    | body   | object | 否  | none |

> 返回示例

```json
{
  "code": 0,
  "data": "string",
  "msg": "string"
}
```

### POST 文章-删除

POST /api/post/del

> Body 请求参数

```json
{
  "id": 12
}
```

请求参数

| 名称      | 位置     | 类型     | 必选 | 说明   |
|---------|--------|--------|----|------|
| x-token | header | string | 是  | none |
| body    | body   | object | 否  | none |

> 返回示例

```json
{
  "code": 0,
  "data": "string",
  "msg": "string"
}
```

## 评论

### POST 评论-新增

POST /api/comment/add

> Body 请求参数

```json
{
  "content": "文章写的真好",
  "postId": 1
}
```

请求参数

| 名称      | 位置     | 类型     | 必选 | 说明   |
|---------|--------|--------|----|------|
| x-token | header | string | 是  | none |
| body    | body   | object | 否  | none |

> 返回示例

```json
{
  "code": 0,
  "data": "string",
  "msg": "string"
}
```

### POST 评论-分页查询

POST /api/comment/page

> Body 请求参数

```json
{
  "page": 2,
  "pageSize": 10
}
```

请求参数

| 名称      | 位置     | 类型     | 必选 | 说明   |
|---------|--------|--------|----|------|
| x-token | header | string | 否  | none |
| body    | body   | object | 否  | none |

> 返回示例

```json
{
  "code": 0,
  "data": {
    "list": [
      {
        "ID": 0,
        "CreatedAt": "string",
        "UpdatedAt": "string",
        "DeletedAt": null,
        "content": "string",
        "userId": 0,
        "User": {
        },
        "postId": 0,
        "Post": {
        }
      }
    ],
    "total": 0,
    "page": 0,
    "pageSize": 0
  },
  "msg": "string"
}
```