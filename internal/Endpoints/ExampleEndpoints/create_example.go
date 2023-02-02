package ExampleEndpoints

import (
	"context"
	"github.com/bmsandoval/flagr/internal/Services"
	"github.com/bmsandoval/flagr/internal/Utilities/AppContext"
	"github.com/bmsandoval/flagr/pbs"
	"github.com/go-kit/kit/endpoint"
)

func MakeCreateExampleEndpoint(appCtx AppContext.Context, services Services.Bundle) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		in := request.(*pbs.CreateReleaseFlagRequest)

		id, err := services.ExampleSvc.InsertExample(in.Label)
		if err != nil {
			appCtx.Logger.Error(err)
			return &pbs.CreateReleaseFlagResponse{}, nil
		}
		_ = id

		return &pbs.CreateReleaseFlagResponse{}, nil
	}
}
