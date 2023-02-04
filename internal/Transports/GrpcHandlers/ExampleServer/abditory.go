package release_flag_server

import (
	"github.com/bmsandoval/flagr/internal/Services"
	"github.com/bmsandoval/flagr/internal/Transports/GrpcHandlers"
	"github.com/bmsandoval/flagr/internal/Utilities/AppContext"
	"github.com/bmsandoval/flagr/pbs"
	"google.golang.org/grpc"
)

func init() {
	GrpcHandlers.RegisterGrpcHandler(
		func(s *grpc.Server, ctx AppContext.Context, svcBundle Services.Bundle) {
			pbs.RegisterExampleServer(s,
				&Server{
					appCtx: ctx, services: svcBundle,
				})
		})
}

type Server struct {
	appCtx   AppContext.Context
	services Services.Bundle
	pbs.UnimplementedExampleServer
}
