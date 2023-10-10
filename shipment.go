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

type ExportTypeCode string

const (
	ExportTypeOther           ExportTypeCode = "OTHER"
	ExportTypePresent         ExportTypeCode = "PRESENT"
	ExportTypeSample          ExportTypeCode = "SAMPLE"
	ExportTypeDocument        ExportTypeCode = "DOCUMENT"
	ExportTypeReturnOfGoods   ExportTypeCode = "RETURN_OF_GOODS"
	ExportTypeCommercialGoods ExportTypeCode = "COMMERCIAL_GOODS"
)

type Customs struct {
	ExportType                      ExportTypeCode `json:"exportType,omitempty"`
	ExportDescription               string         `json:"exportDescription,omitempty"`
	ShippingCosts                   *Amount        `json:"shippingCosts,omitempty"`
	InvoiceNumber                   string         `json:"invoiceNumber,omitempty"`
	SenderCustomsReference          string         `json:"senderCustomsReference,omitempty"`
	ReceiverCustomsReference        string         `json:"receiverCustomsReference,omitempty"`
	HasElectronicExportNotification bool           `json:"hasElectronicExportNotification,omitempty"`
	Items                           []*CustomsItem `json:"items,omitempty"`
}

type CustomsItem struct {
	Description   string  `json:"description,omitempty"`
	Quantity      int     `json:"quantity,omitempty"`
	HsCode        string  `json:"hsCode,omitempty"`
	OriginCountry string  `json:"originCountry,omitempty"`
	UnitValue     *Amount `json:"unitValue,omitempty"`
	Weight        float64 `json:"weight,omitempty"`
}

type Amount struct {
	Value    float64 `json:"value,omitempty"`
	Currency string  `json:"currency,omitempty"`
}
