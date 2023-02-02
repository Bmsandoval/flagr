package ExampleService

// NOTE : Filename 'abditory' chosen to keep this file at the top of the list

import (
	"github.com/bmsandoval/flagr/internal/Utilities/AppContext"
)

// Bundlable adheres to Services.IBundlable and allows the service to be bundled
type Bundlable struct{}

// Create is a copy-pasta property
func (h Bundlable) Create(appCtx AppContext.Context) interface{} {
	return &ExampleSvc{
		AppCtx: appCtx,
	}
}

// Name Return the name of the service as it is defined in Services.Bundle
func (h Bundlable) Name() string {
	return "ExampleSvc"
}

// ExampleSvc adheres to IExampleSvc
type ExampleSvc struct {
	AppCtx AppContext.Context
}
type IExampleSvc interface {
	InsertExample(label string) (string, error)
}
