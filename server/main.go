package main

import (
	"context"
	"log"
	"net"

	pb "github.com/GodlyV/ChatApp-Go-grpc/proto"
	"google.golang.org/grpc"
)

const(
	port = ":8080"
)
type myChatService struct{
	pb.UnimplementedChatServiceServer
}
func (s myChatService) ChatInitiate(context.Context, *pb.InitiateRequest) (*pb.InitiateResponse, error){
	return &pb.InitiateResponse{
		Id: 1,
	},nil
}
func main(){
	lis,err:= net.Listen("tcp",port)
	if err != nil{
		log.Fatalf("Failed to start the server %v",err)
	}
	grpcServer := grpc.NewServer()
	service := &myChatService{}
	log.Printf("server started at %v", lis.Addr())
	pb.RegisterChatServiceServer(grpcServer, service)
	if err := grpcServer.Serve(lis); err !=nil{
		log.Fatalf("Failed to start: %v",err)
	}
}