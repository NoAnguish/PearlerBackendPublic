package core

import (
	"net/http"

	"github.com/NoAnguish/PearlerBackend/backend/objects/account"
	"github.com/NoAnguish/PearlerBackend/backend/objects/cocktail"
	"github.com/NoAnguish/PearlerBackend/backend/objects/pearl"
	"github.com/NoAnguish/PearlerBackend/backend/objects/subscription"
	"github.com/NoAnguish/PearlerBackend/backend/objects/view"
)

func (d *Daemon) registerAccountHandlers() {
	var handlerFunc func(http.ResponseWriter, *http.Request)

	createAccountHandler := account.CreateAccountHandler
	handlerFunc = JsonHandlerWrapper(createAccountHandler, d)
	d.RegisterPOST("/Account/Create/v1", handlerFunc)

	getAccountIdByFirebaseUIdHandler := account.GetIdByFirebaseUIdHandler
	handlerFunc = QueryParametersHandlerWrapper(getAccountIdByFirebaseUIdHandler, d)
	d.RegisterGET("/Account/GetIdByFirebaseUId/v1", handlerFunc)

	updateAccountHandler := account.UpdateAccountHandler
	handlerFunc = InsertImageHandlerWrapper(updateAccountHandler, d)
	d.RegisterPOST("/Account/Update/v1", handlerFunc)
}

func (d *Daemon) registerSubscriptionsHandlers() {
	var handlerFunc func(http.ResponseWriter, *http.Request)

	createSubscriptionHandler := subscription.CreateSubscriptionHandler
	handlerFunc = JsonHandlerWrapper(createSubscriptionHandler, d)
	d.RegisterPOST("/Subscription/Create/v1", handlerFunc)

	deleteSubscriptionHandler := subscription.DeleteSubscriptionHandler
	handlerFunc = JsonHandlerWrapper(deleteSubscriptionHandler, d)
	d.RegisterPOST("/Subscription/Delete/v1", handlerFunc)
}

func (d *Daemon) registerCocktailHandlers() {
	var handlerFunc func(http.ResponseWriter, *http.Request)

	getAllCocktailsHandler := cocktail.GetAllHandler
	handlerFunc = NoRequestHandlerWrapper(getAllCocktailsHandler, d)
	d.RegisterGET("/Cocktail/GetAll/v1", handlerFunc)

	createCocktailHandler := cocktail.CreateHandler
	handlerFunc = InsertImageHandlerWrapper(createCocktailHandler, d)
	d.RegisterPOST("/Cocktail/Create/v1", handlerFunc)

	updateCocktailHandler := cocktail.UpdateHandler
	handlerFunc = JsonHandlerWrapper(updateCocktailHandler, d)
	d.RegisterPOST("/Cocktail/Update/v1", handlerFunc)
}

func (d *Daemon) registerPearlHandlers() {
	var handlerFunc func(http.ResponseWriter, *http.Request)

	createPearlHandler := pearl.CreatePearlHandler
	handlerFunc = InsertImageHandlerWrapper(createPearlHandler, d)
	d.RegisterPOST("/Pearl/Create/v1", handlerFunc)
}

func (d *Daemon) registerViewHandlers() {
	var handlerFunc func(http.ResponseWriter, *http.Request)

	getAccountSelfViewHandlers := view.SelfAccountFullViewHandler
	handlerFunc = QueryParametersHandlerWrapper(getAccountSelfViewHandlers, d)
	d.RegisterGET("/Account/SelfFullView/v1", handlerFunc)

	getAccountGeneralViewHandlers := view.GeneralAccountFullViewHandler
	handlerFunc = QueryParametersHandlerWrapper(getAccountGeneralViewHandlers, d)
	d.RegisterGET("/Account/FullView/v1", handlerFunc)

	getCocktailViewHandler := view.CocktailFullViewHandler
	handlerFunc = QueryParametersHandlerWrapper(getCocktailViewHandler, d)
	d.RegisterGET("/Cocktail/FullView/v1", handlerFunc)

	getByNamePrefixHandler := view.GetByNamePrefixHandler
	handlerFunc = QueryParametersHandlerWrapper(getByNamePrefixHandler, d)
	d.RegisterGET("/Account/GetAllByNamePrefix/v1", handlerFunc)

	getPearlEventsHandler := view.GetPearlEventsHandler
	handlerFunc = QueryParametersHandlerWrapper(getPearlEventsHandler, d)
	d.RegisterGET("/Account/GetFeed/v1", handlerFunc)

	getSubscribersByIdHandler := view.GetAccountSubscriptionsHandler
	handlerFunc = QueryParametersHandlerWrapper(getSubscribersByIdHandler, d)
	d.RegisterGET("/Subscription/GetAccountSubscriptions/v1", handlerFunc)
}

func (d *Daemon) RegisterHandlers() {
	d.registerAccountHandlers()
	d.registerSubscriptionsHandlers()
	d.registerCocktailHandlers()
	d.registerPearlHandlers()
	d.registerViewHandlers()
}
