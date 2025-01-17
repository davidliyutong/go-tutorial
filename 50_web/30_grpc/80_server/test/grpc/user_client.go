package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	pb "github.com/rebirthmonkey/go/50_web/30_grpc/80_server/apiserver/user/controller/grpc/v1"
)

func main() {
	const address = "127.0.0.1:8081"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect.", err)
		return
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)

	ctx := context.Background()
	c := &pb.ListUsersRequest{}
	sl, _ := client.ListUsers(ctx, c)
	fmt.Print("User List is:", sl)
}
