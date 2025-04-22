package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/orangbus/elastic/bootstrap"
	"github.com/orangbus/elastic/facades"
	"github.com/spf13/cast"
	"testing"
)

var (
	indexName = "movie"
)

func init() {
	bootstrap.Boot()
}

// 模拟user数据
type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func generateUser(id int) User {
	return User{
		Id:    id,
		Name:  faker.Username(),
		Phone: faker.Phonenumber(),
	}
}
func generateUserList(total int, startId ...int) []User {
	var list []User
	start := 1
	if len(startId) > 0 {
		start = startId[0]
	}
	for i := 0; i < total; i++ {
		list = append(list, generateUser(start+i))
	}
	return list
}

func TestVersion(t *testing.T) {
	version, err := facades.Elastic().Version()
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Logf("%v", version)
}

// 创建索引
func TestIndexCreate(t *testing.T) {
	err := facades.Elastic().Index().Create(indexName)
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Logf("%s created", indexName)
}
func TestIndexDelete(t *testing.T) {
	err := facades.Elastic().Index().Delete(indexName)
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Logf("%s deleted", indexName)
}

func TestCreate(t *testing.T) {
	user := generateUser(1)
	err := facades.Elastic().Document().Create(indexName, cast.ToString(user.Id), user)
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Logf("user:%d created", user.Id)
}
func TestBulk(t *testing.T) {
	list := generateUserList(100)
	var buf bytes.Buffer
	for _, u := range list {
		meta := []byte(fmt.Sprintf(`{"index":{ "_index" : "%s", "_id" : "%d" }}%s`, indexName, u.Id, "\n"))
		str, err := json.Marshal(u)
		if err != nil {
			continue
		}
		str = append(str, "\n"...)
		buf.Write(meta)
		buf.Write(str)
	}

	err := facades.Elastic().Document().Bulk(indexName, buf.Bytes())
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Logf("done")
}

func TestSearch(t *testing.T) {
	//query := map[string]interface{}{
	//	"query": map[string]interface{}{
	//		"match": map[string]interface{}{
	//			"name": "斗罗",
	//		},
	//	},
	//}
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					//{
					//	"match": map[string]interface{}{
					//		"vod_name": keyword,
					//	},
					//},
					{
						"term": map[string]interface{}{
							"api_id": 82,
						},
					},
					//{
					//	"term": map[string]interface{}{
					//		"type_id": typeId,
					//	},
					//},
				},
			},
		},
	}
	list, total, err := facades.Elastic().Search(indexName, query, 1, 10)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("total:%d", total)
	t.Log(string(list))
}

func TestDelete(t *testing.T) {
	id := "1"
	err := facades.Elastic().Document().Delete(indexName, id)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("document id:%s deleted", id)
}

func TestFirst(t *testing.T) {
	id := 1000
	data, err := facades.Elastic().First(indexName, id)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(data))
}
