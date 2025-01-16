package contracts

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/info"
	"github.com/orangbus/elastic/pkg/document"
	"github.com/orangbus/elastic/pkg/index"
)

type Elastic interface {
	Client() *elasticsearch.TypedClient
	Version() (*info.Response, error)
	Search(indexName string, query map[string]interface{}, page int, limit ...int) ([]interface{}, int64, error)

	Mapping(indexName string, body map[string]interface{}) error
	IndexCreate(indexName string) error
	IndexDelete(indexName string) error

	Index() *index.Index
	Document() *document.Document
}
