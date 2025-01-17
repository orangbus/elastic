package index

import (
	"context"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
)

type Index struct {
	client *elasticsearch.TypedClient
}

func NewIndex(client *elasticsearch.TypedClient) *Index {
	return &Index{client: client}
}

func (i *Index) Create(indexName string) error {
	response, err := i.client.Indices.Create(indexName).Do(context.Background())
	if err != nil {
		return err
	}
	if !response.Acknowledged {
		return errors.New(fmt.Sprintf("%s 创建失败", indexName))
	}
	return nil
}

func (i *Index) Update(indexName string) error {
	return nil
}

func (i *Index) Delete(indexName string) error {
	response, err := i.client.Indices.Delete(indexName).Do(context.Background())
	if err != nil {
		return err
	}
	if !response.Acknowledged {
		return errors.New(fmt.Sprintf("%s 删除失败", indexName))
	}
	return nil
}

func (i *Index) Count(indexName string) (int64, error) {
	response, err := i.client.Count().Index(indexName).Do(context.Background())
	if err != nil {
		return 0, err
	}
	return response.Count, nil
}
