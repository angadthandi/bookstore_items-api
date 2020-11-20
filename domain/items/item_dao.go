package items

import (
	"encoding/json"
	"errors"

	"github.com/angadthandi/bookstore_items-api/domain/queries"

	"github.com/angadthandi/bookstore_items-api/clients/elasticsearch"
	"github.com/angadthandi/bookstore_utils-go/rest_errors"
)

const (
	indexItems = "items"
	typeItem   = "_doc"
)

func (i *Item) Save() rest_errors.RestErr {
	ret, err := elasticsearch.Client.Index(indexItems, typeItem, i)
	if err != nil {
		return rest_errors.NewInternalServerError(
			"error when trying to save item",
			errors.New("database error"),
		)
	}

	i.ID = ret.Id
	return nil
}

func (i *Item) Get() rest_errors.RestErr {
	itemID := i.ID
	ret, err := elasticsearch.Client.Get(indexItems, typeItem, i.ID)
	if err != nil {
		return rest_errors.NewInternalServerError(
			"error when trying to get item",
			errors.New("database error"),
		)
	}

	b, err := ret.Source.MarshalJSON()
	if err != nil {
		return rest_errors.NewInternalServerError(
			"error when trying to marshal database response",
			errors.New("database error"),
		)
	}

	err = json.Unmarshal(b, &i)
	if err != nil {
		return rest_errors.NewInternalServerError(
			"error when trying to unmarshal database response",
			errors.New("database error"),
		)
	}

	i.ID = itemID

	return nil
}

func (i *Item) Search(
	query queries.EsQuery,
) ([]Item, rest_errors.RestErr) {
	ret, err := elasticsearch.Client.Search(
		indexItems,
		query.Build(),
	)
	if err != nil {
		return nil, rest_errors.NewInternalServerError(
			"error when trying to search documents",
			errors.New("database error"),
		)
	}

	items := make([]Item, ret.TotalHits())
	for index, hit := range ret.Hits.Hits {
		b, _ := hit.Source.MarshalJSON()
		var item Item
		if err := json.Unmarshal(b, &item); err != nil {
			return nil, rest_errors.NewInternalServerError(
				"error when trying to parse response",
				errors.New("database error"),
			)
		}
		item.ID = hit.Id
		items[index] = item
	}

	if len(items) == 0 {
		return nil, rest_errors.NewNotFoundError(
			"no mathcing items found",
		)
	}

	return items, nil
}
