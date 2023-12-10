package shippinglabel

type CarrierMetadata struct {
	Code         CarrierCode    `json:"carrierCode,omitempty"`
	Name         string         `json:"name,omitempty"`
	Products     []*Product     `json:"products,omitempty"`
	LabelFormats []*LabelFormat `json:"labelFormats,omitempty"`
}

type Product struct {
	ID                int     `json:"id,omitempty"`
	Product           string  `json:"product,omitempty"`
	Name              string  `json:"name,omitempty"`
	UserParticipation string  `json:"userParticipation,omitempty"`
	Annotation        string  `json:"annotation,omitempty"`
	Price             float64 `json:"price,omitempty"`
	IsInternational   bool    `json:"isInternational,omitempty"`
	MinWeight         int     `json:"minWeight,omitempty"`
	MaxWeight         int     `json:"maxWeight,omitempty"`
	MinLength         int     `json:"minLength,omitempty"`
	MaxLength         int     `json:"maxLength,omitempty"`
	MinWidth          int     `json:"minWidth,omitempty"`
	MaxWidth          int     `json:"maxWidth,omitempty"`
	MinHeight         int     `json:"minHeight,omitempty"`
	MaxHeight         int     `json:"maxHeight,omitempty"`
}

type LabelFormat struct {
	ID              int    `json:"id,omitempty"`
	LabelFormat     string `json:"labelFormat,omitempty"`
	Name            string `json:"name,omitempty"`
	HasAddressField bool   `json:"hasAddressField,omitempty"`
	LabelCountX     int    `json:"labelCountX,omitempty"`
	LabelCountY     int    `json:"labelCountY,omitempty"`
}
