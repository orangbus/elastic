package elastic

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/facades"
	"github.com/spf13/cast"
	"net/http"
	"strings"
	"time"
)

const Binding = "elastic"

var App foundation.Application

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register(app foundation.Application) {
	App = app
	app.MakeConfig()

	app.Bind(Binding, func(app foundation.Application) (any, error) {
		hosts := []string{}
		host := cast.ToString(facades.Config().Get("elastic.hosts"))
		if strings.Contains(host, ",") {
			split := strings.Split(host, ",")
			for _, h := range split {
				hosts = append(hosts, h)
			}
		} else {
			hosts = append(hosts, host)
		}

		c := elasticsearch.Config{
			Addresses: hosts,
			Username:  cast.ToString(facades.Config().GetString("elastic.name")),
			Password:  cast.ToString(facades.Config().GetString("elastic.password")),
			Transport: &http.Transport{
				MaxIdleConns:        100,              // 最大空闲连接数
				MaxIdleConnsPerHost: 2,                // 每个主机的最大空闲连接数
				IdleConnTimeout:     time.Second * 10, // 空闲连接超时时间
			},
		}
		return elasticsearch.NewTypedClient(c)
	})
}

func (receiver *ServiceProvider) Boot(app foundation.Application) {

}
