package shippinglabel

import (
	"bytes"
	"context"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

// APIContext represents the context for making API requests with an authenticated client and token.
//
// It provides methods for sending requests to different endpoints.
type APIContext struct {
	client     *Client
	token      *AuthToken
	tokenMutex sync.RWMutex
}

// NewAPIContext creates an API context
func NewAPIContext(c *Client, token *AuthToken) (*APIContext, error) {
	if c == nil {
		return nil, ErrRequiredClient
	}
	if token == nil {
		return nil, ErrRequiredToken
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
		tk, err := c.client.RefreshToken(ctx, c.token.RefreshToken)
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
func (c *APIContext) GetUser(ctx context.Context) (resp *User, err error) {
	req := c.request().SetMethod(http.MethodGet).ToJSON(&resp).SetPath("/user")
	return resp, c.send(ctx, req)
}

// METADATA

// Metadata returns the carrier metadata
// [GET]: /metadata/carriers
func (c *APIContext) Metadata(ctx context.Context) (resp []*CarrierMetadata, err error) {
	req := c.request().SetMethod(http.MethodGet).ToJSON(&resp).SetPath("/metadata/carriers")
	return resp, c.send(ctx, req)
}

// ADDRESSES

// ListAddresses returns all available user addresses
// [GET]: /addresses
func (c *APIContext) ListAddresses(ctx context.Context) (resp []*Address, err error) {
	req := c.request().SetMethod(http.MethodGet).ToJSON(&resp).SetPath("/addresses")
	return resp, c.send(ctx, req)
}

// CreateAddress creates a new shipment address
// [POST]: /addresses
func (c *APIContext) CreateAddress(ctx context.Context, v *Address) (resp *Address, err error) {
	req := c.request().SetMethod(http.MethodPost).SetJSON(v).ToJSON(&resp).SetPath("/addresses")
	return resp, c.send(ctx, req)
}

// GetAddress returns an address
// [GET]: /addresses/{id}
func (c *APIContext) GetAddress(ctx context.Context, id int) (resp *Address, err error) {
	req := c.request().SetMethod(http.MethodGet).ToJSON(&resp).SetPathf("/addresses/%d", id)
	return resp, c.send(ctx, req)
}

// UpdateAddress updates a shipment address
// [PUT]: /addresses/{id}
func (c *APIContext) UpdateAddress(ctx context.Context, v *Address) (err error) {
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
func (c *APIContext) ListParcels(ctx context.Context) (resp []*Parcel, err error) {
	req := c.request().SetMethod(http.MethodGet).ToJSON(&resp).SetPath("/parcels")
	return resp, c.send(ctx, req)
}

// CreateParcel creates a parcel
// [POST]: /parcels
func (c *APIContext) CreateParcel(ctx context.Context, v *Parcel) (resp *Parcel, err error) {
	req := c.request().SetMethod(http.MethodPost).ToJSON(&resp).SetJSON(v).SetPath("/parcels")
	return resp, c.send(ctx, req)
}

// GetParcel returns a parcel
// [GET]: /parcels/{id}
func (c *APIContext) GetParcel(ctx context.Context, id int) (resp *Parcel, err error) {
	req := c.request().SetMethod(http.MethodGet).ToJSON(&resp).SetPathf("/parcels/%d", id)
	return resp, c.send(ctx, req)
}

// UpdateParcel updates a parcel
// [PUT]: /parcels/{id}
func (c *APIContext) UpdateParcel(ctx context.Context, v *Parcel) (err error) {
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
func (c *APIContext) ListCarriers(ctx context.Context) (resp []*Carrier, err error) {
	req := c.request().SetMethod(http.MethodGet).ToJSON(&resp).SetPath("/carriers")
	return resp, c.send(ctx, req)
}

// CreateCarrier creates a carrier
// [POST]: /carriers
func (c *APIContext) CreateCarrier(ctx context.Context, v *Carrier) (resp *Carrier, err error) {
	req := c.request().SetMethod(http.MethodPost).ToJSON(&resp).SetJSON(v).SetPath("/carriers")
	return resp, c.send(ctx, req)
}

// GetCarrier returns a carrier
// [GET]: /carriers/{id}
func (c *APIContext) GetCarrier(ctx context.Context, code CarrierCode) (resp *Carrier, err error) {
	req := c.request().SetMethod(http.MethodGet).ToJSON(&resp).SetPathf("/carriers/%s", code)
	return resp, c.send(ctx, req)
}

// UpdateCarrier updates a carrier
// [PUT]: /carriers/{id}
func (c *APIContext) UpdateCarrier(ctx context.Context, v *Carrier) (err error) {
	req := c.request().SetMethod(http.MethodPut).SetJSON(v).SetPathf("/carriers/%s", v.Code)
	return c.send(ctx, req)
}

// UpdateCarrierCredentials updates the user credentials from the carrier
// [PUT]: /carriers/{id}/cred
func (c *APIContext) UpdateCarrierCredentials(ctx context.Context, v *Carrier) (err error) {
	req := c.request().SetMethod(http.MethodPut).SetJSON(v).SetPathf("/carriers/%s/credentials", v.Code)
	return c.send(ctx, req)
}

// VerifyCarrier validates the user credentials
// [POST]: /carriers/{id}/verify
func (c *APIContext) VerifyCarrier(ctx context.Context, code CarrierCode) (err error) {
	req := c.request().SetMethod(http.MethodPost).SetPathf("/carriers/%s/verify", code)
	return c.send(ctx, req)
}

// DeleteCarrier deletes a carrier
// [DELETE]: /carriers/{id}
func (c *APIContext) DeleteCarrier(ctx context.Context, code CarrierCode) (err error) {
	req := c.request().SetMethod(http.MethodDelete).SetPathf("/carriers/%s", code)
	return c.send(ctx, req)
}

// CARRIER PRODUCTS

// CreateDHLProduct creates a DHL product
// [POST]: /carriers/DHL/products
func (c *APIContext) CreateDHLProduct(ctx context.Context, v *Product) (resp *Product, err error) {
	req := c.request().SetMethod(http.MethodPost).SetJSON(v).ToJSON(&resp).SetPath("/carriers/DHL/products")
	return resp, c.send(ctx, req)
}

// UpdateDHLProduct updates a DHL product
// [PUT]: /carriers/DHL/products/{id}
func (c *APIContext) UpdateDHLProduct(ctx context.Context, v *Product) (err error) {
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
func (c *APIContext) ListShipments(ctx context.Context, page int, size int, order string) (resp []*Shipment, err error) {
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
func (c *APIContext) ValidateShipment(ctx context.Context, v *Shipment) (err error) {
	req := c.request().SetMethod(http.MethodPost).SetJSON(v).SetPath("/shipments/validate")
	return c.send(ctx, req)
}

// CreateShipment creates a shipment
// [POST]: /shipments
func (c *APIContext) CreateShipment(ctx context.Context, v *Shipment) (resp *Shipment, err error) {
	req := c.request().SetMethod(http.MethodPost).SetJSON(v).ToJSON(&resp).SetPath("/shipments")
	return resp, c.send(ctx, req)
}

// GetShipment returns a shipment
// [GET]: /shipments/{id}
func (c *APIContext) GetShipment(ctx context.Context, id int) (resp *Shipment, err error) {
	req := c.request().SetMethod(http.MethodGet).ToJSON(&resp).SetPathf("/shipments/%d", id)
	return resp, c.send(ctx, req)
}

// DeleteShipment deletes a shipment
// [DELETE]: /shipments/{id}
func (c *APIContext) DeleteShipment(ctx context.Context, id int) (err error) {
	req := c.request().SetMethod(http.MethodDelete).SetPathf("/shipments/%d", id)
	return c.send(ctx, req)
}

// CreateShipments creates multiple shipments
// [POST]: /shipments/bulk
func (c *APIContext) CreateShipments(ctx context.Context, v []*Shipment) (resp []*Shipment, err error) {
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
// Value: 'ids' can be from type []string or []int
func (c *APIContext) GetLabels(ctx context.Context, ids any) (resp *bytes.Buffer, err error) {
	var sIDs []string

	switch ids.(type) {
	case []string:
		sIDs = ids.([]string)
	case []int:
		for _, id := range ids.([]int) {
			sIDs = append(sIDs, strconv.Itoa(id))
		}
	default:
		return nil, ErrWrongType
	}

	if len(sIDs) == 0 {
		return nil, ErrRequiredID
	}

	resp = bytes.NewBuffer(nil)
	req := c.request().SetMethod(http.MethodGet).ToBytesBuffer(resp).SetPathf("/shipments/labels/%s", strings.Join(sIDs, ","))
	return resp, c.send(ctx, req)
}

// QUEUE ITEMS

// ListQueueItems returns all queue items
// [GET]: /shipments/queue
func (c *APIContext) ListQueueItems(ctx context.Context) (resp []*ShipmentQueueItem, err error) {
	req := c.request().SetMethod(http.MethodGet).ToJSON(&resp).SetPath("/shipments/queue")
	return resp, c.send(ctx, req)
}

// CreateQueueItem creates a queue item
// [POST]: /shipments/queue
func (c *APIContext) CreateQueueItem(ctx context.Context, v *ShipmentQueueItem) (resp *ShipmentQueueItem, err error) {
	req := c.request().SetMethod(http.MethodPost).ToJSON(&resp).SetJSON(v).SetPath("/shipments/queue")
	return resp, c.send(ctx, req)
}

// GetQueueItem returns a queue item
// [GET]: /shipments/queue/{id}
func (c *APIContext) GetQueueItem(ctx context.Context, id int) (resp *ShipmentQueueItem, err error) {
	req := c.request().SetMethod(http.MethodGet).ToJSON(&resp).SetPathf("/shipments/queue/%d", id)
	return resp, c.send(ctx, req)
}

// UpdateQueueItem updates a queue item
// [PUT]: /shipments/queue/{id}
func (c *APIContext) UpdateQueueItem(ctx context.Context, v *ShipmentQueueItem) (err error) {
	req := c.request().SetMethod(http.MethodPut).SetJSON(v).SetPathf("/shipments/queue/%d", v.ID)
	return c.send(ctx, req)
}

// DeleteQueueItem deletes a queue item
// [DELETE]: /shipments/queue/{id}
func (c *APIContext) DeleteQueueItem(ctx context.Context, id int) (err error) {
	req := c.request().SetMethod(http.MethodDelete).SetPathf("/shipments/queue/%d", id)
	return c.send(ctx, req)
}

// UploadCSVFile uploads a csv file and sets the items in the shipment queue
// [POST]: /shipments/queue/csv
func (c *APIContext) UploadCSVFile(ctx context.Context, csv []byte, csvProfileID int) (resp []*ShipmentQueueItem, err error) {
	req := c.request().SetMethod(http.MethodPost).SetBytes(csv).ToJSON(&resp).SetPathf("/shipments/queue/csv?profileId=%d", csvProfileID)
	return resp, c.send(ctx, req)
}

// JOBS

// ListJobs returns all jobs
// [GET]: /shipments/jobs
func (c *APIContext) ListJobs(ctx context.Context) (resp []*ShipmentQueueItem, err error) {
	req := c.request().SetMethod(http.MethodGet).ToJSON(&resp).SetPath("/shipments/jobs")
	return resp, c.send(ctx, req)
}

// CreateJob creates a job
// [POST]: /shipments/jobs
func (c *APIContext) CreateJob(ctx context.Context, v *ShipmentQueueItem) (resp *ShipmentQueueItem, err error) {
	req := c.request().SetMethod(http.MethodPost).ToJSON(&resp).SetJSON(v).SetPath("/shipments/jobs")
	return resp, c.send(ctx, req)
}

// GetJob returns a job
// [GET]: /shipments/jobs/{id}
func (c *APIContext) GetJob(ctx context.Context, id int) (resp *ShipmentQueueItem, err error) {
	req := c.request().SetMethod(http.MethodGet).ToJSON(&resp).SetPathf("/shipments/jobs/%d", id)
	return resp, c.send(ctx, req)
}

// UpdateJob updates a job
// [PUT]: /shipments/jobs/{id}
func (c *APIContext) UpdateJob(ctx context.Context, v *ShipmentQueueItem) (err error) {
	req := c.request().SetMethod(http.MethodPut).SetJSON(v).SetPathf("/shipments/jobs/%d", v.ID)
	return c.send(ctx, req)
}

// DeleteJob deletes a job
// [DELETE]: /shipments/jobs/{id}
func (c *APIContext) DeleteJob(ctx context.Context, id int) (err error) {
	req := c.request().SetMethod(http.MethodDelete).SetPathf("/shipments/jobs/%d", id)
	return c.send(ctx, req)
}

// CSV

// ListCSVProfiles returns all csv profiles
// [GET]: /csv/profiles
func (c *APIContext) ListCSVProfiles(ctx context.Context) (resp []*CSVProfile, err error) {
	req := c.request().SetMethod(http.MethodGet).ToJSON(&resp).SetPath("/csv/profiles")
	return resp, c.send(ctx, req)
}

// CreateCSVProfile creates a csv profile
// [POST]: /csv/profiles
func (c *APIContext) CreateCSVProfile(ctx context.Context, v *CSVProfile) (resp *CSVProfile, err error) {
	req := c.request().SetMethod(http.MethodPost).ToJSON(&resp).SetJSON(v).SetPath("/csv/profiles")
	return resp, c.send(ctx, req)
}

// GetCSVProfile returns a csv profile
// [GET]: /csv/profiles/{id}
func (c *APIContext) GetCSVProfile(ctx context.Context, id int) (resp *CSVProfile, err error) {
	req := c.request().SetMethod(http.MethodGet).ToJSON(&resp).SetPathf("/csv/profiles/%d", id)
	return resp, c.send(ctx, req)
}

// UpdateCSVProfile updates a csv profile
// [PUT]: /csv/profiles/{id}
func (c *APIContext) UpdateCSVProfile(ctx context.Context, v *CSVProfile) (err error) {
	req := c.request().SetMethod(http.MethodPut).SetJSON(v).SetPathf("/csv/profiles/%d", v.ID)
	return c.send(ctx, req)
}

// DeleteCSVProfile deletes a csv profile
// [DELETE]: /csv/profiles/{id}
func (c *APIContext) DeleteCSVProfile(ctx context.Context, id int) (err error) {
	req := c.request().SetMethod(http.MethodDelete).SetPathf("/csv/profiles/%d", id)
	return c.send(ctx, req)
}
