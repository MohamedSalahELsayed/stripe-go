//
//
// File generated from our OpenAPI spec
//
//

// Package scheduledqueryrun provides the /sigma/scheduled_query_runs APIs
// For more details, see: https://stripe.com/docs/api#scheduled_queries
package scheduledqueryrun

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /sigma/scheduled_query_runs APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of a sigma scheduled query run.
func Get(id string, params *stripe.SigmaScheduledQueryRunParams) (*stripe.SigmaScheduledQueryRun, error) {
	return getC().Get(id, params)
}

// Get returns the details of a sigma scheduled query run.
func (c Client) Get(id string, params *stripe.SigmaScheduledQueryRunParams) (*stripe.SigmaScheduledQueryRun, error) {
	path := stripe.FormatURLPath("/v1/sigma/scheduled_query_runs/%s", id)
	scheduledqueryrun := &stripe.SigmaScheduledQueryRun{}
	var err error
	sr := stripe.StripeRequest{Method: http.MethodGet, Path: path, Key: c.Key}
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, scheduledqueryrun)
	return scheduledqueryrun, err
}

// List returns a list of sigma scheduled query runs.
func List(params *stripe.SigmaScheduledQueryRunListParams) *Iter {
	return getC().List(params)
}

// List returns a list of sigma scheduled query runs.
func (c Client) List(listParams *stripe.SigmaScheduledQueryRunListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.SigmaScheduledQueryRunList{}
			err := c.B.Call(stripe.StripeRequest{
				Method: http.MethodGet,
				Path:   "/v1/sigma/scheduled_query_runs",
				Key:    c.Key,
				Params: p,
				Body:   b,
			},
				list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for sigma scheduled query runs.
type Iter struct {
	*stripe.Iter
}

// SigmaScheduledQueryRun returns the sigma scheduled query run which the iterator is currently pointing to.
func (i *Iter) SigmaScheduledQueryRun() *stripe.SigmaScheduledQueryRun {
	return i.Current().(*stripe.SigmaScheduledQueryRun)
}

// SigmaScheduledQueryRunList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) SigmaScheduledQueryRunList() *stripe.SigmaScheduledQueryRunList {
	return i.List().(*stripe.SigmaScheduledQueryRunList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
