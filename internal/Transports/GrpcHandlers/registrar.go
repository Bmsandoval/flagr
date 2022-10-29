package GrpcHandlers

import (
	"github.com/bmsandoval/flagr/internal/Services"
	"github.com/bmsandoval/flagr/internal/Utilities/AppContext"
	"google.golang.org/grpc"
)

type RegisterableGrpcHandler func(*grpc.Server, AppContext.Context, Services.Bundle)

var RegisterableGrpcHandlers []RegisterableGrpcHandler

func RegisterGrpcHandler(registerableGrpcHandler RegisterableGrpcHandler) {
	RegisterableGrpcHandlers = append(RegisterableGrpcHandlers, registerableGrpcHandler)
}

func ConfigureGrpcHandlers(server *grpc.Server, appCtx AppContext.Context, svcBundle Services.Bundle) {
	for _, registerableHandler := range RegisterableGrpcHandlers {
		registerableHandler(server, appCtx, svcBundle)
	}
}
