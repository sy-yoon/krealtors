package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/sy-yoon/krealtors/api-gateway/handlers"
	"github.com/sy-yoon/krealtors/api-gateway/middleware"
	realtorpb "github.com/sy-yoon/krealtors/protos/g1/realtor"
	regionpb "github.com/sy-yoon/krealtors/protos/g1/region"
	userpb "github.com/sy-yoon/krealtors/protos/g1/user"
	"github.com/sy-yoon/krealtors/utils"
)

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
	LoadSettings()
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
		Settings.UserEndPoint,
		options,
	); err != nil {
		log.Fatalf("failed to register gRPC gateway: %v", err)
	}

	if err := regionpb.RegisterRegionServiceHandlerFromEndpoint(
		ctx,
		mux,
		Settings.RegionEndPoint,
		options,
	); err != nil {
		log.Fatalf("failed to register gRPC gateway: %v", err)
	}

	if err := realtorpb.RegisterRealtorServiceHandlerFromEndpoint(
		ctx,
		mux,
		Settings.RealtorEndPoint,
		options,
	); err != nil {
		log.Fatalf("failed to register gRPC gateway: %v", err)
	}

	utils.RedisInit()
	conn, err := grpc.Dial(Settings.UserEndPoint, options...)
	if err != nil {
		log.Fatal(err)
	}
	grpcUserClient := userpb.NewUserServiceClient(conn)

	muxWrapper := middleware.NewCacheHandler(mux, utils.GetRedisClient())
	// Creating a normal HTTP server
	server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.Recovery())
	server.Use(static.Serve("/assets", static.LocalFile(Settings.StaticContents, false)))
	server.Use(static.Serve("/realty", static.LocalFile(Settings.RealtyItemDirectory, false)))
	server.Any("v1/users/*w", gin.WrapH(muxWrapper))
	server.Any("v1/region/*w", gin.WrapH(mux))
	server.Any("v1/realtor/*w", gin.WrapH(muxWrapper))
	server.POST("v1/storage/images", middleware.IsAuthorized, handlers.SaveAndResizeImage)

	server.GET("/", func(c *gin.Context) {
		c.File("web/www/static/index.html")
	})

	accountHandler := handlers.AccountHandler{grpcUserClient}

	server.POST("account/login", accountHandler.SignIn)
	//	server.POST("account/register", accountHandler.SignUp)

	server.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Ok")
	})

	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:3000", "http://10.0.0.135:3000"}
	// server.Use(cors.New(config))

	// start server
	err = server.Run("0.0.0.0:" + strconv.Itoa(Settings.GetBindPort()))
	if err != nil {
		log.Fatal(err)
	}
}
