package ExampleEndpoints

import (
	"context"
	"github.com/bmsandoval/flagr/internal/Services"
	"github.com/bmsandoval/flagr/internal/Utilities/AppContext"
	"github.com/bmsandoval/flagr/pbs"
	"github.com/go-kit/kit/endpoint"
	"github.com/ztrue/tracerr"
)

func MakeCreateExampleEndpoint(appCtx AppContext.Context, services Services.Bundle) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		in := request.(*pbs.CreateExampleRequest)

		id, err := services.ExampleSvc.InsertExample(in.Label)
		if err != nil {
			tracerr.Print(err)
			return &pbs.CreateExampleResponse{}, nil
		}
		_ = id

		return &pbs.CreateExampleResponse{}, nil
	}
}
