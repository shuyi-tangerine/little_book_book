include "base.thrift"

namespace go tangerine.little_book_book

service LittleBookBooker {
    // 保存本子内容
    SaveContentResponse SaveContent(1:SaveContentRequest req)
    // 获取本子内容
    GetContentResponse GetContent(1:GetContentRequest req)
}

struct TimeRange {
    1: required i64 s  // start
    2: required i64 e  // end
}

struct SaveContentRequest {
    1: required string Text
    2: optional i64 ContentID = 0
    3: required i64 ContentType
    4: optional string Extra
    5: optional string Title = ''
    6: optional string Operator = ''
    255: optional base.RPCRequest Base
}

struct SaveContentResponse {
    1: optional ContentData Content
    255: required base.RPCResponse Base
}

struct ContentData {
    1: required string Text
    2: required i64 ContentID
    3: required i64 ContentType
    4: required string Title
    5: optional string Backup
    6: optional string Extra
    7: required i64 CreatedAt
    8: required string CreatedBy
    9: required i64 UpdatedAt
    10: required string UpdatedBy
}

struct GetContentRequest {
    1: required i64 ContentID
    255: optional base.RPCRequest Base
}

struct GetContentResponse {
    1: optional ContentData Content
    255: required base.RPCResponse Base
}