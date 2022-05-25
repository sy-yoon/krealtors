package main

import (
	"log"

	"github.com/sy-yoon/krealtors/gms"
	"github.com/sy-yoon/krealtors/gms-realtor/grpcs"
	"github.com/sy-yoon/krealtors/gms-region/config"
	realtorpb "github.com/sy-yoon/krealtors/protos/v1/realtor"
	"google.golang.org/grpc"
)

// ListUsers returns all user messages

func main() {
	// interface parameter는 포인터로 전달
	gmService := gms.NewService()
	if err := gmService.Configure(&config.Settings); err != nil {
		log.Fatal(err)
		return
	}
	gmService.StartUp(func(grpcServer *grpc.Server) {
		gs := grpcs.RealtorServer{}
		gs.AddDBContext(gmService.GetDBContext())
		realtorpb.RegisterRealtorServiceServer(grpcServer, &gs)
	})
}
