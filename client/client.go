// @Title  client
// @Description  服务端
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-10-08 19:26
package client

import (
	"context"
	"fmt"
	"log"
	pb "gitee.com/shirdonl/goWebActualCombat/chapter5/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// @title    main
// @description   用于调用服务
// @auth      MGAronya（张健）             2022-10-08 19:26
// @param     void
// @return    void
func main() {
	conn, err := grpc.Dial(":8078", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial error: %v\n", err)
	}
	defer conn.Close()

	// TODO 实例化 ProgrammerService
	client := pb.NewProgrammerServiceClient(conn)

	req := new(pb.Request)
	req.Name = "shirdon"

	// TODO 调用服务
	resp, err := client.GetProgrammerInfo(context.Background(), req)
	if err != nil {
		log.Fatalf("resp error: %v\n", err)
	}

	fmt.Printf("Recevied: %v\n", resp)
}