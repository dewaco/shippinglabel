package shippinglabel

import "time"

type CarrierCode string

const (
	CarrierDHL    CarrierCode = "DHL"
	CarrierDP     CarrierCode = "DP"
	CarrierDPD    CarrierCode = "DPD"
	CarrierGLS    CarrierCode = "GLS"
	CarrierHermes CarrierCode = "HERMES"
)

type CarrierServiceCode string

const (
	CarrierServicePreferredNeighbour  CarrierServiceCode = "PREFERRED_NEIGHBOUR"
	CarrierServicePreferredLocation   CarrierServiceCode = "PREFERRED_LOCATION"
	CarrierServiceVisualCheckOfAge    CarrierServiceCode = "VISUAL_CHECK_OF_AGE"
	CarrierServiceNamedPersonOnly     CarrierServiceCode = "NAMED_PERSON_ONLY"
	CarrierServiceIdentCheck          CarrierServiceCode = "IDENT_CHECK"
	CarrierServicePreferredDay        CarrierServiceCode = "PREFERRED_DAY"
	CarrierServiceNoNeighbourDelivery CarrierServiceCode = "NO_NEIGHBOUR_DELIVERY"
	CarrierServiceAdditionalInsurance CarrierServiceCode = "ADDITIONAL_INSURANCE"
	CarrierServiceBulkyGoods          CarrierServiceCode = "BULKY_GOODS"
	CarrierServiceCashOnDelivery      CarrierServiceCode = "CASH_ON_DELIVERY"
	CarrierServicePackagingReturn     CarrierServiceCode = "PACKAGING_RETURN"
	CarrierServiceParcelOutletRouting CarrierServiceCode = "PARCEL_OUTLET_ROUTING"
	CarrierServiceFlexDelivery        CarrierServiceCode = "FLEX_DELIVERY"
	CarrierServiceNextDay             CarrierServiceCode = "NEXT_DAY"
	CarrierServiceShopReturn          CarrierServiceCode = "SHOP_RETURN"
	CarrierServiceShopDelivery        CarrierServiceCode = "SHOP_DELIVERY"
	CarrierServiceIdentPin            CarrierServiceCode = "IDENT_PIN"
)

type Carrier struct {
	Code                 CarrierCode       `json:"carrierCode,omitempty"`
	Name                 string            `json:"name,omitempty"`
	IsDefault            bool              `json:"isDefault,omitempty"`
	Username             string            `json:"username,omitempty"`
	UserSecret           string            `json:"userSecret,omitempty"`
	UserSecretExpiration *time.Time        `json:"userSecretExpiration,omitempty"`
	Product              string            `json:"product,omitempty"`
	LabelFormat          string            `json:"labelFormat,omitempty"` // Default label format for shipments
	Created              *time.Time        `json:"created,omitempty"`
	Services             []*CarrierService `json:"services,omitempty"`   // Additional carrier services
	Parameters           map[string]any    `json:"parameters,omitempty"` // Additional parameters for the carrier (e.h. DHL EKP)
}

type CarrierService struct {
	Service    CarrierServiceCode `json:"service,omitempty"`
	Parameters map[string]any     `json:"parameters,omitempty"`
}

func (m *CarrierService) AddParameter(key string, param any) {
	if m.Parameters == nil {
		m.Parameters = make(map[string]any)
	}
	m.Parameters[key] = param
}
