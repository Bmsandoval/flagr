package ReleaseFlagService

// NOTE : Filename 'abditory' chosen to keep this file at the top of the list

import (
	"github.com/bmsandoval/flagr/internal/Utilities/AppContext"
)

// Bundlable adheres to Services.IBundlable and allows the service to be bundled
type Bundlable struct{}

// Create is a copy-pasta property
func (h Bundlable) Create(appCtx AppContext.Context) interface{} {
	return &ReleaseFlagSvc{
		AppCtx: appCtx,
	}
}

// Name Return the name of the service as it is defined in Services.Bundle
func (h Bundlable) Name() string {
	return "ReleaseFlagSvc"
}

// ReleaseFlagSvc adheres to IReleaseFlagSvc
type ReleaseFlagSvc struct {
	AppCtx AppContext.Context
}
type IReleaseFlagSvc interface {
	InsertReleaseFlag(label string) (string, error)
}
