package shippinglabel

type Parcel struct {
	ID           int            `json:"id,omitempty"`
	Name         string         `json:"name,omitempty"`
	Description  string         `json:"description,omitempty"`
	Weight       float64        `json:"weight,omitempty"`
	Length       float64        `json:"length,omitempty"`
	Width        float64        `json:"width,omitempty"`
	Height       float64        `json:"height,omitempty"`
	Reference    string         `json:"reference,omitempty"`
	CustomFields map[string]any `json:"customFields,omitempty"`
	IsDefault    bool           `json:"isDefault,omitempty"`
}
