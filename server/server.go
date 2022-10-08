// @Title  server
// @Description  服务端
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-10-08 19:26
package server

import (
	"fmt"
	"log"
	"net"

	// 导入生成的protobuf包
	pb "gitee.com/shirdonl/goWebActualCombat/chapter5/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// ProgrammerServiceServer			定义服务结构体
type ProgrammerServiceServer struct {}

// @title    GetProgrammerInfo
// @description   实现开放调用的服务
// @auth      MGAronya（张健）             2022-10-08 19:26
// @param     ctx context.Context, req *pb.Request	接收一个Request结构体
// @return    resp *pb.Response, err error			返回一个Response结构体
func (p *ProgrammerServiceServer) GetProgrammerInfo (ctx context.Context, req *pb.Request) (resp *pb.Response, err error) {
	name := req.Name
	if name == "shirdon" {
		resp = &pb.Response{
			Uid: 6,
			Username: name,
			Job: "CTO",
			GoodAt: []string{"Go", "JAVA", "PHP", "Python"},
		}
	}
	err = nil
	return
}

// @title    main
// @description   用于启动服务
// @auth      MGAronya（张健）             2022-10-08 19:26
// @param     void
// @return    void
func main() {
	port := ":8078"
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}
	fmt.Printf("listen %s\n", port)
	s := grpc.NewServer()

	// TODO 注册服务
	pb.RegisterProGrammerServiceServer(s, &ProgrammerServiceServer{})
	s.Server(l)
}