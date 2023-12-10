package shippinglabel

import "time"

type Shipment struct {
	ID                int            `json:"id,omitempty"`
	ShipmentNumber    string         `json:"shipmentNumber,omitempty"`
	Carrier           *Carrier       `json:"carrier,omitempty"`
	Parcels           []*Parcel      `json:"parcels,omitempty"`
	Sender            *Address       `json:"sender,omitempty"`
	Receiver          *Address       `json:"receiver,omitempty"`
	Reference         string         `json:"reference,omitempty"`
	ShipmentDate      *time.Time     `json:"shipmentDate,omitempty"`
	Customs           *Customs       `json:"customs,omitempty"`
	Created           *time.Time     `json:"created,omitempty"`
	AdditionalDetails map[string]any `json:"additionalDetails,omitempty"`
	Label             string         `json:"label,omitempty"`
	Status            *Status        `json:"status,omitempty"`
}

type ShipmentQueueItem struct {
	ID          int        `json:"id,omitempty"`
	Shipment    *Shipment  `json:"shipment,omitempty"`
	ProcessTime *time.Time `json:"processTime,omitempty"`
	Created     *time.Time `json:"created,omitempty"`
}
