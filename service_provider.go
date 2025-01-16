package elastic

import (
	"github.com/goravel/framework/contracts/foundation"
)

const Binding = "elastic"

var App foundation.Application

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register(app foundation.Application) {
	App = app
	app.MakeConfig()

	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		return NewClient()
	})
}

func (receiver *ServiceProvider) Boot(app foundation.Application) {
	app.Publishes("github.com/orangbus/goravel-elastic", map[string]string{
		"config/elastic.go": app.ConfigPath("elastic.go"),
	})
}
