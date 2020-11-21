package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/angadthandi/bookstore_items-api/src/domain/queries"

	"github.com/angadthandi/bookstore_items-api/src/domain/items"
	"github.com/angadthandi/bookstore_items-api/src/services"
	"github.com/angadthandi/bookstore_items-api/src/utils/http_utils"
	"github.com/angadthandi/bookstore_oauth-go/oauth"
	"github.com/angadthandi/bookstore_utils-go/rest_errors"
	"github.com/gorilla/mux"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		// TODO fix error
		http_utils.RespondEror(w, err)
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
	vars := mux.Vars(r)
	itemID := strings.TrimSpace(vars["id"])

	ret, err := services.ItemsService.Get(itemID)
	if err != nil {
		http_utils.RespondEror(w, err)
		return
	}

	http_utils.RespondJson(w, http.StatusOK, ret)
}

func (c *itemsController) Search(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		apiErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondEror(w, apiErr)
		return
	}
	defer r.Body.Close()

	var query queries.EsQuery
	if err := json.Unmarshal(b, &query); err != nil {
		apiErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondEror(w, apiErr)
		return
	}

	items, searchErr := services.ItemsService.Search(query)
	if searchErr != nil {
		http_utils.RespondEror(w, searchErr)
		return
	}

	http_utils.RespondJson(w, http.StatusOK, items)
}
