package main

import (
	"context"
	"github.com/yuwe1/case/rpc-test/demo1/user"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

)

const (
	port = ":50051"
)

type UserService struct {
	// 实现 User 服务的业务对象
}

// UserService 实现了 User 服务接口中声明的所有方法
func (userService *UserService) UserIndex(ctx context.Context, in *user.UserIndexRequest) (*user.UserIndexResponse, error) {
	log.Printf("receive user index request: page %d page_size %d", in.Page, in.PageSize)

	return &user.UserIndexResponse{
		Err: 0,
		Msg: "success",
		Data: []*user.UserEntity{
			{Name: "big_cat", Age: 28},
			{Name: "sqrt_cat", Age: 29},
		},
	}, nil
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 创建 RPC 服务容器
	grpcServer := grpc.NewServer()

	// 为 User 服务注册业务实现 将 User 服务绑定到 RPC 服务容器上
	user.RegisterUserServer(grpcServer, &UserService{})
	// 注册反射服务 这个服务是CLI使用的 跟服务本身没有关系

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}