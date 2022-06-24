package grpcs

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sy-yoon/krealtors/gms"
	"github.com/sy-yoon/krealtors/gms-realtor/config"
	realtorpb "github.com/sy-yoon/krealtors/protos/v1/realtor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	gmService := gms.NewService()
	if err := gmService.Configure(&config.Settings); err != nil {
		log.Fatal(err)
		return
	}

	grpcServer := grpc.NewServer()
	gs := RealtorServer{}
	gs.AddDBContext(gmService.GetDBContext())
	realtorpb.RegisterRealtorServiceServer(grpcServer, &gs)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestCreateReItem(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := realtorpb.NewRealtorServiceClient(conn)

	resp, err := client.CreateReItem(ctx, &realtorpb.ReItem{
		Title:  "Test",
		ReType: 1,
		TxType: 1,
		CityId: "CA0101",
		UserId: "123456",
	})
	if err != nil {
		t.Fatalf("TestGetCountry failed: %v", err)
	}

	assert.Equal(t, 1, resp.Id)
}

func TestGetReItem(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := realtorpb.NewRealtorServiceClient(conn)

	resp, err := client.GetReItem(ctx, &realtorpb.GetReItemRequest{
		Id: 1,
	})
	if err != nil {
		t.Fatalf("TestGetCountry failed: %v", err)
	}

	fmt.Println(resp)

	assert.Equal(t, int64(1), resp.Id)
}

func TestListReItemHeader(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := realtorpb.NewRealtorServiceClient(conn)

	resp, err := client.ListReItemHeader(ctx, &realtorpb.ListReItemHeaderRequest{
		PrevId: 1,
	})
	if err != nil {
		t.Fatalf("TestGetCountry failed: %v", err)
	}

	assert.Equal(t, true, len(resp.Items) > 0)
}
