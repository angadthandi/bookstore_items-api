package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/angadthandi/bookstore_items-api/domain/items"
	"github.com/angadthandi/bookstore_items-api/services"
	"github.com/angadthandi/bookstore_items-api/utils/http_utils"
	"github.com/angadthandi/bookstore_oauth-go/oauth"
	"github.com/angadthandi/bookstore_utils-go/rest_errors"
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
		// TODO fix error
		http_utils.RespondEror(w, nil)
		return
	}

	sellerID := oauth.GetCallerID(r)
	if sellerID == 0 {
		respErr := rest_errors.NewUnauthorizedError(
			"Cannot get Caller ID from Access Token",
		)
		http_utils.RespondEror(w, respErr)
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondEror(w, respErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(reqBody, &itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid item json body")
		http_utils.RespondEror(w, respErr)
		return
	}

	itemRequest.Seller = sellerID

	ret, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		http_utils.RespondEror(w, createErr)
		return
	}

	http_utils.RespondJson(w, http.StatusCreated, ret)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
