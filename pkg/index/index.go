package index

import (
	"context"
	"errors"
	"github.com/elastic/go-elasticsearch/v8"
)

type Index struct {
	client *elasticsearch.TypedClient
}

func NewIndex(client *elasticsearch.TypedClient) *Index {
	return &Index{client: client}
}

func (i *Index) Create(indexName string) error {
	version, _ := i.client.Info().Do(context.Background())
	return errors.New(version.Name)
}

func (i *Index) Update(indexName string) error {
	return nil
}

func (i *Index) Delete(indexName string) error {
	return nil
}
