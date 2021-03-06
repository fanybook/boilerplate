package app

import (
	"github.com/lara-go/boilerplate/app/conf"
	"github.com/lara-go/boilerplate/app/providers"
	"github.com/lara-go/larago"
	"github.com/lara-go/larago/cache"
	"github.com/lara-go/larago/cli"
	"github.com/lara-go/larago/database"
	"github.com/lara-go/larago/events"
	"github.com/lara-go/larago/foundation"
	"github.com/lara-go/larago/http"
	"github.com/lara-go/larago/validation"
)

// Bootstrap application.
func Bootstrap(name, version, description string) *larago.Application {
	// Make application instance.
	application := foundation.MakeApplication(name, version, description)

	// Register signals handler.
	application.Bind(&cli.SignalsHandler{}, (*larago.SignalsHandler)(nil))

	// Register exit handler.
	application.Bind(&foundation.ExitHandler{}, (*larago.ExitHandler)(nil))

	// Register application kernel.
	application.Bind(cli.NewKernel(), (*larago.Kernel)(nil))

	// Import application config.
	application.SetConfig(func() larago.Config {
		return conf.NewConfig()
	})

	// Register application services.
	application.Register(
		// Common service providers.
		&cache.ServiceProvider{},
		&database.ServiceProvider{},
		&foundation.ServiceProvider{},
		&http.ServiceProvider{},
		&validation.ServiceProvider{},

		// Default application service provider.
		&providers.ApplicationServiceProvider{},
	)

	// Register facades.
	application.Facade(
		events.FacadeWrapper,
		http.FacadeWrapper,
	)

	return application
}
