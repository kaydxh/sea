# SeaDateService API 文档

## 服务说明

`SeaDateService` 提供与日期时间相关的服务，主要用于获取当前时间。所有接口均采用 gRPC 协议，消息体采用 proto3 语法定义。

---

## 1. 接口列表

### 1.1. Now

- **功能**：生成并返回当前时间。
- **方法签名**：
  ```
  rpc Now(NowRequest) returns (NowResponse);
  ```

#### 请求参数

| 字段名      | 类型   | 说明                       | 备注                         |
|-------------|--------|----------------------------|------------------------------|
| request_id  | string | 请求ID，用于追踪请求链路   | proto字段名: request_id，json字段名: RequestId |

#### 响应参数

| 字段名      | 类型        | 说明                       | 备注                         |
|-------------|-------------|----------------------------|------------------------------|
| request_id  | string      | 请求ID，回传请求标识       | proto字段名: request_id，json字段名: RequestId |
| date        | string      | 当前时间，格式如 `2024-06-01T12:34:56Z` | proto字段名: date，json字段名: Date |
| error       | types.Error | 错误信息，详见 error.proto | proto字段名: error，json字段名: Error，字段号1000 |

---

### 1.2. NowError

- **功能**：与 Now 类似，但用于测试错误返回场景。
- **方法签名**：
  ```
  rpc NowError(NowErrorRequest) returns (NowErrorResponse);
  ```

#### 请求参数

| 字段名      | 类型   | 说明                       | 备注                         |
|-------------|--------|----------------------------|------------------------------|
| request_id  | string | 请求ID，用于追踪请求链路   | proto字段名: request_id，json字段名: RequestId |

#### 响应参数

| 字段名      | 类型        | 说明                       | 备注                         |
|-------------|-------------|----------------------------|------------------------------|
| request_id  | string      | 请求ID，回传请求标识       | proto字段名: request_id，json字段名: RequestId |
| date        | string      | 当前时间，格式如 `2024-06-01T12:34:56Z` | proto字段名: date，json字段名: Date |
| error       | types.Error | 错误信息，详见 error.proto | proto字段名: error，json字段名: Error，字段号1000 |

---

## 2. 数据结构说明

### 2.1. NowRequest / NowErrorRequest

| 字段名      | 类型   | 说明                       |
|-------------|--------|----------------------------|
| request_id  | string | 请求ID，建议每次请求唯一    |

### 2.2. NowResponse / NowErrorResponse

| 字段名      | 类型        | 说明                       |
|-------------|-------------|----------------------------|
| request_id  | string      | 请求ID，回传请求标识       |
| date        | string      | 当前时间字符串             |
| error       | types.Error | 错误信息，结构见 error.proto |

---

## 3. 错误结构（types.Error）

该结构定义在 `api/protoapi-spec/types/error.proto`，常见字段如下（需参考实际 error.proto 文件）：

| 字段名      | 类型   | 说明                       |
|-------------|--------|----------------------------|
| code        | int32  | 错误码                     |
| message     | string | 错误描述                   |

---

## 4. 示例

### 4.1. Now 接口调用示例

#### 请求

```json
{
  "RequestId": "abc123"
}
```

#### 响应（成功）

```json
{
  "RequestId": "abc123",
  "Date": "2024-06-01T12:34:56Z",
  "Error": {
    "code": 0,
    "message": "success"
  }
}
```

#### 响应（失败）

```json
{
  "RequestId": "abc123",
  "Date": "",
  "Error": {
    "code": 1001,
    "message": "internal error"
  }
}
```

---

### 4.2. NowError 接口调用示例

#### 请求

```json
{
  "RequestId": "xyz789"
}
```

#### 响应（模拟错误）

```json
{
  "RequestId": "xyz789",
  "Date": "",
  "Error": {
    "code": 1234,
    "message": "simulated error"
  }
}
```

---

## 5. 备注

- 字段的 proto 名称与 JSON 名称不同，JSON 名称采用大驼峰（如 RequestId），proto 名称为小写下划线（如 request_id）。
- 错误结构需结合 `error.proto` 文件具体定义。
- 建议每次请求都传递唯一的 `request_id`，便于链路追踪和问题排查。

如需更详细的错误码说明或其他接口扩展，请补充相关 proto 文件内容。

