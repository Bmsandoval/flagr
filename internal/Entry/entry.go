package Entry

import (
	"fmt"
	"github.com/bmsandoval/flagr/configs"
	"github.com/bmsandoval/flagr/internal/Db"
	"github.com/bmsandoval/flagr/internal/Services"
	"github.com/bmsandoval/flagr/internal/Transports/GrpcHandlers"
	"github.com/bmsandoval/flagr/internal/Utilities/AppContext"
	"github.com/chalvern/sugar"
	"google.golang.org/grpc"
	"log"
	"net"
	//gocache "github.com/pmylund/go-cache"
)

func Entry() {
	// Build Context and add logger
	ctx := AppContext.Context{}
	ctx.Logger = sugar.Logger{}

	// Get Configs
	config, err := configs.Configure()
	if err != nil {
		panic(err)
	}
	ctx.Config = *config

	// Setup Database
	connection, err := Db.Start(ctx)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := Db.Stop(); err != nil {
			panic(err)
		}
	}()
	ctx.DB = *connection

	// Bundle Services
	serviceBundle, err := Services.NewBundle(ctx)
	if err != nil {
		panic(err)
	}

	// Bundle Servers
	grpcS := grpc.NewServer()
	GrpcHandlers.ConfigureGrpcHandlers(grpcS, ctx, *serviceBundle)

	// Start Server
	log.Println("Starting Server...")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.SrvPort))
	if err != nil {
		panic(err)
	}
	if err := grpcS.Serve(lis); err != nil {
		panic(err)
	}
}
