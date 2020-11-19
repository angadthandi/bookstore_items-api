package items

import (
	"encoding/json"
	"errors"

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
