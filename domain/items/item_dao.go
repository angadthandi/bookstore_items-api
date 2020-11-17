package items

import (
	"errors"

	"github.com/angadthandi/bookstore_items-api/clients/elasticsearch"
	"github.com/angadthandi/bookstore_utils-go/rest_errors"
)

const (
	indexItems = "items"
)

func (i *Item) Save() *rest_errors.RestErr {
	ret, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return rest_errors.NewInternalServerError(
			"error when trying to save item",
			errors.New("database error"),
		)
	}

	i.ID = ret.Id
	return nil
}
