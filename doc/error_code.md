# SeaDateService 错误码文档

本文件根据 `sea/api/protoapi-spec/sea-date/v1/error.proto` 自动生成，详细列举了 SeaDateService 相关接口可能返回的错误码及其含义。

---

## 错误码列表

| 错误码（枚举名）      | 数值 | 说明 |
|----------------------|------|------|
| OK                   | 0    | 无错误 |
| CANCELLED            | 1    | 请求被客户端取消 |
| UNKNOWN              | 2    | 出现未知的服务器错误。通常是服务器错误。|
| INVALID_ARGUMENT     | 3    | 客户端指定了无效参数。如需了解详情，请查看错误消息和错误详细信息 |
| DEADLINE_EXCEEDED    | 4    | 超出请求时限。仅当调用者设置的时限比方法的默认时限短（即请求的时限不足以让服务器处理请求）并且请求未在时限范围内完成时，才会发生这种情况 |
| NOT_FOUND            | 5    | 未找到指定的资源 |
| ALREADY_EXISTS       | 6    | 客户端尝试创建的资源已存在 |
| PERMISSION_DENIED    | 7    | 客户端权限不足。这可能是因为 OAuth 令牌没有正确的范围、客户端没有权限或者 API 尚未启用 |
| RESOURCE_EXHAUSTED   | 8    | 资源配额不足或达到速率限制。如需了解详情，请查看错误消息和错误详细信息 |
| FAILED_PRECONDITION  | 9    | 请求无法在当前系统状态下执行，例如删除非空目录 |
| ABORTED              | 10   | 并发冲突，例如读取/修改/写入冲突 |
| OUT_OF_RANGE         | 11   | 客户端指定了无效范围 |
| UNIMPLEMENTED        | 12   | API 方法未通过服务器实现 |
| INTERNAL             | 13   | 出现内部服务器错误。通常是服务器错误 |
| UNAVAILABLE          | 14   | 服务不可用。通常是服务器已关闭 |
| DATA_LOSS            | 15   | 出现不可恢复的数据丢失或数据损坏。客户端应该向用户报告错误 |
| UNAUTHENTICATED      | 16   | 由于 OAuth 令牌丢失、无效或过期，请求未通过身份验证 |
| InvalidParameter     | 1000 | 业务自定义：参数无效 |

---

## 备注

- 错误码与 gRPC 标准错误码保持一致，部分为业务自定义扩展。
- 建议客户端根据错误码进行相应的错误处理和提示。
- 业务自定义错误码（如 InvalidParameter=1000）可根据实际业务扩展。
- 错误码的详细使用场景请结合接口文档和实际返回内容理解。 