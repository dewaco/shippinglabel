package shippinglabel

import (
	"context"
	"github.com/dewaco/shippinglabel/types"
	"os"
	"reflect"
	"sync"
	"testing"
)

var (
	client *Client
	api    *APIContext
)

func initClientAndAPIContext(tb *testing.T) {
	tb.Helper()
	if client != nil && api != nil {
		return
	}

	clientID, exists := os.LookupEnv("SHIPPINGLABEL_CLIENT_ID")
	if !exists {
		tb.Fatalf("could not found env: 'SHIPPINGLABEL_CLIENT_ID'")
	}

	secret, exists := os.LookupEnv("SHIPPINGLABEL_CLIENT_SECRET")
	if !exists {
		tb.Fatalf("could not found env: 'SHIPPINGLABEL_CLIENT_SECRET'")
	}
	var err error
	client, err = NewClient(clientID, secret)
	isNoError(tb, err)

	tk, err := client.ClientCredentials(context.Background())
	isNoError(tb, err)
	isNotNil(tb, tk)

	api, err = client.APIContext(tk)
	isNoError(tb, err)
	isNotNil(tb, api)
}

func TestClient_ClientCredentials(t *testing.T) {
	initClientAndAPIContext(t)
}

func TestClient_Error(t *testing.T) {
	initClientAndAPIContext(t)

	_, err := api.GetParcel(context.Background(), 0)
	if err == nil {
		t.Fatalf("expected error")
	}

	_, ok := err.(*types.Error)
	if !ok {
		t.Fatalf("expected error from type: %T", err)
	}

	switch err.(type) {
	case *types.Error:
		break
	default:
		t.Fatalf("expected error from type: %T", err)
	}
}

func TestClient_RefreshToken(t *testing.T) {
	initClientAndAPIContext(t)
	api.token.ExpiresIn = -5000
	api.token.SetExpirationTime()

	user, err := api.GetUser(context.Background())
	isNoError(t, err)
	isNotNil(t, user)
}

func TestAPIContext_GetUser(t *testing.T) {
	initClientAndAPIContext(t)

	user, err := api.GetUser(context.Background())
	isNoError(t, err)
	isNotNil(t, user)
}

func TestAPIContext_GetLabels(t *testing.T) {
	initClientAndAPIContext(t)

	ctx := context.Background()
	shipments, err := api.ListShipments(ctx, 0, 2, "asc")
	isNoError(t, err)

	if len(shipments) == 0 {
		t.Fatalf("expected shipments, but is empty")
	}

	ids := make([]int, 0)
	for _, shipment := range shipments {
		ids = append(ids, shipment.ID)
	}

	labels, err := api.GetLabels(ctx, ids)
	isNoError(t, err)
	isNotNil(t, labels)
}

func TestAPIContext_GetUserParallelism(t *testing.T) {
	initClientAndAPIContext(t)

	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			user, err := api.GetUser(context.Background())
			isNoError(t, err)
			isNotNil(t, user)
			wg.Done()
		}(&wg)
	}
	wg.Wait()
}

func TestAPIContext_ParcelCRUD(t *testing.T) {
	initClientAndAPIContext(t)

	var err error
	parcel := &types.Parcel{
		Name:        "Test",
		Description: "Test Parcel",
		Weight:      150,
	}
	ctx := context.Background()

	// Create parcel
	parcel, err = api.CreateParcel(ctx, parcel)
	isNoError(t, err)
	isNotNil(t, parcel)

	// Update parcel
	parcel.Weight = 200
	err = api.UpdateParcel(ctx, parcel)
	isNoError(t, err)

	// Get parcel
	serverParcel, err := api.GetParcel(ctx, parcel.ID)
	isNoError(t, err)
	isNotNil(t, serverParcel)
	isEqual(t, parcel, serverParcel)

	// List parcels
	ps, err := api.ListParcels(ctx)
	isNoError(t, err)
	isNotNil(t, ps)

	var exists bool
	for _, p := range ps {
		if p.ID == parcel.ID {
			exists = true
			break
		}
	}

	if !exists {
		t.Errorf("parcel (id: %d) not found", parcel.ID)
	}

	// Delete parcel
	err = api.DeleteParcel(ctx, parcel.ID)
	isNoError(t, err)
}

// Helper

func isNoError(tb testing.TB, err error) {
	tb.Helper()
	if err != nil {
		tb.Fatalf("unexpected error: %v", err)
	}
}

func isEqual(tb testing.TB, a any, b any) {
	if !reflect.DeepEqual(a, b) {
		tb.Fatalf("expected %#v (type %T) == %#v (type %T)", a, a, b, b)
	}
}

func isNotNil(tb testing.TB, v any) {
	if isNil(v) {
		tb.Fatalf("expected value %T is empty", v)
	}
}

func isNil(v any) bool {
	if v == nil {
		return true
	}

	value := reflect.ValueOf(v)
	kind := value.Kind()
	return kind >= reflect.Chan && kind <= reflect.Slice && value.IsNil()
}
