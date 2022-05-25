package grpcs

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/test/bufconn"

	"github.com/sy-yoon/krealtors/gms"
	"github.com/sy-yoon/krealtors/gms-region/config"
	regionpb "github.com/sy-yoon/krealtors/protos/v1/region"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

var creds = credentials.NewTLS(&tls.Config{
	InsecureSkipVerify: true,
})

func init() {
	lis = bufconn.Listen(bufSize)
	gmService := gms.NewService()
	if err := gmService.Configure(&config.Settings); err != nil {
		log.Fatal(err)
		return
	}

	grpcServer := grpc.NewServer()
	gs := RegionServer{}
	gs.AddDBContext(gmService.GetDBContext())
	regionpb.RegisterRegionServiceServer(grpcServer, &gs)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

/*
func TestCreateCountry(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := regionpb.NewRegionServiceClient(conn)

	resp, err := client.CreateCountry(ctx, &regionpb.Country{Id: "NK", Name: ""})
	if err != nil {
		t.Fatalf("TestGetCountry failed: %v", err)
	}
	assert.Equal(t, "캐나다", resp.Name)
}

func TestGetCountry(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := regionpb.NewRegionServiceClient(conn)

	resp, err := client.GetCountry(ctx, &regionpb.GetCountryRequest{Id: "CA"})
	if err != nil {
		t.Fatalf("TestGetCountry failed: %v", err)
	}
	assert.Equal(t, "캐나다", resp.Name)
}

func TestGetProvince(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := regionpb.NewRegionServiceClient(conn)

	resp, err := client.GetProvince(ctx, &regionpb.GetProvinceRequest{Id: "CA01"})
	if err != nil {
		t.Fatalf("TestGetProvince failed: %v", err)
	}
	assert.Equal(t, "BC", resp.Name)
}
*/
func TestGetCity(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := regionpb.NewRegionServiceClient(conn)

	resp, err := client.GetCity(ctx, &regionpb.GetCityRequest{Id: "CA0101"})
	if err != nil {
		t.Fatalf("TestGetCity failed: %v", err)
	}
	fmt.Println(resp.GeoLocation)
	assert.Equal(t, "밴쿠버", resp.Name)
}

func TestListCities(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := regionpb.NewRegionServiceClient(conn)

	resp, err := client.ListCities(ctx, &regionpb.ListCitiesRequest{})
	if err != nil {
		t.Fatalf("TestListCities failed: %v", err)
	}

	fmt.Println(len(resp.Cities))
	assert.Equal(t, 21, len(resp.Cities))
}
