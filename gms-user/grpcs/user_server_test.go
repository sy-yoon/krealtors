package grpcs

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sy-yoon/krealtors/gms"
	"github.com/sy-yoon/krealtors/gms-user/config"
	userpb "github.com/sy-yoon/krealtors/protos/v1/user"
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
	gs := UserServer{}
	gs.AddDBContext(gmService.GetDBContext())
	userpb.RegisterUserServiceServer(grpcServer, &gs)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := userpb.NewUserServiceClient(conn)

	resp, err := client.CreateUser(ctx, &userpb.User{
		Name:  "Administrator",
		Email: "admin@krealtors.net",
	})
	if err != nil {
		t.Fatalf("TestGetCountry failed: %v", err)
	}

	assert.Equal(t, 1, resp.Id)
}
