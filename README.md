# blog_grpc


http
cd http
// goctl api go -api=http.api -dir=./

proto

cd rpc/user
goctl rpc protoc user.proto --go_out=./ --go-grpc_out=./ --zrpc_out=.


model 

cd rpc
goctl model mysql ddl --src ./mysql/sql/blog_user.sql --dir ./mysql/model/user





