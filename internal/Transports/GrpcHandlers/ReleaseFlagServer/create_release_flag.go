package release_flag_server

import (
	"context"
	"github.com/bmsandoval/flagr/internal/Endpoints/ExampleEndpoints"
	"github.com/bmsandoval/flagr/pbs"
)

func (s *Server) CreateReleaseFlag(ctx context.Context, in *pbs.CreateReleaseFlagRequest) (*pbs.CreateReleaseFlagResponse, error) {
	endpoint := ExampleEndpoints.MakeCreateExampleEndpoint(s.appCtx, s.services)

	response, err := endpoint(ctx, in)

	if err != nil {
		return nil, err
	}

	if response == nil {
		return nil, nil
	}

	return response.(*pbs.CreateReleaseFlagResponse), nil
}
