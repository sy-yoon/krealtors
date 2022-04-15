package main

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/sy-yoon/krealtors/api-gateway/handlers"
	realtorpb "github.com/sy-yoon/krealtors/protos/g1/realtor"
	regionpb "github.com/sy-yoon/krealtors/protos/g1/region"
	userpb "github.com/sy-yoon/krealtors/protos/g1/user"
)

const portNumber = "9000"
const userServerPortNumber = "9001"
const regionServerPortNumber = "9002"
const realtorServerPortNumber = "9003"

var allowedHeaders = map[string]struct{}{
	"x-request-id": {},
}

func isHeaderAllowed(s string) (string, bool) {
	// check if allowedHeaders contain the header
	if _, isAllowed := allowedHeaders[s]; isAllowed {
		// send uppercase header
		return strings.ToUpper(s), true
	}
	// if not in the allowed header, don't send the header
	return s, false
}

func main() {
	ctx := context.Background()
	mux := runtime.NewServeMux(
		runtime.WithOutgoingHeaderMatcher(isHeaderAllowed),
		runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
			header := request.Header.Get("Authorization")
			// send all the headers received from the client
			md := metadata.Pairs("auth", header)
			return md
		}),
		runtime.WithErrorHandler(func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {
			//creating a new HTTTPStatusError with a custom status, and passing error
			newError := runtime.HTTPStatusError{
				HTTPStatus: 400,
				Err:        err,
			}
			// using default handler to do the rest of heavy lifting of marshaling error and adding headers
			runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, writer, request, &newError)
		}),
	)

	options := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	if err := userpb.RegisterUserServiceHandlerFromEndpoint(
		ctx,
		mux,
		"localhost:"+userServerPortNumber,
		options,
	); err != nil {
		log.Fatalf("failed to register gRPC gateway: %v", err)
	}

	if err := regionpb.RegisterRegionServiceHandlerFromEndpoint(
		ctx,
		mux,
		"localhost:"+regionServerPortNumber,
		options,
	); err != nil {
		log.Fatalf("failed to register gRPC gateway: %v", err)
	}

	if err := realtorpb.RegisterRealtorServiceHandlerFromEndpoint(
		ctx,
		mux,
		"localhost:"+realtorServerPortNumber,
		options,
	); err != nil {
		log.Fatalf("failed to register gRPC gateway: %v", err)
	}

	// Creating a normal HTTP server
	server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.Recovery())
	server.Use(static.Serve("/assets", static.LocalFile("web/www/static", false)))
	server.Use(static.Serve("/realty", static.LocalFile("../realty", false)))
	server.Any("v1/users/*w", gin.WrapH(mux))
	server.Any("v1/region/*w", gin.WrapH(mux))
	server.Any("v1/realtor/*w", gin.WrapH(mux))
	server.POST("v1/storage/images", handlers.SaveAndResizeImage)

	server.GET("/", func(c *gin.Context) {
		c.File("web/www/static/index.html")
	})

	server.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Ok")
	})

	// start server
	err := server.Run("0.0.0.0:" + portNumber)
	if err != nil {
		log.Fatal(err)
	}
}
