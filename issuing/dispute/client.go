//
//
// File generated from our OpenAPI spec
//
//

// Package dispute provides the /issuing/disputes APIs
package dispute

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /issuing/disputes APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new issuing dispute.
func New(params *stripe.IssuingDisputeParams) (*stripe.IssuingDispute, error) {
	return getC().New(params)
}

// New creates a new issuing dispute.
func (c Client) New(params *stripe.IssuingDisputeParams) (*stripe.IssuingDispute, error) {
	dispute := &stripe.IssuingDispute{}
	var err error
	sr := stripe.StripeRequest{Method: http.MethodPost, Path: "/v1/issuing/disputes", Key: c.Key}
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, dispute)
	return dispute, err
}

// Get returns the details of an issuing dispute.
func Get(id string, params *stripe.IssuingDisputeParams) (*stripe.IssuingDispute, error) {
	return getC().Get(id, params)
}

// Get returns the details of an issuing dispute.
func (c Client) Get(id string, params *stripe.IssuingDisputeParams) (*stripe.IssuingDispute, error) {
	path := stripe.FormatURLPath("/v1/issuing/disputes/%s", id)
	dispute := &stripe.IssuingDispute{}
	var err error
	sr := stripe.StripeRequest{Method: http.MethodGet, Path: path, Key: c.Key}
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, dispute)
	return dispute, err
}

// Update updates an issuing dispute's properties.
func Update(id string, params *stripe.IssuingDisputeParams) (*stripe.IssuingDispute, error) {
	return getC().Update(id, params)
}

// Update updates an issuing dispute's properties.
func (c Client) Update(id string, params *stripe.IssuingDisputeParams) (*stripe.IssuingDispute, error) {
	path := stripe.FormatURLPath("/v1/issuing/disputes/%s", id)
	dispute := &stripe.IssuingDispute{}
	var err error
	sr := stripe.StripeRequest{Method: http.MethodPost, Path: path, Key: c.Key}
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, dispute)
	return dispute, err
}

// Submit is the method for the `POST /v1/issuing/disputes/{dispute}/submit` API.
func Submit(id string, params *stripe.IssuingDisputeSubmitParams) (*stripe.IssuingDispute, error) {
	return getC().Submit(id, params)
}

// Submit is the method for the `POST /v1/issuing/disputes/{dispute}/submit` API.
func (c Client) Submit(id string, params *stripe.IssuingDisputeSubmitParams) (*stripe.IssuingDispute, error) {
	path := stripe.FormatURLPath("/v1/issuing/disputes/%s/submit", id)
	dispute := &stripe.IssuingDispute{}
	var err error
	sr := stripe.StripeRequest{Method: http.MethodPost, Path: path, Key: c.Key}
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, dispute)
	return dispute, err
}

// List returns a list of issuing disputes.
func List(params *stripe.IssuingDisputeListParams) *Iter {
	return getC().List(params)
}

// List returns a list of issuing disputes.
func (c Client) List(listParams *stripe.IssuingDisputeListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.IssuingDisputeList{}
			err := c.B.Call(stripe.StripeRequest{
				Method: http.MethodGet,
				Path:   "/v1/issuing/disputes",
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

// Iter is an iterator for issuing disputes.
type Iter struct {
	*stripe.Iter
}

// IssuingDispute returns the issuing dispute which the iterator is currently pointing to.
func (i *Iter) IssuingDispute() *stripe.IssuingDispute {
	return i.Current().(*stripe.IssuingDispute)
}

// IssuingDisputeList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) IssuingDisputeList() *stripe.IssuingDisputeList {
	return i.List().(*stripe.IssuingDisputeList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
