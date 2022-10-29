package Services

import (
	"github.com/bmsandoval/flagr/internal/Utilities/AppContext"
	"reflect"
)

type Bundle struct {
}

var bundlableServices = []IBundlable{}

type IBundlable interface {
	Create(appCtx AppContext.Context) (interface{}, error)
	Name() string
}

func NewBundle(appCtx AppContext.Context) (*Bundle, error) {
	bundle := &Bundle{}

	for _, bundlableService := range bundlableServices {
		helper, err := bundlableService.Create(appCtx)
		if err != nil {
			return nil, err
		}
		SetField(bundle, bundlableService.Name(), helper)
	}

	return bundle, nil
}

func SetField(bundler *Bundle, field string, value interface{}) {
	v := reflect.ValueOf(bundler).Elem().FieldByName(field)
	if v.IsValid() {
		v.Set(reflect.ValueOf(value))
	}
}
