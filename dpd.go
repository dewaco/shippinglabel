package shippinglabel

type DPDDetails struct {
	DefaultProduct string `json:"defaultProduct,omitempty"`
	LabelFormat    string `json:"labelFormat,omitempty"`
}

type DPDProduct struct {
	Product string `json:"product,omitempty"`
	Name    string `json:"name,omitempty"`
}

type DPDLabelFormat struct {
	LabelFormat string `json:"labelFormat,omitempty"`
	Name        string `json:"name,omitempty"`
}
