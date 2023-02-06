package shippinglabel

import (
	"bytes"
	"context"
	"github.com/dewaco/shippinglabel/types"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type APIContext struct {
	client     *Client
	token      *types.AuthToken
	tokenMutex sync.RWMutex
}

// NewAPIContext creates an API context
func NewAPIContext(c *Client, token *types.AuthToken) (*APIContext, error) {
	if c == nil {
		return nil, types.ErrRequiredClient
	}
	if token == nil {
		return nil, types.ErrRequiredToken
	}
	return &APIContext{client: c, token: token}, nil
}

// request creates a *request struct and sets the bearer token
func (c *APIContext) request() *request {
	return newRequest(c.client.baseURL).SetBearer(c.token.AccessToken)
}

// send checks the status of the access token and sends the request
func (c *APIContext) send(ctx context.Context, req *request) error {
	c.tokenMutex.Lock()
	if c.token != nil && c.token.IsExpired() {
		tk, err := c.client.RefreshToken(context.Background(), c.token.RefreshToken)
		if err != nil {
			c.tokenMutex.Unlock()
			return err
		}

		c.token.SetAccessToken(tk)
		req.SetBearer(c.token.AccessToken)
	}
	c.tokenMutex.Unlock()

	// Send request
	return c.client.send(ctx, req)
}

// USER

// GetUser returns the user details
// [GET]: /user
func (c *APIContext) GetUser(ctx context.Context) (resp *types.User, err error) {
	req := c.request().SetMethod(http.MethodGet).ToJSON(&resp).SetPath("/user")
	return resp, c.send(ctx, req)
}

// METADATA

// Metadata returns the carrier metadata
// [GET]: /metadata/carriers/details
func (c *APIContext) Metadata(ctx context.Context) (resp *types.CarrierMetadata, err error) {
	req := c.request().SetMethod(http.MethodGet).ToJSON(&resp).SetPath("/metadata/carriers/details")
	return resp, c.send(ctx, req)
}

// ADDRESSES

// ListAddresses returns all available user addresses
// [GET]: /addresses
func (c *APIContext) ListAddresses(ctx context.Context) (resp []*types.Address, err error) {
	req := c.request().SetMethod(http.MethodGet).ToJSON(&resp).SetPath("/addresses")
	return resp, c.send(ctx, req)
}

// CreateAddress creates a new shipment address
// [POST]: /addresses
func (c *APIContext) CreateAddress(ctx context.Context, v *types.Address) (resp *types.Address, err error) {
	req := c.request().SetMethod(http.MethodPost).SetJSON(v).ToJSON(&resp).SetPath("/addresses")
	return resp, c.send(ctx, req)
}

// GetAddress returns an address
// [GET]: /addresses/{id}
func (c *APIContext) GetAddress(ctx context.Context, id int) (resp *types.Address, err error) {
	req := c.request().SetMethod(http.MethodGet).ToJSON(&resp).SetPathf("/addresses/%d", id)
	return resp, c.send(ctx, req)
}

// UpdateAddress updates a shipment address
// [PUT]: /addresses/{id}
func (c *APIContext) UpdateAddress(ctx context.Context, v *types.Address) (err error) {
	req := c.request().SetMethod(http.MethodPut).SetJSON(v).SetPathf("/addresses/%d", v.ID)
	return c.send(ctx, req)
}

// DeleteAddress deletes a shipment address
// [DELETE]: /addresses/{id}
func (c *APIContext) DeleteAddress(ctx context.Context, id int) (err error) {
	req := c.request().SetMethod(http.MethodDelete).SetPathf("/addresses/%d", id)
	return c.send(ctx, req)
}

// PARCELS

// ListParcels returns all parcels
// [GET]: /parcels
func (c *APIContext) ListParcels(ctx context.Context) (resp []*types.Parcel, err error) {
	req := c.request().SetMethod(http.MethodGet).ToJSON(&resp).SetPath("/parcels")
	return resp, c.send(ctx, req)
}

// CreateParcel creates a parcel
// [POST]: /parcels
func (c *APIContext) CreateParcel(ctx context.Context, v *types.Parcel) (resp *types.Parcel, err error) {
	req := c.request().SetMethod(http.MethodPost).ToJSON(&resp).SetJSON(v).SetPath("/parcels")
	return resp, c.send(ctx, req)
}

// GetParcel returns a parcel
// [GET]: /parcels/{id}
func (c *APIContext) GetParcel(ctx context.Context, id int) (resp *types.Parcel, err error) {
	req := c.request().SetMethod(http.MethodGet).ToJSON(&resp).SetPathf("/parcels/%d", id)
	return resp, c.send(ctx, req)
}

// UpdateParcel updates a parcel
// [PUT]: /parcels/{id}
func (c *APIContext) UpdateParcel(ctx context.Context, v *types.Parcel) (err error) {
	req := c.request().SetMethod(http.MethodPut).SetJSON(v).SetPathf("/parcels/%d", v.ID)
	return c.send(ctx, req)
}

// DeleteParcel deletes a parcel
// [DELETE]: /parcels/{id}
func (c *APIContext) DeleteParcel(ctx context.Context, id int) (err error) {
	req := c.request().SetMethod(http.MethodDelete).SetPathf("/parcels/%d", id)
	return c.send(ctx, req)
}

// CARRIERS

// ListCarriers returns all user created carriers
// [GET]: /carriers
func (c *APIContext) ListCarriers(ctx context.Context) (resp []*types.Carrier, err error) {
	req := c.request().SetMethod(http.MethodGet).ToJSON(&resp).SetPath("/carriers")
	return resp, c.send(ctx, req)
}

// CreateCarrier creates a carrier
// [POST]: /carriers
func (c *APIContext) CreateCarrier(ctx context.Context, v *types.Carrier) (resp *types.Carrier, err error) {
	req := c.request().SetMethod(http.MethodPost).ToJSON(&resp).SetJSON(v).SetPath("/carriers")
	return resp, c.send(ctx, req)
}

// GetCarrier returns a carrier
// [GET]: /carriers/{id}
func (c *APIContext) GetCarrier(ctx context.Context, code types.CarrierCode) (resp *types.Carrier, err error) {
	req := c.request().SetMethod(http.MethodGet).ToJSON(&resp).SetPathf("/carriers/%s", code)
	return resp, c.send(ctx, req)
}

// UpdateCarrier updates a carrier
// [PUT]: /carriers/{id}
func (c *APIContext) UpdateCarrier(ctx context.Context, v *types.Carrier) (err error) {
	req := c.request().SetMethod(http.MethodPut).SetJSON(v).SetPathf("/carriers/%s", v.Code)
	return c.send(ctx, req)
}

// UpdateCarrierCredentials updates the user credentials from the carrier
// [PUT]: /carriers/{id}/cred
func (c *APIContext) UpdateCarrierCredentials(ctx context.Context, v *types.Carrier) (err error) {
	req := c.request().SetMethod(http.MethodPut).SetJSON(v).SetPathf("/carriers/%s/credentials", v.Code)
	return c.send(ctx, req)
}

// VerifyCarrier validates the user credentials
// [POST]: /carriers/{id}/verify
func (c *APIContext) VerifyCarrier(ctx context.Context, code types.CarrierCode) (err error) {
	req := c.request().SetMethod(http.MethodPost).SetPathf("/carriers/%s/verify", code)
	return c.send(ctx, req)
}

// DeleteCarrier deletes a carrier
// [DELETE]: /carriers/{id}
func (c *APIContext) DeleteCarrier(ctx context.Context, code types.CarrierCode) (err error) {
	req := c.request().SetMethod(http.MethodDelete).SetPathf("/carriers/%s", code)
	return c.send(ctx, req)
}

// CARRIER PRODUCTS

// CreateDHLProduct creates a DHL product
// [POST]: /carriers/DHL/products
func (c *APIContext) CreateDHLProduct(ctx context.Context, v *types.DHLProduct) (resp *types.DHLProduct, err error) {
	req := c.request().SetMethod(http.MethodPost).SetJSON(v).ToJSON(&resp).SetPath("/carriers/DHL/products")
	return resp, c.send(ctx, req)
}

// UpdateDHLProduct updates a DHL product
// [PUT]: /carriers/DHL/products/{id}
func (c *APIContext) UpdateDHLProduct(ctx context.Context, v *types.DHLProduct) (err error) {
	req := c.request().SetMethod(http.MethodPut).SetJSON(v).SetPathf("/carriers/DHL/products/%d", v.ID)
	return c.send(ctx, req)
}

// DeleteDHLProduct deletes a DHL product
// [DELETE]: /carriers/DHL/products/{id}
func (c *APIContext) DeleteDHLProduct(ctx context.Context, id int) (err error) {
	req := c.request().SetMethod(http.MethodDelete).SetPathf("/carriers/DHL/products/%d", id)
	return c.send(ctx, req)
}

// SHIPMENTS

// ListShipments returns all shipments
// [GET]: /shipments
// Query parameters:
// page: The page number to retrieve for the list of shipments. For example page = 0 and page_size = 10 return the
// first 10 shipments. page = 1 and page_size=10 return the next shipments (11-20). Default: 0
// page_size: The maximum number of shipments to return in the response. Must be an integer between 0 and 10000. Default: 10000
// order: Specifies the order of shipments. Available values are 'asc' or 'desc'. Default: asc
func (c *APIContext) ListShipments(ctx context.Context, page int, size int, order string) (resp []*types.Shipment, err error) {
	if page < 0 {
		page = 0
	}

	if size < 0 {
		size = 10000
	}

	if order != "asc" && order != "desc" {
		order = "asc"
	}

	req := c.request().SetMethod(http.MethodGet).ToJSON(&resp).SetPathf("/shipments?page=%d&page_size=%d&order=%s", page, size, order)
	return resp, c.send(ctx, req)
}

// ValidateShipment validates a shipment
// [POST]: /shipments/validate
func (c *APIContext) ValidateShipment(ctx context.Context, v *types.Shipment) (err error) {
	req := c.request().SetMethod(http.MethodPost).SetJSON(v).SetPath("/shipments/validate")
	return c.send(ctx, req)
}

// CreateShipment creates a shipment
// [POST]: /shipments
func (c *APIContext) CreateShipment(ctx context.Context, v *types.Shipment) (resp *types.Shipment, err error) {
	req := c.request().SetMethod(http.MethodPost).SetJSON(v).ToJSON(&resp).SetPath("/shipments")
	return resp, c.send(ctx, req)
}

// GetShipment returns a shipment
// [GET]: /shipments/{id}
func (c *APIContext) GetShipment(ctx context.Context, id int) (resp *types.Shipment, err error) {
	req := c.request().SetMethod(http.MethodGet).ToJSON(&resp).SetPathf("/shipments/%d", id)
	return resp, c.send(ctx, req)
}

// UpdateShipment updates a shipment
// [PUT]: /shipments/{id}
func (c *APIContext) UpdateShipment(ctx context.Context, v *types.Shipment) (err error) {
	req := c.request().SetMethod(http.MethodPut).SetJSON(v).SetPathf("/shipments/%d", v.ID)
	return c.send(ctx, req)
}

// DeleteShipment deletes a shipment
// [DELETE]: /shipments/{id}
func (c *APIContext) DeleteShipment(ctx context.Context, id int) (err error) {
	req := c.request().SetMethod(http.MethodDelete).SetPathf("/shipments/%d", id)
	return c.send(ctx, req)
}

// CreateShipments creates multiple shipments
// [POST]: /shipments/bulk
func (c *APIContext) CreateShipments(ctx context.Context, v []*types.Shipment) (resp []*types.Shipment, err error) {
	req := c.request().SetMethod(http.MethodPost).SetJSON(v).ToJSON(&resp).SetPath("/shipments/bulk")
	return resp, c.send(ctx, req)
}

// GetLabel returns a shipment label in PDF format
// [GET]: /shipments/{id}
func (c *APIContext) GetLabel(ctx context.Context, id int) (resp *bytes.Buffer, err error) {
	resp = bytes.NewBuffer(nil)
	req := c.request().SetMethod(http.MethodGet).ToBytesBuffer(resp).SetPathf("/shipments/%d/label", id)
	return resp, c.send(ctx, req)
}

// GetLabels returns labels in PDF format
// [GET]: /shipments/labels/{id1,id2,...,idn}
func (c *APIContext) GetLabels(ctx context.Context, ids []int) (resp *bytes.Buffer, err error) {
	if len(ids) == 0 {
		return nil, types.ErrRequiredID
	}

	sids := make([]string, 0)
	for _, id := range ids {
		sids = append(sids, strconv.Itoa(id))
	}

	resp = bytes.NewBuffer(nil)
	req := c.request().SetMethod(http.MethodGet).ToBytesBuffer(resp).SetPathf("/shipments/labels/%s", strings.Join(sids, ","))
	return resp, c.send(ctx, req)
}
