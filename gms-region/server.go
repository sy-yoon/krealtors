package main

import (
	"log"

	"github.com/sy-yoon/krealtors/gms"
	"github.com/sy-yoon/krealtors/gms-region/grpcs"
	regionpb "github.com/sy-yoon/krealtors/protos/v1/region"
	"google.golang.org/grpc"
)

// ListUsers returns all user messages
func main() {
	// interface parameter는 포인터로 전달
	gmService := gms.NewService()
	if err := gmService.Configure(&Settings); err != nil {
		log.Fatal(err)
		return
	}
	gmService.StartUp(func(grpcServer *grpc.Server) {
		gs := grpcs.RegionServer{}
		gs.AddDBContext(gmService.GetDBContext())
		regionpb.RegisterRegionServiceServer(grpcServer, &gs)
	})
}
