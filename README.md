# An ElasticSearch For A Goravel Extend Package

只需简单配置，即可快速接入elastic搜索。

elasticsearch: 8.x

## 快速开始
1、install
```bash
go get -u github.com/orangbus/goravel-elastic
```

2、Register service provider
```go
// config/app.go
import "github.com/orangbus/elastic"

"providers": []foundation.ServiceProvider{
    &elastic.ServiceProvider{},
},

3、 Publish Configuration
```bash
go run . artisan vendor:publish --package=github.com/orangbus/goravel-elastic
```
3、add `.env`(多个用`,`分割)
```
ELASTIC_HOST=http://127.0.0.1:9200 # http://127.0.0.1:9200,http://127.0.0.1:9201
ELASTIC_USERNAME=elastic
ELASTIC_PASSWORD=
```
4、Test
```go
import elastic "github.com/orangbus/elastic/facades"

func (r *WebElastic) Version(ctx http.Context) http.Response {
	version, err := elastic.Elastic().Version()
	if err != nil {
		return resp.Error(ctx, err.Error())
	}
	return resp.Data(ctx, version)
}
```

## 主要功能
- [x] 索引创建
- [x] 索引删除
- [x] 搜索
- [ ] 异步同步
- [ ] 命令行导入、删除

使用到的 elastic 官方包
```bash
go get -u github.com/elastic/go-elasticsearch/v8
```
