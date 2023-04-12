include "base.thrift"

namespace go tangerine.little_book_book

service LittleBookBooker {
    // 保存本子内容
    SaveContentResponse SaveContent(1:SaveContentRequest req)
    // 获取本子内容
    GetContentResponse GetContent(1:GetContentRequest req)
}

struct SaveContentRequest {
    1: required string Text
    255: optional base.RPCRequest Base
}

struct SaveContentResponse {
    1: optional ContentData Content
    255: required base.RPCResponse Base
}

struct ContentData {
    1: required string Text
}

struct GetContentRequest {
    255: optional base.RPCRequest Base
}

struct GetContentResponse {
    1: optional ContentData Content
    255: required base.RPCResponse Base
}