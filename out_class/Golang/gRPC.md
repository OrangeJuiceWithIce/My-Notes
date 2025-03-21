# gRPC
## Proto
微服务之间的"中立语言"，可以作为不同语言服务之间的桥梁

### 安装
1.Protobuf(github上下载解压，加环境变量)  
2.gRPC核心库
```
go get google.golang.org/grpc
```
3.语言代码生成器
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
### proto文件示例
```proto
syntax = "proto3";  //声明语法
option go_package = ".;service";  //.代表生成的go文件存放路径，service是生成的go文件的包名

service SayHello{
    rpc SayHello(HelloRequest) returns (HelloResponse){}  //定义一个rpc方法，方法名为SayHello，请求参数类型为HelloRequest，返回类型为HelloResponse
}

message HelloRequest{
    string requestName = 1;  //定义一个字符串类型的字段，字段名为requestName，序号为1
}

message HelloResponse{
    string responseMsg = 1;  //定义一个字符串类型的字段，字段名为responseMsg，序号为1
}
```
### 生成go文件
```proto
protoc --go_out=xxx xxx.proto  //第一参数为生成的go文件存放路径，第二参数为proto文件路径
protoc --go-grpc_out=xxx xxx.proto  //若proto文件中定义了rpc方法，可以生成文件
```

### 完善rpc方法
1.生成的xxx_grpc.pb.go文件中，包含了未完善的结构体UnimplementedUserServiceServer,以及与其绑定的未完善service
```go
type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) UserLogin(context.Context, *UserRequest) (*UserDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedUserServiceServer) UserRegister(context.Context, *UserRequest) (*UserDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}
func (UnimplementedUserServiceServer) testEmbeddedByValue()                     {}
```
2.新建一个servermain.go文件，import上述生成的go文件，新建一个struct对象,嵌入unimplementedUserServiceServer结构体，重写特定的func
```go
package main
import pb "./proto"

type Server struct {
    pb.UnimplementedUserServiceServer
}

func (s *Server) UserLogin(ctx context.Context,req *pb.UserRequest) (*pb.UserDetailResponse, error) {
	...
}
```
3.开启端口
```go
listener, err := net.listen("tcp", ":8080")  //监听端口
grpcServer := grpc.NewServer()  //创建白板Server
pb.RegisterUserServiceServer(grpcServer, &Server{})  //通过嵌入结构体的方法注册服务
err := grpcServer.Serve(listener)  //启动服务
if err != nil {
    fmt.Println("failed to serve: ", err)
}
```
4.客户端调用
```go
package main

func main(){
    conn,err:=grpc.Dial("localhost:8080",grpc.WithTransportCredentials(insecure.NewCredentials()))  //参数一：服务端口，参数二：加密和认证，这里采用空模板
    if err!= nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    client=pb.NewUserServiceClient(conn)  //创建客户端对象

    resp,_:=client.UserLogin(context.Background(),&pb.UserRequest{Username:"liuha",Password:"123456"})
    
    fmt.Println(resp.GetResponseMsg())
}
```