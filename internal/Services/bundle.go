package Services

import (
	"github.com/bmsandoval/flagr/internal/Services/ReleaseFlagService"
	"github.com/bmsandoval/flagr/internal/Utilities/AppContext"
	"reflect"
)

type Bundle struct {
	ReleaseFlagSvc ReleaseFlagService.IReleaseFlagSvc
}

var bundlableServices = []IBundlable{
	ReleaseFlagService.Bundlable{},
}

type IBundlable interface {
	Create(appCtx AppContext.Context) interface{}
	Name() string
}

func NewBundle(appCtx AppContext.Context) *Bundle {
	bundle := &Bundle{}

	for _, bundlableService := range bundlableServices {
		helper := bundlableService.Create(appCtx)
		SetField(bundle, bundlableService.Name(), helper)
	}

	return bundle
}

func SetField(bundler *Bundle, field string, value interface{}) {
	v := reflect.ValueOf(bundler).Elem().FieldByName(field)
	if v.IsValid() {
		v.Set(reflect.ValueOf(value))
	}
}
