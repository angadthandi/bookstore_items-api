package services

import (
	"net/http"

	"github.com/angadthandi/bookstore_items-api/domain/items"
	"github.com/angadthandi/bookstore_utils-go/rest_errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, *rest_errors.RestErr)
	Get(string) (*items.Item, *rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(
	item items.Item,
) (*items.Item, *rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("", http.StatusNotImplemented, "", nil)
}

func (s *itemsService) Get(
	id string,
) (*items.Item, *rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("", http.StatusNotImplemented, "", nil)
}