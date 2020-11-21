package elasticsearch

import (
	"context"
	"fmt"
	"time"

	// "github.com/angadthandi/bookstore_items-api/logger"
	"github.com/angadthandi/bookstore_utils-go/logger"
	"github.com/olivere/elastic"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(*elastic.Client)
	Index(string, string, interface{}) (*elastic.IndexResponse, error)
	Get(string, string, string) (*elastic.GetResult, error)
	Search(string, elastic.Query) (*elastic.SearchResult, error)
}

type esClient struct {
	client *elastic.Client
}

func Init() {
	log := logger.GetLogger()

	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		// elastic.SetURL("http://0.0.0.0:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetSniff(false),
		// elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		// elastic.SetRetrier(NewCustomRetrier()),
		// elastic.SetGzip(true),
		// elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetErrorLog(log),
		// elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
		elastic.SetInfoLog(log),
		// elastic.SetHeaders(http.Header{
		// 	"X-Caller-Id": []string{"..."},
		// }),
	)

	// TODO : can add a timeout and try the above conn in a loop
	// that will give docker the chance to create network connection
	// in 2 containers & maybe we can connect this service to ES
	// within docker containers

	if err != nil {
		panic(err)
	}

	Client.setClient(client)
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

func (c *esClient) Index(
	index string,
	docType string,
	doc interface{},
) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	ret, err := c.client.Index().
		Index(index).
		Type(docType).
		BodyJson(doc).
		Do(ctx)

	if err != nil {
		logger.Error(
			fmt.Sprintf(
				"errr when trying to index document in index %s",
				index,
			),
			err,
		)
		return nil, err
	}

	return ret, nil
}

func (c *esClient) Get(
	index string,
	docType string,
	id string,
) (*elastic.GetResult, error) {
	ctx := context.Background()
	ret, err := c.client.Get().
		Index(index).
		Type(docType).
		Id(id).
		Do(ctx)

	if err != nil {
		logger.Error(
			fmt.Sprintf(
				"error when trying to get document for id: %s",
				id,
			),
			err,
		)
		return nil, err
	}

	return ret, nil
}

func (c *esClient) Search(
	index string,
	query elastic.Query,
) (*elastic.SearchResult, error) {
	ctx := context.Background()
	ret, err := c.client.
		Search(index).
		Query(query).
		RestTotalHitsAsInt(true).
		Do(ctx)

	if err != nil {
		logger.Error(
			fmt.Sprintf(
				"error when trying to search document for index: %s",
				index,
			),
			err,
		)
		return nil, err
	}

	return ret, nil
}
