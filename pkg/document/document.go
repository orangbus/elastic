package document

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/elastic/go-elasticsearch/v8"
)

type Document struct {
	client *elasticsearch.TypedClient
}

func NewDocument(client *elasticsearch.TypedClient) *Document {
	return &Document{client: client}
}

func (e *Document) Create(indexName, id string, doc any) error {
	_, err := e.client.Index(indexName).Id(id).Document(doc).Do(context.Background())
	return err
}
func (e *Document) First(indexName string, id int) ([]byte, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"term": map[string]interface{}{
							"id": id,
						},
					},
				},
			},
		},
	}
	marshal, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(marshal)
	res, err2 := e.client.Search().Index(indexName).Raw(reader).Do(context.Background())
	if err2 != nil {
		return nil, err2
	}
	for _, item := range res.Hits.Hits {
		return item.Source_, nil
	}
	return nil, errors.New("not found")
}

func (e *Document) Update(indexName, id string, doc any) error {
	_, err := e.client.Index(indexName).Id(id).Document(doc).Do(context.Background())
	return err
}
func (e *Document) Bulk(indexName string, data []byte) error {
	reader := bytes.NewReader(data)
	res, err := e.client.Bulk().Index(indexName).Raw(reader).Do(context.Background())
	if err != nil {
		return err
	}
	if res.Errors {
		return errors.New("批量插入失败")
	}
	return nil
}

func (e *Document) Delete(indexName, id string) error {
	_, err := e.client.Delete(indexName, id).Do(context.Background())
	return err
}
