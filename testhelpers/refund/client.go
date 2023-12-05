//
//
// File generated from our OpenAPI spec
//
//

// Package refund provides the /refunds APIs
package refund

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
)

// Client is used to invoke /refunds APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Expire is the method for the `POST /v1/test_helpers/refunds/{refund}/expire` API.
func Expire(id string, params *stripe.TestHelpersRefundExpireParams) (*stripe.Refund, error) {
	return getC().Expire(id, params)
}

// Expire is the method for the `POST /v1/test_helpers/refunds/{refund}/expire` API.
func (c Client) Expire(id string, params *stripe.TestHelpersRefundExpireParams) (*stripe.Refund, error) {
	path := stripe.FormatURLPath("/v1/test_helpers/refunds/%s/expire", id)
	refund := &stripe.Refund{}
	var err error
	sr := stripe.StripeRequest{Method: http.MethodPost, Path: path, Key: c.Key}
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, refund)
	return refund, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
