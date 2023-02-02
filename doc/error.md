# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/protoapi-spec/sea-date/v1/error.proto](#api_protoapi-spec_sea-date_v1_error-proto)
    - [SeaDateReasonEnum](#sea-api-seadate-SeaDateReasonEnum)
  
    - [SeaDateReasonEnum.SeaDateReason](#sea-api-seadate-SeaDateReasonEnum-SeaDateReason)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_protoapi-spec_sea-date_v1_error-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/protoapi-spec/sea-date/v1/error.proto



<a name="sea-api-seadate-SeaDateReasonEnum"></a>

### SeaDateReasonEnum






 


<a name="sea-api-seadate-SeaDateReasonEnum-SeaDateReason"></a>

### SeaDateReasonEnum.SeaDateReason
Enum describing possible face fusion reasons.
Enum 描述了人脸融合业务出错的可能原因 -- 映射为具体业务错误码
大驼峰式命名法
一二级错误码之间用&#39;_&#39;隔开，服务会自动将之转换为&#39;.&#39;

| Name | Number | Description |
| ---- | ------ | ----------- |
| OK | 0 | The canonical error codes for gRPC APIs. https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto

无错误 |
| CANCELLED | 1 | 请求被客户端取消 |
| UNKNOWN | 2 | 出现未知的服务器错误。通常是服务器错误。 |
| INVALID_ARGUMENT | 3 | 客户端指定了无效参数。如需了解详情，请查看错误消息和错误详细信息 |
| DEADLINE_EXCEEDED | 4 | 超出请求时限。仅当调用者设置的时限比方法的默认时限短（即请求的时限不足以让服务器处理请求）并且请求未在时限范围内完成时，才会发生这种情况 |
| NOT_FOUND | 5 | 未找到指定的资源 |
| ALREADY_EXISTS | 6 | 客户端尝试创建的资源已存在 |
| PERMISSION_DENIED | 7 | 客户端权限不足。这可能是因为 OAuth |
| UNAUTHENTICATED | 16 | 令牌没有正确的范围、客户端没有权限或者 API 尚未启用

由于 OAuth 令牌丢失、无效或过期，请求未通过身份验证 |
| RESOURCE_EXHAUSTED | 8 | 资源配额不足或达到速率限制。如需了解详情，请查看错误消息和错误详细信息 |
| FAILED_PRECONDITION | 9 | 请求无法在当前系统状态下执行，例如删除非空目录 |
| ABORTED | 10 | 并发冲突，例如读取/修改/写入冲突 |
| OUT_OF_RANGE | 11 | 客户端指定了无效范围 |
| UNIMPLEMENTED | 12 | API 方法未通过服务器实现 |
| INTERNAL | 13 | 出现内部服务器错误。通常是服务器错误 |
| UNAVAILABLE | 14 | 服务不可用。通常是服务器已关闭 |
| DATA_LOSS | 15 | 出现不可恢复的数据丢失或数据损坏。客户端应该向用户报告错误 |
| InvalidParameter | 1000 |  |


 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

