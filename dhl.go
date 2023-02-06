package shippinglabel

import "time"

type DHLServiceCode string

const (
	DHLServicePreferredNeighbour          DHLServiceCode = "PREFERRED_NEIGHBOUR"
	DHLServicePreferredLocation           DHLServiceCode = "PREFERRED_LOCATION"
	DHLServiceVisualCheckOfAge            DHLServiceCode = "VISUAL_CHECK_OF_AGE"
	DHLServiceNamedPersonOnly             DHLServiceCode = "NAMED_PERSON_ONLY"
	DHLServiceIdentCheck                  DHLServiceCode = "IDENT_CHECK"
	DHLServicePreferredDay                DHLServiceCode = "PREFERRED_DAY"
	DHLServiceNoNeighbourDelivery         DHLServiceCode = "NO_NEIGHBOUR_DELIVERY"
	DHLServiceAdditionalInsurance         DHLServiceCode = "ADDITIONAL_INSURANCE"
	DHLServiceBulkyGoods                  DHLServiceCode = "BULKY_GOODS"
	DHLServiceCashOnDelivery              DHLServiceCode = "CASH_ON_DELIVERY"
	DHLServiceIndividualSenderRequirement DHLServiceCode = "INDIVIDUAL_SENDER_REQUIREMENT"
	DHLServicePackagingReturn             DHLServiceCode = "PACKAGING_RETURN"
	DHLServiceParcelOutletRouting         DHLServiceCode = "PARCEL_OUTLET_ROUTING"
)

type DHLDetails struct {
	EKP         string        `json:"ekp,omitempty"`
	LabelFormat string        `json:"labelFormat,omitempty"`
	Products    []*DHLProduct `json:"products,omitempty"`
	Services    []*DHLService `json:"services,omitempty"`
}

type DHLProduct struct {
	ID                          int    `json:"id,omitempty"`
	Product                     string `json:"product,omitempty"`
	Name                        string `json:"name,omitempty"`
	Procedure                   string `json:"procedure,omitempty"`
	UserParticipation           string `json:"userParticipation,omitempty"`
	PreferredNeighbour          bool   `json:"preferredNeighbour,omitempty"`
	PreferredLocation           bool   `json:"preferredLocation,omitempty"`
	VisualCheckOfAge            bool   `json:"visualCheckOfAge,omitempty"`
	NamedPersonOnly             bool   `json:"namedPersonOnly,omitempty"`
	IdentCheck                  bool   `json:"identCheck,omitempty"`
	PreferredDay                bool   `json:"preferredDay,omitempty"`
	NoNeighbourDelivery         bool   `json:"noNeighbourDelivery,omitempty"`
	GoGreen                     bool   `json:"goGreen,omitempty"`
	AdditionalInsurance         bool   `json:"additionalInsurance,omitempty"`
	BulkyGoods                  bool   `json:"bulkyGoods,omitempty"`
	CashOnDelivery              bool   `json:"cashOnDelivery,omitempty"`
	IndividualSenderRequirement bool   `json:"individualSenderRequirement,omitempty"`
	PackagingReturn             bool   `json:"packagingReturn,omitempty"`
	ParcelOutletRouting         bool   `json:"parcelOutletRouting,omitempty"`
}

type DHLService struct {
	Service    DHLServiceCode `json:"service,omitempty"`
	Value      string         `json:"value,omitempty"`
	FloatValue float64        `json:"floatValue,omitempty"`
	DateValue  *time.Time     `json:"dateValue,omitempty"`

	// COD
	AccountOwner     string `json:"accountOwner,omitempty"`
	BankName         string `json:"bankName,omitempty"`
	IBAN             string `json:"iban,omitempty"`
	BIC              string `json:"bic,omitempty"`
	Note             string `json:"note,omitempty"`
	Note2            string `json:"note2,omitempty"`
	AccountReference string `json:"accountReference,omitempty"`

	// IdentCheck
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"LastName,omitempty"`
}

type DHLLabelFormat struct {
	LabelFormat string `json:"labelFormat,omitempty"`
	Name        string `json:"name,omitempty"`
}
