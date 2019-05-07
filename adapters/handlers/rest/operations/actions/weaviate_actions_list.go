/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/semi-technologies/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@semi.technology
 */ // Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// WeaviateActionsListHandlerFunc turns a function with the right signature into a weaviate actions list handler
type WeaviateActionsListHandlerFunc func(WeaviateActionsListParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateActionsListHandlerFunc) Handle(params WeaviateActionsListParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// WeaviateActionsListHandler interface for that can handle valid weaviate actions list params
type WeaviateActionsListHandler interface {
	Handle(WeaviateActionsListParams, *models.Principal) middleware.Responder
}

// NewWeaviateActionsList creates a new http.Handler for the weaviate actions list operation
func NewWeaviateActionsList(ctx *middleware.Context, handler WeaviateActionsListHandler) *WeaviateActionsList {
	return &WeaviateActionsList{Context: ctx, Handler: handler}
}

/*WeaviateActionsList swagger:route GET /actions actions weaviateActionsList

Get a list of Actions.

Lists all Actions in reverse order of creation, owned by the user that belongs to the used token.

*/
type WeaviateActionsList struct {
	Context *middleware.Context
	Handler WeaviateActionsListHandler
}

func (o *WeaviateActionsList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateActionsListParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
