package AppContext

import (
	"context"
	"github.com/bmsandoval/flagr/configs"
	"github.com/bmsandoval/flagr/pkg/DbContext"
	"github.com/chalvern/sugar"
)

type Context struct {
	DB        DbContext.Connection
	Config    configs.Env
	GoContext context.Context
	Logger    sugar.Logger
}
