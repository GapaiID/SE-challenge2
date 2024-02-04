package mails

import "go.uber.org/fx"

var Module = fx.Module(
	"mails",
	fx.Provide(NewAuthMail),
)
