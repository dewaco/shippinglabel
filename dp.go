package shippinglabel

type DPDetails struct {
	LabelFormat    string `json:"labelFormat,omitempty"`
	DefaultProduct string `json:"defaultProduct,omitempty"`
	OrderID        string `json:"orderId,omitempty"`
}

type DPProduct struct {
	Product         string  `json:"product,omitempty"`
	Name            string  `json:"name,omitempty"`
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
	LabelFormat       string `json:"labelFormat,omitempty"`
	Name              string `json:"name,omitempty"`
	IsAddressPossible bool   `json:"isAddressPossible,omitempty"`
	LabelCountX       int    `json:"labelCountX,omitempty"`
	LabelCountY       int    `json:"labelCountY,omitempty"`
}
