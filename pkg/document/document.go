package document

import "github.com/elastic/go-elasticsearch/v8"

type Document struct {
	client *elasticsearch.TypedClient
}

func NewDocument(client *elasticsearch.TypedClient) *Document {
	return &Document{client: client}
}

func (i *Document) Create(indexName string) error {
	return nil
}

func (i *Document) Update(indexName string) error {
	return nil
}

func (i *Document) Delete(indexName string) error {
	return nil
}
