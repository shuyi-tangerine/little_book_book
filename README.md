

# 预计支持功能

* CRUD 本子
* 保存本子内容
* 获取本子内容
* 本子历史回溯


# 部署

* 环境变量：BOOK_DIR=/tmp


``` shell
# server
BOOK_DIR=/tmp go run *.go -server=true -P=compact -buffered=true -framed=false -addr=localhost:9090 -secure=false -web_addr=localhost:9091
# client rpc
go run *.go -server=false -P=compact -buffered=true -framed=false -addr=localhost:9090 -secure=false
# client web
curl -XGET localhost:9091/ping
```
