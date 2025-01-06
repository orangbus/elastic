package config

import (
	"github.com/goravel/framework/facades"
)

func init() {
	config := facades.Config()
	config.Add("elastic", map[string]any{
		"host":     config.Env("ELASTIC_HOST", "http://localhost:9200"),
		"username": config.Env("ELASTIC_USERNAME", "elastic"),
		"password": config.Env("ELASTIC_PASSWORD", ""),
	})
}
