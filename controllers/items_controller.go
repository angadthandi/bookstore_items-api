package controllers

import (
	"fmt"
	"net/http"

	"github.com/angadthandi/bookstore_items-api/domain/items"
	"github.com/angadthandi/bookstore_items-api/services"
	"github.com/angadthandi/bookstore_oauth-go/oauth"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		return
	}

	item := items.Item{
		Seller: oauth.GetCallerID(r),
	}

	ret, err := services.ItemsService.Create(item)
	if err != nil {
		return
	}

	fmt.Println(ret)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
