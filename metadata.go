package shippinglabel

type CarrierMetadata struct {
	Carriers []*Carrier `json:"carriers,omitempty"`
	DHL      struct {
		Products     []*DHLProduct  `json:"products,omitempty"`
		LabelFormats []*LabelFormat `json:"labelFormats,omitempty"`
	} `json:"dhl,omitempty"`
	DP struct {
		Products     []*DPProduct     `json:"products,omitempty"`
		LabelFormats []*DPLabelFormat `json:"labelFormats,omitempty"`
	} `json:"dp,omitempty"`
	DPD struct {
		Products     []*Product     `json:"products,omitempty"`
		LabelFormats []*LabelFormat `json:"labelFormats,omitempty"`
	} `json:"dpd,omitempty"`
}

type Product struct {
	Product string `json:"product,omitempty"`
	Name    string `json:"name,omitempty"`
}

type LabelFormat struct {
	LabelFormat string `json:"labelFormat,omitempty"`
	Name        string `json:"name,omitempty"`
}

type DHLProduct struct {
	ID int `json:"id,omitempty"`
	Product
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

type DPProduct struct {
	Product
	Annotation      string  `json:"annotation,omitempty"`
	Price           float64 `json:"price,omitempty"`
	IsInternational bool    `json:"isInternational,omitempty"`
	MinWeight       int     `json:"minWeight,omitempty"`
	MaxWeight       int     `json:"maxWeight,omitempty"`
	MinLength       int     `json:"minLength,omitempty"`
	MaxLength       int     `json:"maxLength,omitempty"`
	MinWidth        int     `json:"minWidth,omitempty"`
	MaxWidth        int     `json:"maxWidth,omitempty"`
	MinHeight       int     `json:"minHeight,omitempty"`
	MaxHeight       int     `json:"maxHeight,omitempty"`
}

type DPLabelFormat struct {
	LabelFormat
	IsAddressPossible bool `json:"isAddressPossible,omitempty"`
	LabelCountX       int  `json:"labelCountX,omitempty"`
	LabelCountY       int  `json:"labelCountY,omitempty"`
}
