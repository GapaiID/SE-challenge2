package policies

import "go.uber.org/fx"

var Module = fx.Module(
	"policies",
	fx.Provide(NewUserPolicy),
)
