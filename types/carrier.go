package types

import "time"

type CarrierCode string

const (
	CarrierDHL CarrierCode = "DHL"
	CarrierDP  CarrierCode = "DP"
	CarrierDPD CarrierCode = "DPD"
)

type Carrier struct {
	Code       CarrierCode `json:"carrierCode,omitempty"`
	Name       string      `json:"name,omitempty"`
	IsDefault  bool        `json:"isDefault,omitempty"`
	Created    *time.Time  `json:"created,omitempty"`
	Username   string      `json:"username,omitempty"`
	UserSecret string      `json:"userSecret,omitempty"`
	Product    string      `json:"product,omitempty"`
	DHLDetails *DHLDetails `json:"dhlDetails,omitempty"`
	DPDetails  *DPDetails  `json:"dpDetails,omitempty"`
	DPDDetails *DPDDetails `json:"dpdDetails,omitempty"`
}

type CarrierMetadata struct {
	DHLProducts []*DHLProduct     `json:"dhlProducts,omitempty"`
	DHLFormats  []*DHLLabelFormat `json:"dhlFormats,omitempty"`
	DPProducts  []*DPProduct      `json:"dpProducts,omitempty"`
	DPFormats   []*DPLabelFormat  `json:"dpFormats,omitempty"`
	DPDProducts []*DPDProduct     `json:"dpdProducts,omitempty"`
	DPDFormats  []*DPDLabelFormat `json:"dpdFormats,omitempty"`
}
