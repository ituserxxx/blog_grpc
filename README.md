# blog_grpc

genate http
./goctl http new http

http 
cd http
../goctl api go -api=http.api -dir=./



### genate mysql model 

cd rpc
../goctl model mysql  ddl --src ./mysql/sql/blog_user.sql --dir ./mysql/model/user
../goctl model mysql  ddl --src ./mysql/sql/blog_article.sql --dir ./mysql/model/article

### genate  rpc code

cd rpc/
../goctl rpc new user

### genate  proto code

cd rpc/user
../../goctl rpc protoc user.proto --go_out=./ --go-grpc_out=./ --zrpc_out=.

cd rpc/article
../../goctl rpc protoc article.proto --go_out=./ --go-grpc_out=./ --zrpc_out=.

### genate dockerfile

cd http
goctl docker --go http.go --exe httpapi
docker build -t httpapi:v1 .



### start rpc service

go run rpc/user/user.go -f=rpc/user/etc/user.yaml
go run rpc/tag/tag.go -f=rpc/tag/etc/tag.yaml
go run http/http.go -f=http/etc/http-api.yaml
